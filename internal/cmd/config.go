package cmd

import (
	"os"
	"os/user"
	"path"

	"github.com/urfave/cli"
)

const (
	appDir = ".snip"
)

var (
	// InitConfigCommand re-initializes app config
	InitConfigCommand = cli.Command{
		Name:     "config",
		Category: "Config",
		Usage:    "configuration options",
		Subcommands: []cli.Command{
			{
				Name:   "init",
				Usage:  "reinitialize the snip configuration",
				Action: initConfig,
			},
		},
	}
)

func initConfig(c *cli.Context) error {

	ok, err := userDirExists()
	if err != nil {
		return err
	}

	if !ok {
		err = os.Mkdir(getUserDirPath(), 0700)
		if err != nil {
			return err
		}
	}

	return nil
}

func getUserDirPath() string {
	usr, _ := user.Current()
	return path.Join(usr.HomeDir, appDir)
}

func userDirExists() (ok bool, err error) {
	s, e := os.Stat(getUserDirPath())
	if e != nil {
		if os.IsNotExist(e) {
			return false, nil
		}
		return false, e
	}
	return s.IsDir(), nil
}
