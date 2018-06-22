package main

import (
	"os"
  "log"
  "github.com/urfave/cli"
  "errors"
)

func main() {
  var secret_path string
	app := cli.NewApp()

  app.Flags = []cli.Flag {
    cli.StringFlag{
      Name:        "secret, s",
      Usage:       "Path to your secrets entry",
      Destination: &secret_path,
    },
  }

  app.Action = func(c *cli.Context) error {
    if len(secret_path) == 0 {
      return errors.New("Must specify secret with `-s`")
    }
    // c.Args() contains [script, to, run]
    RunScript(secret_path, c.Args())
    return nil
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
