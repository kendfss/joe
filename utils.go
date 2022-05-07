package main

import (
	"io"
	"net/http"
	"os"
	"path"

	"github.com/termie/go-shutil"

	"github.com/kendfss/but"
)

func DownloadFiles(url, dataPath string) (err error) {
	archivePath := path.Join("/tmp", "master.zip")

	// Create the file
	out, err := os.Create(archivePath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	// Unzip
	// err = unzip(archivePath, "/tmp")
	err = unzipSource(archivePath, "/tmp")
	but.Must(err)

	err = shutil.CopyTree(path.Join("/tmp", "gitignore-main"), dataPath, nil)
	but.Must(err)

	return nil
}
