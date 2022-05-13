/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package step

import (
	"fmt"
	"github.com/spf13/cobra"
)

// GrayCmd represents the gray command
var GrayCmd = &cobra.Command{
	Use:   "gray",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gray called")
	},
}

func init() {
	//cmd.StepCmd.AddCommand(GrayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// GrayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// GrayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
