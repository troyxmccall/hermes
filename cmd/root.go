package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/troyxmccall/hermes/core"
)

var cachePath string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hermes",
	Short: "Takes a yaml file containing commands and bash snippits and executes each command while showing a simple (vertical) progress bar",
	Long:  `Takes a yaml file containing commands and bash snippits and executes each command while showing a simple (vertical) progress bar`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initHermes)
	rootCmd.PersistentFlags().StringVar(&cachePath, "cache-path", "", "The path where cached files will be stored. By default '$(pwd)/.hermes' is used")
}

// initConfigDir ...todo
func initHermes() {
	core.Setup()
}
