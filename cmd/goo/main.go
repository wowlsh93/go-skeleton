/*
Copyright Stick Corp. All Rights Reserved.

Written by HAMA
*/

package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wowlsh93/go-skeleton/core/Aapp"
	"github.com/wowlsh93/go-skeleton/core/version"
	"os"
)

var mainCmd = &cobra.Command{
	Use:   "goo",
	Short: "go skeleton",
	Long:  `go skeleton`,
	//	Run: func(cmd *cobra.Command, args []string) { },
}

func main() {

	mainCmd.AddCommand(Aapp.Cmd())
	mainCmd.AddCommand(version.Cmd())

	if err := mainCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}