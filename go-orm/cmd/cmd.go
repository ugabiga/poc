package cmd

import (
	"fmt"
	"go-orm/ent_example"
	"go-orm/internal"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(entCmd)
	rootCmd.AddCommand(migrateUpCmd)
	rootCmd.AddCommand(migrateDownCmd)

	entCmd.AddCommand(entMigrateCmd)

	entMigrateCmd.AddCommand(entMigrateGenerateCmd)
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

var entMigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run ent migration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migrate")
	},
}

var entMigrateGenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate ent migration",
	Run: func(cmd *cobra.Command, args []string) {
		ent_example.GenerateMigration()
	},
}

var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "Process up migration",
	Run: func(cmd *cobra.Command, args []string) {
		internal.UpMigration()
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "Process down migration",
	Run: func(cmd *cobra.Command, args []string) {
		internal.DownMigration()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
