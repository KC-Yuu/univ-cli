package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

// Commande custom : fonctionnalités personnalisées
var customCmd = &cobra.Command{
	Use:   "custom",
	Short: "Commandes personnalisées",
	Long: `La commande custom offre des fonctionnalités personnalisées.

Sous-commandes disponibles:
  sysinfo    Affiche les informations système`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Sous-commande sysinfo : affiche les informations système
var sysinfoCmd = &cobra.Command{
	Use:   "sysinfo",
	Short: "Affiche les informations système",
	Long: `Affiche des informations sur le système :
- OS et architecture
- Nombre de CPUs
- Version de Go
- Répertoire courant
- Variables d'environnement

Exemple:
  univ-cli custom sysinfo`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🖥️  Informations Système")
		fmt.Println("========================")
		fmt.Println()

		fmt.Printf("OS           : %s\n", runtime.GOOS)
		fmt.Printf("Architecture : %s\n", runtime.GOARCH)
		fmt.Printf("CPUs         : %d\n", runtime.NumCPU())
		fmt.Printf("Go version   : %s\n", runtime.Version())

		if cwd, err := os.Getwd(); err == nil {
			fmt.Printf("Répertoire   : %s\n", cwd)
		}

		fmt.Println()
		fmt.Println("Environnement :")
		fmt.Printf("  HOME  : %s\n", os.Getenv("HOME"))
		fmt.Printf("  USER  : %s\n", os.Getenv("USER"))
		fmt.Printf("  SHELL : %s\n", os.Getenv("SHELL"))
	},
}

func init() {
	rootCmd.AddCommand(customCmd)
	customCmd.AddCommand(sysinfoCmd)
}
