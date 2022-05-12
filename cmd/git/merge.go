/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package git

import (
	"fmt"
	"github.com/spf13/cobra"
)

// mergeCmd represents the merge command
var MergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("merge called")
	},
}

func init() {
	//cmd.gitCmd.AddCommand(mergeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mergeCmd.PersistentFlags().String("foo", "", "A help for foo")
	MergeCmd.PersistentFlags().StringP("from","","","from branch ")
	MergeCmd.PersistentFlags().StringP("to","","","to branch" )
	MergeCmd.PersistentFlags().StringP("auto","","","if auto " )

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mergeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}