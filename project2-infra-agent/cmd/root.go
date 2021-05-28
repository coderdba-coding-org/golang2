package cmd

// Cobra example: https://levelup.gitconnected.com/exploring-go-packages-cobra-fce6c4e331d6

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"os"
        "log"
)

// rootCmd is the base command called when run without any subcommands
// You can define another command and then do rootCmd.AddCommand(theOtherCommand) --> see example
var rootCmd = &cobra.Command{
	Use:   "myagent",
	Short: "Node management",
	Long: `myagent runs on each hypervisor. 
	The agent is used to manage the hypervisor and bootstrap servers
	`,
	Run: func(cmd *cobra.Command, args []string) (err error) {
                fmt.Println("Called the agent")
	},
}

// This is called by main.go
// It checks error of its own execution and exits if errors out
func ExecuteContext(ctx context.Context) {
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}

func init() {
   cobra.OnInitialize(initConfig)
}

func initConfig() {

	config.Config.SetConfigFile(cfgFile)
	if err := config.Config.ReadInConfig(); err != nil {
		log.Warn("Failed to read %v - viper %s", cfgFile, err.Error())
		return
	}
	fmt.Println("Using config file:", config.Config.ConfigFileUsed())
}
