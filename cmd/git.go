/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"gmctl/cmd/git"

	"github.com/spf13/cobra"
)

// gitCmd represents the git command
var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("git called")
	},
}

func init() {
	appCmd.AddCommand(gitCmd)
	gitCmd.AddCommand(git.TagCmd)
	gitCmd.AddCommand(git.MergeCmd)
	gitCmd.AddCommand(git.CloneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gitCmd.PersistentFlags().String("foo", "", "A help for foo")

	appCmd.PersistentFlags().StringP("tag", "t", "","tag")
	//appCmd.PersistentFlags().StringP("merge-request","m","","merge-request")
	appCmd.PersistentFlags().StringP("from","","","from branch ")
	appCmd.PersistentFlags().StringP("to","","","to branch" )



	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
