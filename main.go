package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/kendfss/but"
)

const (
	gitignoreUrl = "https://github.com/github/gitignore/archive/master.zip"
	dataDir      = ".joe-data"
)

var (
	userHome, dataPath, version string

	errLogger = log.New(os.Stderr, "", 0)
)

func init() {
	var err error
	userHome, err = os.UserHomeDir()
	but.Exif(err != nil, err)

	dataPath = path.Join(userHome, dataDir)
}

func findGitignores() (a map[string]string, err error) {
	_, err = os.ReadDir(dataPath)
	if err != nil {
		return nil, err
	}

	filelist := make(map[string]string)
	err = filepath.Walk(dataPath, func(filepath string, info os.FileInfo, err error) error {
		but.Exif(err != nil, err)
		if strings.HasSuffix(info.Name(), ".gitignore") {
			name := strings.ToLower(strings.Replace(info.Name(), ".gitignore", "", 1))
			filelist[name] = filepath
		}
		return nil
	})
	return filelist, err
}

func availableFiles() (a []string, err error) {
	gitignores, err := findGitignores()
	if err != nil {
		return nil, err
	}

	availableGitignores := []string{".txt"}
	for key := range gitignores {
		availableGitignores = append(availableGitignores, key)
	}

	return availableGitignores, nil
}

func search(arg string) {
	gitignores, err := findGitignores()
	but.Exif(err != nil, err)

	for ig := range gitignores {
		b, err := regexp.MatchString(arg, ig)
		if err != nil {
			errLogger.Println(err)
		} else if b {
			fmt.Println(ig)
		}
	}
}

func generate(args string) {
	names := strings.Split(args, ",")

	gitignores, err := findGitignores()
	but.Exif(err != nil, err)

	notFound := []string{}
	// output := ".DS_Store\n._*\n"
	output := ""
	for _, name := range names {
		if filepath, ok := gitignores[strings.ToLower(name)]; ok {
			bytes, err := ioutil.ReadFile(filepath)
			if err == nil {
				output += "#### " + name + " ####\n"
				output += string(bytes)
				output += "\n"
			}
		} else {
			notFound = append(notFound, name)
		}
	}

	if len(notFound) > 0 {
		errLogger.Printf("Unsupported files: %s\n", strings.Join(notFound, ", "))
		errLogger.Fatal("Run `joe ls` to see list of available gitignores.")
	}

	fmt.Println(output)
}

func main() {
	err := app.Run(os.Args)
	but.Exif(err != nil, err)
}
