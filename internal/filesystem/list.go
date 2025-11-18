package filesystem

import (
	"fmt"
	"os"
)

func ListDirectory(path string) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("impossible de lire le répertoire '%s': %w", path, err)
	}

	fmt.Printf("📁 Contenu de: %s\n\n", path)

	if len(entries) == 0 {
		fmt.Println("(vide)")
		return nil
	}

	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Printf("  📂 %s/\n", entry.Name())
		} else {
			fmt.Printf("  📄 %s\n", entry.Name())
		}
	}

	fmt.Printf("\nTotal: %d élément(s)\n", len(entries))

	return nil
}
