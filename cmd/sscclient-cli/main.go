package main

import (
	"fmt"
	"os"
	"os/user"
	"path"

	"github.com/adiclepcea/socialservicechain/sscclient"
	"github.com/hyperledger/sawtooth-sdk-go/logging"
	"github.com/jessevdk/go-flags"
)

// Command is the base for all subcommands
type Command interface {
	Register(*flags.Command) error
	Name() string
	KeyfilePassed() string
	URLPassed() string
	Run() error
}

//Opts holds the options to show at the cli
type Opts struct {
	Verbose []bool `short:"v" long:"verbose" description:"Enable more verbose output"`
	Version bool   `short:"V" long:"version" description:"Display version information"`
}

var distributionVersion string

var logger = logging.Get()

func init() {
	if len(distributionVersion) == 0 {
		distributionVersion = "Unknown"
	}
}

func main() {
	arguments := os.Args[1:]
	for _, arg := range arguments {
		if arg == "-V" || arg == "--version" {
			fmt.Println(sscclient.DistributionName + " (Hyperledger Sawtooth) version " + distributionVersion)
			os.Exit(0)
		}
	}

	var opts Opts
	parser := flags.NewParser(&opts, flags.Default)
	parser.Command.Name = "ssc"

	// Add sub-commands
	commands := []Command{
		&CreateNGO{},
		&CreateSC{},
	}
	for _, cmd := range commands {
		err := cmd.Register(parser.Command)
		if err != nil {
			logger.Errorf("Couldn't register command %v: %v", cmd.Name(), err)
			os.Exit(1)
		}
	}

	remaining, err := parser.Parse()
	if e, ok := err.(*flags.Error); ok {
		if e.Type == flags.ErrHelp {
			return
		}
		os.Exit(1)
	}

	if len(remaining) > 0 {
		fmt.Println("Error: Unrecognized arguments passed: ", remaining)
		os.Exit(2)
	}

	switch len(opts.Verbose) {
	case 2:
		logger.SetLevel(logging.DEBUG)
	case 1:
		logger.SetLevel(logging.INFO)
	default:
		logger.SetLevel(logging.WARN)
	}

	// If a sub-command was passed, run it
	if parser.Command.Active == nil {
		os.Exit(2)
	}

	name := parser.Command.Active.Name
	for _, cmd := range commands {
		if cmd.Name() == name {
			err := cmd.Run()
			if err != nil {
				fmt.Println("Error: ", err)
				os.Exit(1)
			}
			return
		}
	}

	fmt.Println("Error: Command not found: ", name)
}

func createClient(args Command, readFile bool) (*sscclient.SSCClient, error) {
	url := args.URLPassed()
	if url == "" {
		url = sscclient.DefaultURL
	}
	keyfile := ""
	if readFile {
		var err error
		keyfile, err = validateKeyfile(args.KeyfilePassed())
		if err != nil {
			return nil, err
		}
	}
	return sscclient.NewSSCClient(url, keyfile)
}

func validateKeyfile(keyfile string) (string, error) {
	if keyfile == "" {
		username, err := user.Current()
		if err != nil {
			return "", err
		}
		return path.Join(
			username.HomeDir, ".sawtooth", "keys", username.Username+".priv"), nil
	}
	return keyfile, nil

}
