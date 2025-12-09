package cmd

import (
	"fmt"
	"os"
	"univ-cli/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Lance l'interface utilisateur textuelle",
	Long: `Lance une interface interactive avec :
  - Affichage de la date et l'heure
  - Message de bienvenue
  - Mini-jeu (deviner un nombre)
  - Sélection de thèmes de couleurs

Navigation : ↑/↓ naviguer • Enter sélectionner • q quitter`,
	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(tui.InitialModel())

		if _, err := p.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Erreur: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}
