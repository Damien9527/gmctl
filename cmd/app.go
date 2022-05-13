/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// appCmd represents the app command
var AppCmd = &cobra.Command{
	Use:   "app",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("app called")
		run()
	},
}

func init() {
	rootCmd.AddCommand(AppCmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// appCmd.PersistentFlags().String("foo", "", "A help for foo")
	AppCmd.PersistentFlags().StringP("env", "e", "","env")
	AppCmd.PersistentFlags().StringP("name","n","","name")
	AppCmd.PersistentFlags().StringP("pipeline_id","p","","pipeline_id")
	AppCmd.Flags().StringVarP(&ConfigFile, "config", "c", "~/.gmctl/config.yaml","config file ")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// appCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	ConfigSetup()
}


func run() {
	fmt.Println("appname: ",os.Getenv("appname"))
	fmt.Println("env: ",os.Getenv("env"))
	fmt.Println("build: ",os.Getenv("BUILD_ID"))
	fmt.Println("build: ",os.Getenv("builder"))
	fmt.Println("Environ: ",os.Environ())

 }