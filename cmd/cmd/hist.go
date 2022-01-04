// Copyright (C) 2021 the Authors. All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// SPDX-License-Identifier: BSD-3-Clause
//

package cmd

import (
	"cryptopals/lib"
	"encoding/binary"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

// histCmd represents the hist command
var histCmd = &cobra.Command{
	Use:   "hist [flags] [file]",
	Short: "histogram",
	Long:  "finds the most freqent sequences and prints them",
	Run: func(cmd *cobra.Command, args []string) {
		data := readIn(args)
		hmap := hist(data, width)
		if lib.V {
			fmt.Println(len(data), "bytes data;", len(hmap), "items")
		}
		elems := hmap.topElems(topN)
		fmt.Println(elems)
	},
}

type ele struct {
	n, val uint64
}

func (e ele) String() string {
	if !lib.V {
		return fmt.Sprintf("0x%0*x", width*2, e.val)
	}
	return fmt.Sprintf("% 4d: 0x%0*x", e.n, width*2, e.val)
}

type elems []ele

func (eles elems) Len() int           { return len(eles) }
func (eles elems) Swap(i, j int)      { eles[i], eles[j] = eles[j], eles[i] }
func (eles elems) Less(i, j int) bool { return eles[i].n < eles[j].n }

func (eles elems) String() string {
	var s []string
	for _, e := range eles {
		s = append(s, e.String())
	}
	if !lib.V {
		return strings.Join(s, ", ")
	}
	return strings.Join(s, "\n")
}

func hist(data []byte, width uint) hmap {
	hmap := make(hmap)
	for len(data) >= int(width) {
		switch width {
		case 4:
			u := binary.LittleEndian.Uint32(data[:4])
			n := hmap[uint64(u)]
			hmap[uint64(u)] = n + 1
		default:
			fmt.Fprintf(os.Stderr, "width %d is not implemented\n", width)
			os.Exit(1)
		}
		data = data[width:]
	}
	return hmap
}

type hmap map[uint64]uint64

func (hm hmap) topElems(n uint) (elems elems) {
	if n >= uint(len(hm)) {
		return hm.toSorted()
	}
	sorted := hm.toSorted()
	return sorted[:n]
}

//return slice, with first being highest
func (hm hmap) toSorted() (elems elems) {
	for k, v := range hm {
		elems = append(elems, ele{
			n:   v,
			val: k,
		})
	}
	sort.Sort(sort.Reverse(elems))
	return
}

var (
	width, topN uint
)

func init() {
	rootCmd.AddCommand(histCmd)
	histCmd.Flags().UintVarP(&width, "width", "w", 4, "arch width, max 8")
	histCmd.Flags().UintVarP(&topN, "top", "n", 10, "number to display")
	addReadInFlags(histCmd)
}
