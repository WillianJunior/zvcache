/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/WillianJunior/zvcache/envmanager"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		// Optionally run one of the validators provided by cobra
		if err := cobra.ExactArgs(2)(cmd, args); err != nil {
			return err
		}

		// TODO: validate: <env_type> <path> always

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")

		envmng, path := args[0], args[1]

		// not sure which permission to give here
		// os.ModeDir gives error when creating folder
		err := os.Mkdir(path, 0777)
		if err != nil {
			fmt.Printf("error: %s\n", err)
			return
		}

		// readable by everyone, writable by owner
		err = os.WriteFile(path+"/.config", []byte("envmng=venv\n"), 0644)
		if err != nil {
			fmt.Printf("error while creating .config file: %s\n", err)
			return
		}

		em, err := envmanager.NewEnvManager(envmng, path)

		err = em.Create()
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
