package filesystem

import (
	"fmt"
	"os"
)

func CreateDirectory(path string) error {
	if info, err := os.Stat(path); err == nil {
		if info.IsDir() {
			return fmt.Errorf("le répertoire '%s' existe déjà", path)
		}
		return fmt.Errorf("'%s' existe déjà mais c'est un fichier", path)
	}

	err := os.MkdirAll(path, 0755)
	if err != nil {
		return fmt.Errorf("impossible de créer le répertoire '%s': %w", path, err)
	}

	return nil
}
