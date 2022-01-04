// Copyright (C) 2021 the Authors. All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// SPDX-License-Identifier: BSD-3-Clause
//

package cmd

import (
	"bufio"
	"cryptopals/lib"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var (
	cfgFile string
	v       int
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cmd",
	Short: "encryption utilities",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig, initVerbose)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cmd.yaml)")
	rootCmd.PersistentFlags().CountVarP(&v, "verbose", "v", "verbosity; -vv for more")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cmd" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cmd")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func initVerbose() {
	lib.V = v > 0
	lib.VV = v > 1
	if v > 2 {
		fmt.Fprintln(os.Stderr, "sorry, exceeded max verbosity")
	}
}

var (
	offs, length uint
	b64          bool
	in           string
	elfSect      string
)

//reads input file `in`, honoring `b64`
func readIn(args []string) []byte {
	if len(elfSect) > 0 {
		fmt.Fprintf(os.Stderr, "--elfsect unimplemented, sorry\n")
		os.Exit(1)
	}
	if len(args) == 0 {
		if len(in) == 0 {
			fmt.Fprintf(os.Stderr, "requires either --in or a trailing input-file arg\n")
			os.Exit(1)
		}
	} else {
		in = args[0]
	}
	var data []byte
	var err error
	if !b64 {
		data, err = os.ReadFile(in)
		cobra.CheckErr(err)
	} else {
		f, err := os.Open(in)
		cobra.CheckErr(err)
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			l, err := base64.StdEncoding.DecodeString(scanner.Text())
			cobra.CheckErr(err)
			data = append(data, l...)
		}
		cobra.CheckErr(scanner.Err())
	}
	if offs > 0 {
		data = data[offs:]
	}
	if length > uint(len(data)) {
		data = data[:length]
	}
	return data
}

func addReadInFlags(cmd *cobra.Command) {
	cmd.Flags().UintVarP(&offs, "off", "o", 0, "offset into input file")
	cmd.Flags().UintVarP(&length, "len", "l", 0, "length to read from offset (0 for until EOF)")
	cmd.Flags().StringVar(&in, "in", "", "input file")
	cmd.Flags().StringVar(&elfSect, "elf", "", "name of elf section to work on (unimplemented)")
	cmd.Flags().BoolVar(&b64, "base64", false, "true for base64-encoded input, false for binary")
}
