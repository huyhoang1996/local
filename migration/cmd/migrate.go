// Copyright © 2020 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/local/migration/migration"
	"github.com/spf13/cobra"
)

var cmdMigrate string

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: Work your own magic here
		// env := os.Getenv("ENV")
		// logLevel := os.Getenv("LOG_LEVEL")
		// var log logger.Logger
		// if env == "local" {
		// 	log = logger.Logger{Stdout: true, Level: logLevel}
		// } else {
		// 	log = logger.Logger{Stdout: false, Level: logLevel, OutputFile: "logs/migration.log"}
		// }

		// util.Log = log.NewLogger()
		// fmt.Println("===== hoang")
		a, err := migration.New()
		if err != nil {
			return err
		}
		err = a.Migration(cmdMigrate)
		if err != nil {
			return err
		}
		fmt.Println("migrate called")
		return nil
	},
}

func init() {
	RootCmd.PersistentFlags().StringVar(&cmdMigrate, "m", "init", "commands for migration")

	RootCmd.AddCommand(migrateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
