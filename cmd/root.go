package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "zvcache",
	Short: "A system-wide cache manager for virtual environment dependencies",
	Long: `zvcache is a daemon that maintains a system-wide cache of dependencies
for virtual environment tools. It helps reduce storage usage and improve file-cache
efficiency in multi-user systems by preventing dependency duplication across
different virtual environments.

When a user creates a new environment and installs a dependency, zvcache first
checks if it exists in the local cache. If available, it adds the dependency
to the user environment as a soft link with no write privileges. If not available,
the cache daemon downloads the dependency into its local cache for future use.

This approach particularly benefits:
- Multi-user systems like clusters and shared workstations
- Environments with network/distributed file systems
- Systems with multiple virtual environments sharing common dependencies`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.zvcache.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
