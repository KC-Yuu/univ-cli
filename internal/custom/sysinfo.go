package custom

import (
	"fmt"
	"os"
	"runtime"
)

func DisplaySystemInfo() error {
	fmt.Println("🖥️  Informations Système")
	fmt.Println("=======================")
	fmt.Println()

	fmt.Printf("OS              : %s\n", runtime.GOOS)

	fmt.Printf("Architecture    : %s\n", runtime.GOARCH)

	fmt.Printf("Nombre de CPUs  : %d\n", runtime.NumCPU())

	fmt.Printf("Version de Go   : %s\n", runtime.Version())

	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("impossible d'obtenir le répertoire courant: %w", err)
	}
	fmt.Printf("Répertoire      : %s\n", cwd)

	fmt.Println()
	fmt.Println("Variables d'environnement :")
	fmt.Printf("  HOME          : %s\n", os.Getenv("HOME"))
	fmt.Printf("  USER          : %s\n", os.Getenv("USER"))
	fmt.Printf("  SHELL         : %s\n", os.Getenv("SHELL"))
	fmt.Printf("  PATH          : %s\n", truncateString(os.Getenv("PATH"), 60))

	fmt.Println()
	fmt.Println("Statistiques mémoire Go :")
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("  Alloc         : %d KB\n", m.Alloc/1024)
	fmt.Printf("  TotalAlloc    : %d KB\n", m.TotalAlloc/1024)
	fmt.Printf("  Sys           : %d KB\n", m.Sys/1024)
	fmt.Printf("  NumGC         : %d\n", m.NumGC)

	return nil
}

func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}
