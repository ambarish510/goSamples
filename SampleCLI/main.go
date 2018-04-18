package main
import (
  "fmt"
  "os"
  "github.com/ambsflip/goSamples/SampleCLI/search"
  "github.com/urfave/cli"
)

func main() {
  app := cli.NewApp()
	app.Name = "Sample CLI"
	app.Usage = "To understand creating CLI in GOlang "
	app.EnableBashCompletion = true
	app.Version = "v1.0"

  app.Before = func(c *cli.Context) error{
		fmt.Println("Run before the cli")
    //code to setup log files
    //code to fetch from config
    //code to check for version update incase of auto-update enabled
    return nil
	}
  app.After = func(c *cli.Context) error{
    fmt.Println("Run after the cli")
    //close the log fileDescriptor
    return nil
  }
  app.Flags = []cli.Flag {
    cli.StringFlag{
      Name: "lang",
      Value: "english",
      Usage: "language for the greeting",
    },
    cli.StringFlag{
      Name: "country",
      Usage: "language for the country",
    },
  }
  app.Action = func(c *cli.Context) error {
    lang := "defaultLang"
    if c.String("lang") == "" {
      fmt.Println("missing flag lang")
    }else{
      lang = c.String("lang")
      fmt.Println("Language:", lang)
    }

    country := "defaultCountry"
    if c.String("country") == "" {
      fmt.Println("missing flag country")
    }else{
      country = c.String("country")
      fmt.Println("Country:", country)
    }
    name := "defaultArg"
    if c.NArg() > 0 {
      name = c.Args().Get(0)
    }
    fmt.Println("Command Line argument :", name)

    return nil
  }
  app.Commands = getCommands()
  app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Printf("%s: '%s' is not a valid command. See '%s --help'.", c.App.Name, command, os.Args[0])
		//prints the stack trace to log file
		//utils.PrintStackTraceToLogFile()
		//log.Fatalf("%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, os.Args[0])
	}
  //os.Args fetches the runtime/command line arguments
  app.Run(os.Args)
}
func getCommands() []cli.Command {
	commands := []cli.Command{
		{
			Name:    "search",
			Aliases: []string{"S"},
			Usage:   "Used to search a single artifact in the ART repo",
			Flags:   search.GetSearchFlags(),
			Action:  search.SearchArtifactory,
		},
	}
	return commands
}
/*
sample outputs
---------------------------
Ambarish-Mac:bin ambarish.a$ SampleCLI  abcd efgh
Run before the cli
Language: english
missing flag country
Command Line argument : abcd
Run after the cli

the defaults values mentioned in the apps.flag section takes effect for flag lang

app.Action defines is the global options (similar to --help / --version)
COMMANDS:
     search, S  Used to search a single artifact in the ART repo
     help, h    Shows a list of commands or help for one command
GLOBAL OPTIONS:
   --lang value     language for the greeting (default: "english")
   --country value  language for the country
   --help, -h       show help
   --version, -v    print the version

The commands has a higher order of precedence. If commands are issued then the action
part of the command is executed and not the app.action part.

Ambarish-Mac:bin ambarish.a$ SampleCLI --country england search --packagename qwert --store asdf
Run before the cli
SearchArtifactory
Store:  ASDF Package Name:  qwert
Run after the cli

If command is not issued, the global action is executed

Ambarish-Mac:bin ambarish.a$ SampleCLI --country england
Run before the cli
Language: english
Country: england
Command Line argument : defaultArg
Run after the cli

Ambarish-Mac:bin ambarish.a$ SampleCLI abcd
Run before the cli
Language: english
missing flag country
Command Line argument : abcd
Run after the cli

if a global action is mentioned, the app.CommandNotFound is never executed

*/
