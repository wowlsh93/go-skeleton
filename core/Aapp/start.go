/*
Copyright Stick Corp. All Rights Reserved.

Written by HAMA
*/

package Aapp

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wowlsh93/go-skeleton/core/Aapp/config"
	"github.com/wowlsh93/go-skeleton/core/Aapp/minus"
	"github.com/wowlsh93/go-skeleton/core/Aapp/plus"
	"github.com/wowlsh93/go-skeleton/logging"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var configPath string

func startCmd() *cobra.Command {
	scannerStartCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "", "config file (default is ./configurations/bs_config.yaml)")
	scannerStartCmd.MarkPersistentFlagRequired("mode")

	return scannerStartCmd
}

var scannerStartCmd = &cobra.Command{
	Use:   "start",
	Short: "start aapp..",
	Long:  `start aapp..`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return fmt.Errorf("trailing args detected")
		}

		cmd.SilenceUsage = true
		return serve(args)
	},
}

func serve(args []string) error {
	var conf = config.InitConfig(configPath)
	logging.InitLog(&conf)

	cleanUp := make(chan bool)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	//service lifecycle
	for {
		select {
		case <-time.After(5 * time.Second):
			fmt.Println(minus.Minus(5,3))
			fmt.Println(plus.Plus(5,3))
		//case <- resume:
		//case <- suspend:
		//case <- reconfiguring:
		case <-cleanUp:
			fmt.Println("scanner cleanUP was Finished")
			time.Sleep(500 * time.Millisecond)
			return nil
		case sig := <-sigs:
			fmt.Println(sig)
			fmt.Println("scanner closing by signal")
		}
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println("scanner exit")
	return nil
}
