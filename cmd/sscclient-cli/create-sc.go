package main

import (
	"github.com/jessevdk/go-flags"

	sscstate "github.com/adiclepcea/socialservicechain/cmd/ssc/state"
)

//CreateNGO is the command structure for createNGO
type CreateSC struct {
	Args struct {
		Name  string `positional-arg-name:"name" required:"true" description:"Name of the social case"`
		Needs []sscstate.Need
	} `positional-args:"true"`
	URL     string `long:"url" description:"Specify URL of REST API"`
	Keyfile string `long:"keyfile" description:"Identify file containing user's private key"`
	Wait    uint   `long:"wait" description:"Set time, in seconds, to wait for transaction to commit"`
}

//Name returns name of the command
func (args *CreateSC) Name() string {
	return "create-sc"
}

//KeyfilePassed returns the keyfile passed in command
func (args *CreateSC) KeyfilePassed() string {
	return args.Keyfile
}

//URLPassed returns the passed in URL
func (args *CreateSC) URLPassed() string {
	return args.URL
}

//Register will register this command in the command processor
func (args *CreateSC) Register(parent *flags.Command) error {
	_, err := parent.AddCommand(args.Name(), "Adds a new SC", "Sends a create SC transaction to set <name> to <value>.", args)
	if err != nil {
		return err
	}
	return nil
}

//Run will run the action associated with this command
func (args *CreateSC) Run() error {
	name := args.Args.Name
	needs := args.Args.Needs
	wait := args.Wait

	sscClient, err := createClient(args, true)
	if err != nil {
		return err
	}
	_, err = sscClient.CreateSC(name, needs, wait)
	return err
}
