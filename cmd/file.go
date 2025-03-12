/*
Copyright © 2025 Hamza Boukhentiche hamza@boukh.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/sqweek/dialog"
)

// fileCmd represents the file command
var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "Open a file selection dialog",
	Long: `Open a native file selection dialog that allows you to pick a file.
The path of the selected file is returned to stdout.`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		directory, _ := cmd.Flags().GetString("directory")
		save, _ := cmd.Flags().GetBool("save")
		filter, _ := cmd.Flags().GetStringSlice("filter")

		// Prepare the dialog
		var builder dialog.FileBuilder
		if title != "" {
			builder.Title(title)
		}
		if directory != "" {
			builder.SetStartDir(directory)
		}
		if len(filter) > 0 {
			for _, f := range filter {
				builder.Filter(f)
			}
		}

		var result string
		var err error

		if save {
			// Save file dialog
			result, err = builder.Save()
		} else {
			// Load file dialog
			result, err = builder.Load()
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}

		// Output the selected file path
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(fileCmd)

	// Flags for the file command
	fileCmd.Flags().StringP("title", "t", "", "Set the dialog title")
	fileCmd.Flags().StringP("directory", "d", "", "Set the starting directory")
	fileCmd.Flags().BoolP("save", "s", false, "Show a save dialog instead of open dialog")
	fileCmd.Flags().StringSliceP("filter", "f", []string{}, "File filters (e.g. 'Images (*.jpg *.png)' 'Documents (*.pdf *.doc)')")
}
