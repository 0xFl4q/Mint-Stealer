package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"syscall"
	"unsafe"
)

var (
	systemParametersInfo = syscall.NewLazyDLL("user32.dll").NewProc("SystemParametersInfoW")
)

const (
	SPI_SETDESKWALLPAPER = 0x0014
)

func main() {
	imageURL := "https://wallpapercave.com/wp/wp4817329.png" // Lien vers votre image

	imagePath, err := downloadImage(imageURL)
	if err != nil {
		fmt.Println("Erreur lors du téléchargement de l'image:", err)
		os.Exit(1)
	}

	err = setWallpaper(imagePath)
	if err != nil {
		fmt.Println("Erreur lors du changement du fond d'écran:", err)
		os.Exit(1)
	}

	fmt.Println("Le fond d'écran a été changé avec succès!")
}

func downloadImage(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	tempDir := os.TempDir()
	imagePath := filepath.Join(tempDir, "image.jpg")

	file, err := os.Create(imagePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return "", err
	}

	return imagePath, nil
}

func setWallpaper(imagePath string) error {
	imagePtr, err := syscall.UTF16PtrFromString(imagePath)
	if err != nil {
		return err
	}

	ret, _, err := systemParametersInfo.Call(
		uintptr(SPI_SETDESKWALLPAPER),
		0,
		uintptr(unsafe.Pointer(imagePtr)),
		uintptr(0),
	)
	if ret == 0 {
		return err
	}

	return nil
}
