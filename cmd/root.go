package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "univ-cli",
	Short: "Une application CLI pour manipuler le système de fichiers et plus",
	Long: `univ-cli est une application en ligne de commande développée en Go qui permet de :
- Naviguer et manipuler le système de fichiers (fs)
- Lancer une interface utilisateur textuelle (tui)
- Utiliser des fonctionnalités personnalisées (custom)

Utilisez les sous-commandes pour accéder aux différentes fonctionnalités.`,

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func PrintError(err error) {
	fmt.Fprintf(os.Stderr, "\033[31mErreur:\033[0m %v\n", err)
}

func PrintSuccess(message string) {
	fmt.Printf("\033[32m✓\033[0m %s\n", message)
}
func PrintInfo(message string) {
	fmt.Printf("\033[34mℹ\033[0m %s\n", message)
}
