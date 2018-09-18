

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/troyxmccall/hermes/core"
	"path/filepath"
)

// bundleCmd represents the bundle command
var bundleCmd = &cobra.Command{
	Use:   "bundle",
	Short: "Bundle yaml and referenced resources into a single executable (experimental)",
	Long:  `Bundle yaml and referenced resources into a single executable (experimental)`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		userYamlPath := args[0]
		bundlePath := filepath.Base(userYamlPath[0:len(userYamlPath)-len(filepath.Ext(userYamlPath))]) + ".hermes"

		core.Bundle(userYamlPath, bundlePath)
	},
}

func init() {
	rootCmd.AddCommand(bundleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bundleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bundleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}