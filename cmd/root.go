package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task-cli",
	Short: "Create Jira task and Git branch",
}

func initConfig() {
	viper.SetConfigName(".task-cli")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// No config file found, no problem
		} else {
			fmt.Println("Error reading config:", err)
			os.Exit(1)
		}
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
	}
}
