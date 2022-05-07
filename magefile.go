//go:build mage

package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/magefile/mage/mg"
	sh "github.com/magefile/mage/sh"
	"log"
	"os"
	"path"
	"time"
)

var (
	buildDir         string
	buildBinary      string
	buildBinaryLocal string
	buildTime        string
	nuclinoApiKey    string
)

var Default = Debugrun

// build the binary locally
func LocalBuild() error {

	mg.Deps(Prepare)

	buildPath := path.Join(buildDir, buildBinaryLocal)

	fmt.Printf("Building locally %s...\n", buildPath)

	return sh.RunWith(map[string]string{"GOOS": "windows"}, "go", "build", "-ldflags", "-s -X main.buildTime="+buildTime, "-o", buildPath, "./cmd/nuclino-cli/")

}

// building and running locally
func Debugrun() error {

	mg.Deps(Prepare)

	mg.Deps(LocalBuild)

	buildPath := path.Join(buildDir, buildBinaryLocal)

	fmt.Printf("Running locally %s...\n", buildPath)

	err := sh.Run(buildPath, "-env", "development", "-apikey", nuclinoApiKey)

	return err
}

func LoadEnvironment() {
	fmt.Println("Loading environment variables...")
	buildDir = os.Getenv("BUILD_DIR")
	buildBinary = os.Getenv("BUILD_BINARY")
	buildBinaryLocal = os.Getenv("BUILD_BINARY_LOCAL")
	nuclinoApiKey = os.Getenv("NUCLINO_API_KEY")
	buildTime = time.Now().Format("2006-01-02_15:04:05")
}

// Prepare directory for builds
func Prepare() {
	mg.Deps(LoadEnvironment)
	fmt.Printf("Prepare %s directory...\n", buildDir)
	if err := os.Mkdir(buildDir, os.ModePerm); err != nil {
		log.Printf("Creating %s directory didn't work: ", buildDir, err.Error())
	}
}

// Clean up after yourself
func Clean() {
	mg.Deps(LoadEnvironment)
	fmt.Println("Cleaning...")
	os.RemoveAll(buildDir)
}
