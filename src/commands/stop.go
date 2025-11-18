package commands

import (
	"fmt"
	"wolremote/src/checker" // Import the checker package to check if the program is running
	"wolremote/src/config"
	"wolremote/src/ssh"

	"github.com/spf13/cobra"
)

var StopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the remote program",
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

		// Check if the program is running on the remote machine
		running, err := checker.CheckIfRunning(client, cfg.RemoteProgram.ProgramName)
		if err != nil {
			fmt.Println("Error checking if program is running:", err)
			return
		}

		// If the program is running, attempt to stop it
		if running {
			// Prepare the command to stop the program using pkill
			remoteCmd := fmt.Sprintf("pkill -f %s", cfg.RemoteProgram.ProgramName)
			if cfg.RemoteProgram.SudoUser {
				remoteCmd = "sudo " + remoteCmd
			}

			// Run the stop command
			_, err := client.RunCommand(remoteCmd)
			if err != nil {
				fmt.Println("[INFO] Cannot stop the program:", err)
				return
			}

			// Output success message
			fmt.Println("Remote program stopped:", cfg.RemoteProgram.ProgramName)
			fmt.Println("[INFO] Program successfully stopped.")
		} else {
			// If the program is not running, output that the program is not found
			fmt.Println("Remote program stopped:", cfg.RemoteProgram.ProgramName)
			fmt.Println("[INFO] Program not found.")
		}
	},
}

func init() {
	// Add the --remote flag
	StopCmd.Flags().String("remote", "", "Path to the remote PC config file")
	StopCmd.MarkFlagRequired("remote")

	// Add StopCmd to the root command
	RootCmd.AddCommand(StopCmd)
}
