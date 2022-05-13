/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"gmctl/cmd/step"

	"github.com/spf13/cobra"
)

// stepCmd represents the step command
var StepCmd = &cobra.Command{
	Use:   "step",
	Short: "流水线阶段功能",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("step called")
	},
}

func init() {
	rootCmd.AddCommand(StepCmd)
	StepCmd.AddCommand(step.GrayCmd)
	StepCmd.AddCommand(step.AutotestCmd)
	StepCmd.AddCommand(step.CheckApproveCmd)
	StepCmd.AddCommand(step.CreateApproveCmd)
	StepCmd.AddCommand(step.DeloyCmd)
	StepCmd.AddCommand(step.PrebuildCmd)
	StepCmd.AddCommand(step.PromoteCmd)
	StepCmd.AddCommand(step.RollbackCmd)
	StepCmd.AddCommand(step.ScanCmd)


	StepCmd.PersistentFlags().StringP("env", "e", "","env")
	StepCmd.PersistentFlags().StringP("name","n","","name")
	StepCmd.PersistentFlags().StringP("pipeline_id","p","","pipeline_id")
	StepCmd.Flags().StringVarP(&ConfigFile, "config", "c", "~/.gmctl/config.yaml","config file ")


	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stepCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stepCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
