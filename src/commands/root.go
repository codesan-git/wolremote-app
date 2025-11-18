package commands

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "wolremote [config file] <command>",
	Short: "WOL Remote CLI to manage remote PCs",
	Long: `wolremote is a CLI tool to control remote PCs:
- Wake-on-LAN
- SSH connect/disconnect
- Run, check and stop remote programs
- Shutdown remote PCs

Example usage:
  wolremote on --remote "config_1.toml" (Turn on remote computer using Wake-On-LAN)
  wolremote off --remote "config_1.toml" (Turn off remote computer)
  wolremote connect --remote "config_1.toml" (Connect to remote computer via SSH)
  wolremote disconnect --remote "config_1.toml" (Disconnect from remote computer)
  wolremote run --remote "config_1.toml" (Run the speciefied program)
  wolremote check --remote "config_1.toml" (Check the speciefied program)
  wolremote stop --remote "config_1.toml" (Stop the specified program)`,
}
