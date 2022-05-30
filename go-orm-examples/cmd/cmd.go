package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-orm/examples/ent_example"
	"go-orm/examples/sqlboiler_example"
	"go-orm/examples/sqlc_example"
	"go-orm/internal"
)

func init() {
	rootCmd.AddCommand(entCmd)
	rootCmd.AddCommand(sqlcCmd)
	rootCmd.AddCommand(sqlboilerCmd)

	rootCmd.AddCommand(migrateUpCmd)
	rootCmd.AddCommand(migrateDownCmd)

	sqlboilerCmd.AddCommand(sqlboilerSeedCmd)

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

var sqlboilerCmd = &cobra.Command{
	Use:   "sqlboiler",
	Short: "Run ent example",
	Run: func(cmd *cobra.Command, args []string) {
		sqlboiler_example.Run()
	},
}

var sqlboilerSeedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Run sqlboiler seeder",
	Run: func(cmd *cobra.Command, args []string) {
		sqlboiler_example.Seed(10, 100)
	},
}

var sqlcCmd = &cobra.Command{
	Use:   "sqlc",
	Short: "Run ent example",
	Run: func(cmd *cobra.Command, args []string) {
		sqlc_example.Run()
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
	Short: "Generate migration from ent",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migrate")
	},
}

var entMigrateGenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate migration from ent",
	Run: func(cmd *cobra.Command, args []string) {
		ent_example.GenerateMigration()
	},
}

var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "Process migration up",
	Run: func(cmd *cobra.Command, args []string) {
		internal.UpMigration()
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "Process migration down",
	Run: func(cmd *cobra.Command, args []string) {
		internal.DownMigration()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
