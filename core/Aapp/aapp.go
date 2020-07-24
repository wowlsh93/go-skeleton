/*
Copyright Stick Corp. All Rights Reserved.

Written by HAMA
*/

package Aapp

import (
	"fmt"
	"github.com/spf13/cobra"
)

const (
	scannerFuncName = "aapp"
	scannerCmdDes   = "aapp start"
)

func Cmd() *cobra.Command {
	scannerCmd.AddCommand(startCmd())
	return scannerCmd
}

var scannerCmd = &cobra.Command{
	Use:   scannerFuncName,
	Short: fmt.Sprint(scannerCmdDes),
	Long:  fmt.Sprint(scannerCmdDes),
}

