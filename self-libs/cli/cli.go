package cli

import (
	"flag"
	"fmt"
	"log"
	"os"

	"roshar/services/helpers/svc"
	"roshar/services/users-svc/pkg/datastore/pg"
	"roshar/services/users-svc/pkg/users"
)

const (
	createIntersts     = "create-interests"
	overwriteInterests = "overwrite-interests"
)

// CLI defines struct for executing commands
type CLI struct {
	*svc.Service

	Users users.Users
}

// Config contains the service configuration
type Config struct {
	AuthDBConnectionStr string
}

type cmdMap map[string]*flag.FlagSet

// New initializes the command line
func New(cfg *Config) CLI {
	db, err := pg.NewDatastore(cfg.AuthDBConnectionStr)
	if err != nil {
		panic(err)
	}

	c := CLI{
		Users:   users.New(db),
		Service: svc.NewService("Users"),
	}
	c.RegisterShutdownFunc(db.Close)

	return c
}

// Run runs the command line
func (cli *CLI) Run() {
	cli.validateArgs()

	cmd := cli.cmd()

	// command parameter values
	createInterestsFile := cmd[createIntersts].String("file", "", "The file to get user interests")
	overwriteInterestsFile := cmd[overwriteInterests].String("file", "", "The file to get user interests")

	cli.parse(cmd)

	if cmd[createIntersts].Parsed() {
		if *createInterestsFile == "" {
			cmd[createIntersts].Usage()
			os.Exit(1)
		}

		err := cli.CreateInterests(*createInterestsFile)
		if err != nil {
			log.Panic(err)
		}
	}

	if cmd[overwriteInterests].Parsed() {
		if *overwriteInterestsFile == "" {
			cmd[overwriteInterests].Usage()
			os.Exit(1)
		}

		err := cli.OverwriteInterests(*overwriteInterestsFile)
		if err != nil {
			log.Panic(err)
		}
	}
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Printf("  %s -file FILENAME - Inserts intersts names from FILENAME which is in JSON format with key \"interests\"\n", createIntersts)
	fmt.Printf("  %s -file FILENAME - Remove all interests from db and insert new interest names FILENAME which is in JSON format with key \"interests\"", overwriteInterests)
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

// cmd returns a map of commands with command string as key and command usage flatset as value
func (cli *CLI) cmd() cmdMap {
	c := cmdMap{}

	cmds := []string{createIntersts, overwriteInterests}
	for _, cmd := range cmds {
		c[cmd] = flag.NewFlagSet(cmd, flag.ExitOnError)
	}

	return c
}

// parse parses flag definitions from the argument list
func (cli *CLI) parse(c cmdMap) {
	cmd, ok := c[os.Args[1]]
	// command is defined, print usage and exit
	if !ok {
		cli.printUsage()
		os.Exit(1)
	}

	err := cmd.Parse(os.Args[2:])
	if err != nil {
		log.Panic(err)
	}
}
