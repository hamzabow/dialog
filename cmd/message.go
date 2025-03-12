/*
Copyright © 2025 Hamza Boukhentiche hamza@boukh.com
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/sqweek/dialog"
)

// messageCmd represents the message command
var messageCmd = &cobra.Command{
	Use:   "message [text]",
	Short: "Show a message box dialog",
	Long: `Show a native message box dialog with customizable icons and buttons.
The result (which button was pressed) is returned to stdout.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		message := strings.Join(args, " ")
		title, _ := cmd.Flags().GetString("title")
		msgType, _ := cmd.Flags().GetString("type")

		switch strings.ToLower(msgType) {
		case "info":
			dialog.Message("%s", message).Title(title).Info()
		case "error":
			dialog.Message("%s", message).Title(title).Error()
		default:
			dialog.Message("%s", message).Title(title).Info()
		}
	},
}

// confirmCmd represents the confirm command
var confirmCmd = &cobra.Command{
	Use:   "confirm [question]",
	Short: "Show a yes/no confirmation dialog",
	Long: `Show a native confirmation dialog with Yes/No buttons.
Returns 0 (success) if Yes was clicked, and 1 (error) if No was clicked or dialog was cancelled.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		question := strings.Join(args, " ")
		title, _ := cmd.Flags().GetString("title")

		if dialog.Message("%s", question).Title(title).YesNo() {
			fmt.Println("yes")
			os.Exit(0)
		} else {
			fmt.Println("no")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(messageCmd)
	rootCmd.AddCommand(confirmCmd)

	// Flags for the message command
	messageCmd.Flags().StringP("title", "t", "Message", "Set the dialog title")
	messageCmd.Flags().StringP("type", "y", "info", "Message type: 'info' or 'error'")

	// Flags for the confirm command
	confirmCmd.Flags().StringP("title", "t", "Confirmation", "Set the dialog title")
}
