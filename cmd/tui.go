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
	Long: `Lance une interface utilisateur textuelle (TUI) interactive avec un menu proposant :
  - Affichage de la date et l'heure
  - Message de bienvenue
  - Mini-jeu : Deviner un nombre
  - Quitter

Navigation :
  ↑/↓ ou j/k : Naviguer dans le menu
  Enter       : Sélectionner une option
  q ou Esc    : Retour au menu / Quitter`,
	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(tui.InitialModel())

		if _, err := p.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Erreur lors du lancement du TUI: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}
