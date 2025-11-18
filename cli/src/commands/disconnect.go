package commands

import (
	"fmt"
	"wolremote/src/config"
	"wolremote/src/ssh"

	"github.com/spf13/cobra"
)

var DisconnectCmd = &cobra.Command{
	Use:   "disconnect",
	Short: "Disconnect SSH from remote PC",
	Run: func(cmd *cobra.Command, args []string) {
		// Get the config file path from the --remote flag
		cfgFile, _ := cmd.Flags().GetString("remote")

		// Load the configuration
		cfg, err := config.LoadConfig(cfgFile)
		if err != nil {
			fmt.Println("Error loading config:", err)
			return
		}

		// Connect via SSH
		client, err := ssh.NewClient(cfg.SSH)
		if err != nil {
			fmt.Println("SSH connection failed:", err)
			return
		}
		defer client.Close()

		// Successfully disconnected (you might want to add custom disconnect handling if needed)
		fmt.Println("SSH disconnected from", cfgFile)
	},
}

func init() {
	// Add the --remote flag
	DisconnectCmd.Flags().String("remote", "", "Path to the remote PC config file")
	DisconnectCmd.MarkFlagRequired("remote")

	// Add DisconnectCmd to the root command
	RootCmd.AddCommand(DisconnectCmd)
}
