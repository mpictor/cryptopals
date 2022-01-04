// Copyright (C) 2021 the Authors. All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// SPDX-License-Identifier: BSD-3-Clause
//

package cmd

import (
	"cryptopals/lib"
	"cryptopals/lib/key"
	"cryptopals/lib/xor"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:   "find [flags] [input-file]",
	Short: "find an xor key",
	Long: `Using the given constraints, find key that produces
the most realistic plaintext of the given type (format).`,
	Run: func(cmd *cobra.Command, args []string) {
		pt, err := key.PTFromStr(inFmt)
		cobra.CheckErr(err)

		ksd := key.KeySearchData{
			Data:   readIn(args),
			PTType: pt,
		}
		ksd.SizeSearch(int(min), int(max))
		if lib.V {
			fmt.Println("best key sizes:", ksd.Sizes)
		}
		ksd.KeySearch()
		if lib.V {
			fmt.Printf("best key: %s\n", ksd.BestKey)
		}
		if len(ksd.BestKey.K) == 0 {
			return
		}
		if len(kout) > 0 {
			cobra.CheckErr(os.WriteFile(kout, ksd.BestKey.K, 0644))
		}
		if lib.VV {
			fmt.Printf("plaintext: %q [...]\n", string(xor.EncryptXor(ksd.Data, ksd.BestKey.K))[:60])
		}
		if len(out) > 0 {
			cobra.CheckErr(os.WriteFile(out, xor.EncryptXor(ksd.Data, ksd.BestKey.K), 0644))
		}
	},
}

var (
	inFmt, kout, out string
	min, max         uint
)

func init() {
	xorCmd.AddCommand(findCmd)

	findCmd.Flags().StringVar(&inFmt, "fmt", "arm32le", "expected format to score keys against")
	findCmd.Flags().UintVarP(&min, "min", "n", 2, "minimum key length to search")
	findCmd.Flags().UintVarP(&max, "max", "x", 8, "maximum key length to search")
	addReadInFlags(findCmd)
	findCmd.Flags().StringVar(&out, "out", "", "output file name")
	findCmd.Flags().StringVar(&kout, "kout", "", "output key to file")
}
