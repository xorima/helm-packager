/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.xom/xorima/helm-variant-packager/internal/config"
	"github.xom/xorima/helm-variant-packager/internal/core"
	"github.xom/xorima/helm-variant-packager/internal/internallogger"
	"os"
)

// packageCmd represents the package command
var packageCmd = &cobra.Command{
	Use:   "package",
	Short: "Packages all the variant charts",
	Long: `package which will create helm charts based on
	the various values.*.yaml files in the chart-path directory. 
	These will wrap the base chart and be outputted as packages in
	the output-path dirtectory.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("package called")
		err := config.DefaultConfig.Validate()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		log := internallogger.NewLogger()
		handler := core.NewHandler(log, config.DefaultConfig)
		handler.Handle()

	},
}

func init() {
	rootCmd.AddCommand(packageCmd)

	packageCmd.Flags().StringVarP(&config.DefaultConfig.ChartPath, "chart-path", "p", "", "The path for the charts")
	packageCmd.Flags().StringVarP(&config.DefaultConfig.OutputPath, "output-path", "o", "", "The path for the written certificates")

}
