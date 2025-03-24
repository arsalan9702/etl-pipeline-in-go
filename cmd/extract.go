/*
Copyright © 2025 NAME HERE <arsalansayed9702@gmail.com>
*/
package cmd

import (
	"fmt"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
	"log"
	"time"
)

// extractCmd represents the extract command
var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "Extract data from various sources",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Extracting data...")

		// Simulate data extraction (e.g., fetching 100 records)
		totalRecords := 100
		bar := progressbar.Default(int64(totalRecords))

		for i := 0; i < totalRecords; i++ {
			time.Sleep(50 * time.Millisecond)  // Simulating network delay
			if err := bar.Add(1); err != nil { // ✅ Checking for errors
				log.Fatalf("Failed to update progress bar: %v", err)
			}
		}

		fmt.Println("\nData extraction completed successfully!")
	},
}

func init() {
	rootCmd.AddCommand(extractCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// extractCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// extractCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
