package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/urfave/cli"

	"github.com/kendfss/but"
)

const joe = `
 ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄ 
▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌
 ▀▀▀▀▀█░█▀▀▀ ▐░█▀▀▀▀▀▀▀█░▌▐░█▀▀▀▀▀▀▀▀▀ 
      ▐░▌    ▐░▌       ▐░▌▐░▌          
      ▐░▌    ▐░▌       ▐░▌▐░█▄▄▄▄▄▄▄▄▄ 
      ▐░▌    ▐░▌       ▐░▌▐░░░░░░░░░░░▌
      ▐░▌    ▐░▌       ▐░▌▐░█▀▀▀▀▀▀▀▀▀ 
      ▐░▌    ▐░▌       ▐░▌▐░▌          
 ▄▄▄▄▄█░▌    ▐░█▄▄▄▄▄▄▄█░▌▐░█▄▄▄▄▄▄▄▄▄ 
▐░░░░░░░▌    ▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌
 ▀▀▀▀▀▀▀      ▀▀▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀▀▀▀▀ 
`

var app = cli.NewApp()

func init() {
	app.Name = joe
	app.Usage = "generate .gitignore files from the command line"
	app.UsageText = "joe command [arguments...]"
	app.Version = version
	// app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name:    "l",
			Aliases: []string{"list"},
			Usage:   "list all available files",
			Action: func(c *cli.Context) error {
				availableGitignores, err := availableFiles()
				if err != nil {
					defer errLogger.Fatal(err)
					return err
				}
				fmt.Printf("%d supported .gitignore files:\n", len(availableGitignores))
				sort.Strings(availableGitignores)
				for _, gnore := range availableGitignores {
					fmt.Println(gnore)
				}
				return nil
			},
		},
		{
			Name:    "u",
			Aliases: []string{"update"},
			Usage:   "update all available gitignore files",
			Action: func(c *cli.Context) error {
				fmt.Println("Updating gitignore files..")
				// err := RemoveContents(dataPath)
				err := os.RemoveAll(dataPath)
				but.Exif(err != nil, err)

				err = DownloadFiles(gitignoreUrl, dataPath)
				if err != nil {
					defer errLogger.Fatal(err)
					return err
				}
				return nil
			},
		},
		{
			Name:    "g",
			Aliases: []string{"generate"},
			Usage:   "generate gitignore files",
			Action: func(c *cli.Context) error {
				if c.NArg() != 1 {
					return cli.ShowAppHelp(c)
				} else {
					generate(c.Args()[0])
				}
				return nil
			},
		},
		{
			Name:    "s",
			Aliases: []string{"search"},
			Usage:   "search for gitignore files (one word per query)",
			Action: func(c *cli.Context) error {
				if c.NArg() != 1 {
					return cli.ShowAppHelp(c)
				} else {
					search(c.Args()[0])
				}
				return nil
			},
		},
	}
}
