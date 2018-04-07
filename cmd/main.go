package cmd

import (
	"github.com/urfave/cli"

	"os"
)



func main() {
	app := cli.NewApp()
	app.Version = Version
	app.Usage = "An URL Shortener for Go."
	app.Commands = []cli.Command{
		//TODO:
	}
	app.Run(os.Args)
}