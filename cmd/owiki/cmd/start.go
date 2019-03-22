// Copyright © 2019 NAME HERE <EMAIL ADDRESS>

package cmd

import (
	"github.com/oletizi/owiki/internal/server"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start [port]",
	Short: "Starts the web service",
	Long:  `Starts the web service on port 8080 or the specified port`,
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetInt("port")
		dataDir, _ := cmd.Flags().GetString("data-dir")
		server.Run(port, dataDir)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().IntP("port", "p", 8080, "Port number")
	startCmd.Flags().String("data-dir", "d", "Directory to store page data")
}
