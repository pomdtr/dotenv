package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "dotenv [flags] command [args...]",
		Short: "Run a command with env vars injected from one or more .env files",
		Example: `dotenv -- YOUR_COMMAND --YOUR-FLAG
dotenv --command "YOUR_COMMAND && YOUR_OTHER_COMMAND"`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if cmd.Flags().Changed("completion") {
				return nil
			}
			if cmd.Flags().Changed("command") {
				if len(args) > 1 {
					return fmt.Errorf("command and args are mutually exclusive")
				}
			} else {
				if len(args) == 0 {
					return fmt.Errorf("command is required")
				}
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if cmd.Flags().Changed("completion") {
				shell, _ := cmd.Flags().GetString("completion")
				switch shell {
				case "bash":
					return cmd.GenBashCompletionV2(os.Stdout, true)
				case "zsh":
					return cmd.GenZshCompletion(os.Stdout)
				case "fish":
					return cmd.GenFishCompletion(os.Stdout, true)
				case "powershell":
					return cmd.GenPowerShellCompletion(os.Stdout)
				default:
					return fmt.Errorf("unsupported shell: %s", shell)
				}
			}

			envs, _ := cmd.Flags().GetStringSlice("env")
			preserveEnv, _ := cmd.Flags().GetBool("preserve-env")

			if cmd.Flags().Changed("command") {
				command, _ := cmd.Flags().GetString("command")
				shell, ok := os.LookupEnv("SHELL")
				if !ok {
					shell = "/bin/sh"
				}
				args = []string{shell, "-c", command}
			}

			if err := godotenv.Exec(envs, args[0], args[1:], !preserveEnv); err != nil {
				return fmt.Errorf("failed to run command: %w", err)
			}

			return nil
		},
	}

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.Flags().StringSliceP("env", "e", []string{}, "env files to load")
	rootCmd.Flags().Bool("preserve-env", false, "preserve existing environment variables")
	rootCmd.Flags().StringP("command", "c", "", "command to run")
	rootCmd.Flags().String("completion", "", "generate completion script")

	return rootCmd
}

func main() {
	cmd := NewRootCmd()
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
