package cmd

import (
	"fmt"
	ent_example "go-orm/ent-example"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(entCmd)
}

var rootCmd = &cobra.Command{
	Use:   "run",
	Short: "Run example",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root command")
	},
}

var entCmd = &cobra.Command{
	Use:   "ent",
	Short: "Run ent example",
	Run: func(cmd *cobra.Command, args []string) {
		ent_example.Run()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
