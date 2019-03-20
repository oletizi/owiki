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
		server.Run(port)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().IntP("port", "p", 8080, "Port number")
}
