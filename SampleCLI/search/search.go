package search

import (
	"github.com/urfave/cli"
	"fmt"
  "strings"
)

func GetSearchFlags() []cli.Flag {
	flags := []cli.Flag{
		cli.StringFlag{
    	Name:  "store",
			Usage: "Expected values are Maven or pypi eg: --store Maven or --store pypi",
		},cli.StringFlag{
			Name:  "packagename",
			Usage: "Maven use the format groupid:artifactid  eg: --packagename com.google.inject:guice, For other just the packagename eg: -packagename Flask ",
		},
	}
	return flags
}

func SearchArtifactory(c *cli.Context) error{
	ArtStore := strings.ToUpper(c.String("store"))
	ArtPackageName := c.String("packagename")
  fmt.Println("SearchArtifactory\nStore: ",ArtStore,"Package Name: ",ArtPackageName)
	return nil
}
