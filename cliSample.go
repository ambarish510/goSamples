package main
import (
  "fmt"
  "os"

  "github.com/urfave/cli"
)

func cliSample() {
  app := cli.NewApp()
  app.Before = func(c *cli.Context) error{
		fmt.Println("Run before the cli")
    return nil
	}
  app.After = func(c *cli.Context) error{
    fmt.Println("Run after the cli")
    return nil
  }
  app.Flags = []cli.Flag {
    cli.StringFlag{
      Name: "lang",
      Value: "english",
      Usage: "language for the greeting",
    },
  }
  app.Action = func(c *cli.Context) error {
    name := "SampleName"
    if c.NArg() > 0 {
      name = c.Args().Get(0)
    }
    if c.String("lang") == "spanish" {
      fmt.Println("Hola", name)
    } else {
      fmt.Println("Hello", name)
    }
    return nil
  }
  app.Run(os.Args)
}
