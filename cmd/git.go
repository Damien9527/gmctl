/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gmctl/cmd/git"
)

// gitCmd represents the git command
var GitCmd = &cobra.Command{
	Use:   "git",
	Short: "git subcommands ",
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
	AppCmd.AddCommand(GitCmd)
	GitCmd.AddCommand(git.CloneCmd)
	GitCmd.AddCommand(git.MergeCmd)
	GitCmd.AddCommand(git.TagCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gitCmd.PersistentFlags().String("foo", "", "A help for foo")

	AppCmd.PersistentFlags().StringP("tag", "t", "","tag")
	//appCmd.PersistentFlags().StringP("merge-request","m","","merge-request")
	AppCmd.PersistentFlags().StringP("from","","","from branch ")
	AppCmd.PersistentFlags().StringP("to","","","to branch" )



	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
