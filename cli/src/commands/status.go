package commands

import (
	"fmt"
	"wolremote/src/config"
	"wolremote/src/ssh"

	"github.com/spf13/cobra"
)

// StatusCmd checks if the remote PC is on and prints the IP address
var StatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check if the remote PC is ON and reachable",
	Run: func(cmd *cobra.Command, args []string) {
		// Get the config file path from the --remote flag
		cfgFile, _ := cmd.Flags().GetString("remote")

		// Load the configuration
		cfg, err := config.LoadConfig(cfgFile)
		if err != nil {
			fmt.Println("Error loading config:", err)
			return
		}

		// Try to connect to the remote machine via SSH
		client, err := ssh.NewClient(cfg.SSH)
		if err != nil {
			fmt.Println("Remote PC is OFF or unreachable.")
			return
		}
		defer client.Close()

		// If we successfully connect, we know the remote PC is on
		fmt.Printf("Remote PC is ON and reachable. IP: %s\n", cfg.SSH.Host)
	},
}

func init() {
	// Add the --remote flag
	StatusCmd.Flags().String("remote", "", "Path to the remote PC config file")
	StatusCmd.MarkFlagRequired("remote")

	// Add StatusCmd to the root command
	RootCmd.AddCommand(StatusCmd)
}
