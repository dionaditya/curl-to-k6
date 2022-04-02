/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/dionaditya/curl-to-k6/internal"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var initCommand = &cobra.Command{
	Use:   "init",
	Short: "Generator to create k6 script from curl command",
	Long:  `ClI apps to generate k6 script from curl command using gherkin syntax`,

	Run: func(cmd *cobra.Command, args []string) {
		internal.Run(args[0])
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	var rootCmd = &cobra.Command{Use: "app"}

	rootCmd.AddCommand(initCommand)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cmd.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}
