package filesystem

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func DisplayFileContent(filePath string) error {
	info, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("le fichier '%s' n'existe pas", filePath)
		}
		return fmt.Errorf("impossible d'accéder au fichier '%s': %w", filePath, err)
	}

	if info.IsDir() {
		return fmt.Errorf("'%s' est un répertoire, pas un fichier", filePath)
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("impossible de lire le fichier '%s': %w", filePath, err)
	}

	if !utf8.Valid(content) {
		return fmt.Errorf("le fichier '%s' semble être un fichier binaire", filePath)
	}

	fmt.Print(string(content))

	if len(content) > 0 && content[len(content)-1] != '\n' {
		fmt.Println()
	}

	return nil
}
