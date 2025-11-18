package cmd

import (
	"fmt"
	"univ-cli/internal/custom"

	"github.com/spf13/cobra"
)

var customCmd = &cobra.Command{
	Use:   "custom",
	Short: "Commandes personnalisées",
	Long: `La commande custom offre des fonctionnalités personnalisées supplémentaires.

Sous-commandes disponibles:
  quote              Affiche une citation aléatoire
  sysinfo            Affiche les informations système
  calc <expression>  Calcule une expression mathématique simple`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var quoteCmd = &cobra.Command{
	Use:   "quote",
	Short: "Affiche une citation aléatoire",
	Long: `Affiche une citation inspirante aléatoire.

Exemples:
  univ-cli custom quote`,
	Run: func(cmd *cobra.Command, args []string) {
		quote := custom.GetRandomQuote()
		PrintInfo(fmt.Sprintf("💭 %s", quote))
	},
}

var sysinfoCmd = &cobra.Command{
	Use:   "sysinfo",
	Short: "Affiche les informations système",
	Long: `Affiche des informations sur le système d'exploitation,
l'architecture et d'autres détails techniques.

Exemples:
  univ-cli custom sysinfo`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := custom.DisplaySystemInfo(); err != nil {
			PrintError(err)
			return
		}
	},
}

var calcCmd = &cobra.Command{
	Use:   "calc <expression>",
	Short: "Calcule une expression mathématique simple",
	Long: `Calcule une expression mathématique simple (addition, soustraction, multiplication, division).

Exemples:
  univ-cli custom calc "10 + 5"
  univ-cli custom calc "42 * 2"
  univ-cli custom calc "100 / 4"
  univ-cli custom calc "50 - 8"`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		expression := args[0]

		result, err := custom.Calculate(expression)
		if err != nil {
			PrintError(err)
			return
		}

		PrintSuccess(fmt.Sprintf("%s = %.2f", expression, result))
	},
}

func init() {
	rootCmd.AddCommand(customCmd)

	customCmd.AddCommand(quoteCmd)
	customCmd.AddCommand(sysinfoCmd)
	customCmd.AddCommand(calcCmd)
}
