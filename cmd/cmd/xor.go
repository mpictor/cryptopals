// Copyright (C) 2021 the Authors. All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// SPDX-License-Identifier: BSD-3-Clause
//

package cmd

import (
	"cryptopals/lib/xor"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// xorCmd represents the xor command
var xorCmd = &cobra.Command{
	Use:   "xor",
	Short: "encrypt/decrypt xor",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(out) == 0 || len(xkey) == 0 {
			return fmt.Errorf("--out and --key are required")
		}
		data := readIn(args)
		keydata, err := os.ReadFile(xkey)
		if err != nil {
			return err
		}
		outdata := xor.EncryptXor(data, keydata)
		return os.WriteFile(out, outdata, 0644)
	},
}

var tweakCmd = &cobra.Command{
	Use:   "tweak",
	Short: "tweak xor key byte",
	Long: `Use when you have a close-but-not-exact key and know what one
particular byte in the output should be (i.e. a word is misspelled).

Provide ciphertext, old key, offset to incorrect byte, and the desired
value for that byte. Computes new key; writes new output and key out.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(out) == 0 || len(xkey) == 0 || len(wantChar) != 1 || byteOffs == 0 {
			return fmt.Errorf("--out, --key, --byte, and --want are required")
		}
		data := readIn(args)
		keydata, err := os.ReadFile(xkey)
		if err != nil {
			return err
		}
		oldData := xor.EncryptXor(data, keydata)
		had := data[byteOffs]
		got := oldData[byteOffs]
		koffs := byteOffs % len(keydata)
		kold := keydata[koffs]
		if kold^had != got {
			return fmt.Errorf("sanity check failed: koffs=%d kold=0x%x had=0x%x got=0x%x", koffs, kold, had, got)
		}
		knew := had ^ wantChar[0]
		fmt.Fprintf(os.Stderr, "replacing key char at position %d (0x%02x) with 0x%02x\n", koffs, kold, knew)
		newkey := make([]byte, len(keydata))
		copy(newkey, keydata)
		newkey[koffs] = knew
		outdata := xor.EncryptXor(data, newkey)
		if err := os.WriteFile(out+".key", newkey, 0644); err != nil {
			return err
		}
		return os.WriteFile(out, outdata, 0644)
	},
}

var (
	xkey     string
	byteOffs int
	wantChar string
)

func init() {
	rootCmd.AddCommand(xorCmd)
	xorCmd.Flags().StringVar(&out, "out", "", "output file name")
	xorCmd.Flags().StringVarP(&xkey, "key", "k", "", "key file name")
	addReadInFlags(xorCmd)

	xorCmd.AddCommand(tweakCmd)
	tweakCmd.Flags().StringVar(&out, "out", "", "output file name")
	tweakCmd.Flags().StringVarP(&xkey, "key", "k", "", "key file name")
	tweakCmd.Flags().IntVarP(&byteOffs, "byte", "b", 0, "offset to file byte to tweak")
	tweakCmd.Flags().StringVarP(&wantChar, "want", "w", "", "char we want at this position, i.e. '\\x123'")
	addReadInFlags(tweakCmd)
}
