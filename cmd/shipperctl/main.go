package main

import (
	"flag"
	"fmt"
	"github.com/bookingcom/shipper/cmd/shipperctl/cmd/chart"
	"github.com/bookingcom/shipper/cmd/shipperctl/cmd/clean"
	"github.com/bookingcom/shipper/cmd/shipperctl/cmd/clusters"
	"github.com/bookingcom/shipper/cmd/shipperctl/cmd/list"
	"os"

	"github.com/bookingcom/shipper/cmd/shipperctl/cmd/backup"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "shipperctl [command]",
	Short:         "Command line application to make working with Shipper more enjoyable",
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	rootCmd.AddCommand(clusters.Cmd)
	rootCmd.AddCommand(list.Cmd)
	rootCmd.AddCommand(clean.Cmd)
	rootCmd.AddCommand(backup.Cmd)
	rootCmd.AddCommand(chart.Command)
}

func main() {
	flag.Parse()
	err := rootCmd.Execute()
	if err != nil {
		fmt.Printf("Error! %s\n", err)
		os.Exit(1)
	}
}
