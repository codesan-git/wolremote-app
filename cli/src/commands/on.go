package commands

import (
	"fmt"
	"wolremote/src/config"
	"wolremote/src/wol"

	"github.com/spf13/cobra"
)

var OnCmd = &cobra.Command{
	Use:   "on",
	Short: "Wake remote PC using Wake-on-LAN",
	Run: func(cmd *cobra.Command, args []string) {
		// Get the config file path from the --remote flag
		cfgFile, _ := cmd.Flags().GetString("remote")

		// Load the configuration
		cfg, err := config.LoadConfig(cfgFile)
		if err != nil {
			fmt.Println("Error loading config:", err)
			return
		}

		// Send Wake-on-LAN packet
		err = wol.Wake(cfg.Client.MacAddress, cfg.Client.BroadcastIP)
		if err != nil {
			fmt.Println("Failed to send WOL packet:", err)
			return
		}

		fmt.Println("Wake-on-LAN packet sent for", cfgFile)
	},
}

func init() {
	// Add the --remote flag for passing the config file path
	OnCmd.Flags().String("remote", "", "Path to the remote PC config file")
	OnCmd.MarkFlagRequired("remote") // Make it required

	// Add the OnCmd to the RootCmd
	RootCmd.AddCommand(OnCmd)
}
