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

// directoryCmd represents the directory command
var directoryCmd = &cobra.Command{
	Use:   "directory",
	Short: "Open a directory selection dialog",
	Long: `Open a native directory selection dialog that allows you to pick a folder.
The path of the selected directory is returned to stdout.`,
	Aliases: []string{"dir", "folder"},
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		startDir, _ := cmd.Flags().GetString("directory")

		// Prepare the dialog
		builder := dialog.Directory().Title(title)
		if startDir != "" {
			builder.SetStartDir(startDir)
		}
		dir, err := builder.Browse()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}

		// Output the selected directory path
		fmt.Println(dir)
	},
}

func init() {
	rootCmd.AddCommand(directoryCmd)

	// Flags for the directory command
	directoryCmd.Flags().StringP("title", "t", "Select a directory", "Set the dialog title")
	directoryCmd.Flags().StringP("directory", "d", "", "Set the starting directory")
}
