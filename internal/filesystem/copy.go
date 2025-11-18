package filesystem

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CopyFile(source, destination string) error {
	sourceInfo, err := os.Stat(source)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("le fichier source '%s' n'existe pas", source)
		}
		return fmt.Errorf("impossible d'accéder au fichier source '%s': %w", source, err)
	}

	if sourceInfo.IsDir() {
		return fmt.Errorf("'%s' est un répertoire (seuls les fichiers peuvent être copiés)", source)
	}

	destPath := destination
	destInfo, err := os.Stat(destination)
	if err == nil && destInfo.IsDir() {
		destPath = filepath.Join(destination, filepath.Base(source))
	}

	if _, err := os.Stat(destPath); err == nil {
		return fmt.Errorf("le fichier de destination '%s' existe déjà", destPath)
	}

	sourceFile, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("impossible d'ouvrir le fichier source '%s': %w", source, err)
	}
	defer sourceFile.Close()

	destFile, err := os.OpenFile(destPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, sourceInfo.Mode())
	if err != nil {
		return fmt.Errorf("impossible de créer le fichier de destination '%s': %w", destPath, err)
	}
	defer destFile.Close()

	bytesWritten, err := io.Copy(destFile, sourceFile)
	if err != nil {
		os.Remove(destPath)
		return fmt.Errorf("erreur lors de la copie: %w", err)
	}

	if bytesWritten != sourceInfo.Size() {
		os.Remove(destPath)
		return fmt.Errorf("copie incomplète")
	}

	if err := destFile.Sync(); err != nil {
		return fmt.Errorf("erreur lors de la synchronisation: %w", err)
	}

	return nil
}
