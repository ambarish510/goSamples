TO create CLI from golang, use the package "github.com/urfave/cli"
In the main file, mention the parameters like :-
  app.Name 
  app.Usage
  app.EnableBashCompletion
  app.Version
  app.Before
  app.After
  app.Flags
  app.Action
  app.Commands
  app.CommandNotFound
  
The app.Action defines is the global options (similar to --help / --version). The commands has a higher order of precedence. 
If commands are issued then the action part of the command is executed and not the app.action part. 
If command is not issued, the global action is executed. 
If a global action is mentioned, the app.CommandNotFound is never executed.
