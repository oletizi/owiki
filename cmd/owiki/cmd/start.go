// Copyright © 2019 NAME HERE <EMAIL ADDRESS>

package cmd

import (
	"github.com/oletizi/owiki/internal/server"
	"log"

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
		docroot, _ := cmd.Flags().GetString("docroot")
		templateDir, _ := cmd.Flags().GetString("template-dir")
		log.Printf("port: %i", port)
		log.Printf("data-dir: %s", dataDir)
		log.Printf("docroot: %s", docroot)

		server.Run(port, dataDir, docroot, templateDir)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().IntP("port", "p", 8080, "Port number")
	startCmd.Flags().StringP("data-dir", "d", "/tmp", "Directory to store page data")
	startCmd.Flags().String("docroot", ".", "Web document root.")
	startCmd.Flags().StringP("template-dir", "t", ".", "Template directory")
}
