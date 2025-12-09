package cmd

import (
	"fmt"
	"univ-cli/internal/filesystem"

	"github.com/spf13/cobra"
)

var fsCmd = &cobra.Command{
	Use:   "fs",
	Short: "Manipuler le système de fichiers",
	Long: `La commande fs permet de naviguer et manipuler le système de fichiers local.

Sous-commandes disponibles:
  ls [path]              Liste les fichiers et dossiers
  cat <file>             Affiche le contenu d'un fichier
  cp <file> <dest>       Copie un fichier vers une destination
  mkdir <dir>            Crée un nouveau répertoire`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var lsCmd = &cobra.Command{
	Use:   "ls [path]",
	Short: "Liste les fichiers et dossiers",
	Long: `Liste les fichiers et dossiers dans le répertoire spécifié.
Si aucun chemin n'est fourni, liste le répertoire courant.

Exemples:
  univ-cli fs ls
  univ-cli fs ls /tmp
  univ-cli fs ls ./Documents`,
	Run: func(cmd *cobra.Command, args []string) {
		path := "."
		if len(args) > 0 {
			path = args[0]
		}

		if err := filesystem.ListDirectory(path); err != nil {
			PrintError(err)
			return
		}
	},
}

var catCmd = &cobra.Command{
	Use:   "cat <file>",
	Short: "Affiche le contenu d'un fichier",
	Long: `Affiche le contenu du fichier spécifié.

Exemples:
  univ-cli fs cat fichier.txt
  univ-cli fs cat README.md`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filepath := args[0]

		if err := filesystem.DisplayFileContent(filepath); err != nil {
			PrintError(err)
			return
		}
	},
}

var cpCmd = &cobra.Command{
	Use:   "cp <file> <destination>",
	Short: "Copie un fichier",
	Long: `Copie le fichier source vers la destination spécifiée.
Les permissions du fichier sont préservées.

Exemples:
  univ-cli fs cp source.txt dest.txt
  univ-cli fs cp fichier.txt /tmp/`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		source := args[0]
		destination := args[1]

		if err := filesystem.CopyFile(source, destination); err != nil {
			PrintError(err)
			return
		}

		PrintSuccess(fmt.Sprintf("Fichier copié: %s → %s", source, destination))
	},
}

var mkdirCmd = &cobra.Command{
	Use:   "mkdir <directory>",
	Short: "Crée un nouveau répertoire",
	Long: `Crée un nouveau répertoire avec le nom spécifié.
Les répertoires parents sont créés automatiquement si nécessaire.

Exemples:
  univ-cli fs mkdir nouveau_dossier
  univ-cli fs mkdir /tmp/test/nested`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dirPath := args[0]

		if err := filesystem.CreateDirectory(dirPath); err != nil {
			PrintError(err)
			return
		}

		PrintSuccess(fmt.Sprintf("Répertoire créé: %s", dirPath))
	},
}

func init() {
	rootCmd.AddCommand(fsCmd)

	fsCmd.AddCommand(lsCmd)
	fsCmd.AddCommand(catCmd)
	fsCmd.AddCommand(cpCmd)
	fsCmd.AddCommand(mkdirCmd)
}
