// Copyright (C) 2021 the Authors. All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// SPDX-License-Identifier: BSD-3-Clause
//

package cmd

import (
	"cryptopals/lib/key"
	"cryptopals/lib/xor"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// scoreCmd represents the score command
var scoreCmd = &cobra.Command{
	Use:   "score",
	Short: "compute realness score",
	Long: `Given a key and plaintext, compute realness score. 
This is the same as used when scoring candidate keys.`,
	Run: func(cmd *cobra.Command, args []string) {
		ptt, err := key.PTFromStr(inFmt)
		cobra.CheckErr(err)

		data := readIn(args)
		fmt.Println("score:", key.ScoreSeq(data, ptt))
	},
}

// like scoreCmd but a subcommand of xor, with additional key arg
var xorScoreCmd = &cobra.Command{
	Use:   "score",
	Short: "given key, compute realness score",
	Long: `Given a key and ciphertext, compute realness score. 
This is the same as used when scoring candidate keys.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(xkey) == 0 {
			return fmt.Errorf("--key is required")
		}
		data := readIn(args)
		keydata, err := os.ReadFile(xkey)
		if err != nil {
			return err
		}
		outdata := xor.EncryptXor(data, keydata)

		ptt, err := key.PTFromStr(inFmt)
		if err != nil {
			return err
		}

		fmt.Println("score:", key.ScoreSeq(outdata, ptt))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(scoreCmd)
	scoreCmd.Flags().StringVarP(&inFmt, "fmt", "f", "arm32le", "format to try")
	addReadInFlags(scoreCmd)

	xorCmd.AddCommand(xorScoreCmd)
	xorScoreCmd.Flags().StringVarP(&inFmt, "fmt", "f", "arm32le", "format to try")
	xorScoreCmd.Flags().StringVarP(&xkey, "key", "k", "", "key file name")
	addReadInFlags(xorScoreCmd)
}
