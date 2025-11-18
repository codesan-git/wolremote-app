package commands

import (
	"fmt"
	"wolremote/src/config"
	"wolremote/src/ssh"

	"github.com/spf13/cobra"
)

var ConnectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to remote PC via SSH",
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

		// Successfully connected to SSH
		fmt.Println("SSH connected to", cfgFile)
	},
}

func init() {
	// Add the --remote flag
	ConnectCmd.Flags().String("remote", "", "Path to the remote PC config file")
	ConnectCmd.MarkFlagRequired("remote")

	// Add ConnectCmd to the root command
	RootCmd.AddCommand(ConnectCmd)
}
