/*
Copyright © 2023 helm-variant-packager

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "helm-variant-packager",
	Short: "A helm plugin to compose variant helm charts",
	Long: `A helm plugin which will create helm charts based om
	the various values.*.yaml files in the chart directory. 
	These will wrap the base chart and provide a way to
	deploy the same chart with different values baked in.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//Run: func(cmd *cobra.Command, args []string) {
	//	core.IndentYAMLWithFoobar("foo.yaml", "bar.yaml")
	//
	//},
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
}
