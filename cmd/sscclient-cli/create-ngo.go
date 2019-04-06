package main

import (
	"github.com/jessevdk/go-flags"
)

//CreateNGO is the command structure for createNGO
type CreateNGO struct {
	Args struct {
		Name  string `positional-arg-name:"name" required:"true" description:"Name of the key"`
		Value string `positional-arg-name:"value" required:"true" description:"Name of the NGO"`
	} `positional-args:"true"`
	URL     string `long:"url" description:"Specify URL of REST API"`
	Keyfile string `long:"keyfile" description:"Identify file containing user's private key"`
	Wait    uint   `long:"wait" description:"Set time, in seconds, to wait for transaction to commit"`
}

//Name returns the name of the command
func (args *CreateNGO) Name() string {
	return "create-ngo"
}

//KeyfilePassed returns the keyfile passed in command
func (args *CreateNGO) KeyfilePassed() string {
	return args.Keyfile
}

//URLPassed returns the passed in URL
func (args *CreateNGO) URLPassed() string {
	return args.URL
}

//Register will register this command in the command processor
func (args *CreateNGO) Register(parent *flags.Command) error {
	_, err := parent.AddCommand(args.Name(), "Adds a new NGO", "Sends a create NGO transaction to set <name> to <value>.", args)
	if err != nil {
		return err
	}
	return nil
}

//Run will run the action associated with this command
func (args *CreateNGO) Run() error {
	name := args.Args.Name
	value := args.Args.Value
	wait := args.Wait

	sscClient, err := createClient(args, true)
	if err != nil {
		return err
	}
	_, err = sscClient.CreateNGO(name, value, wait)
	return err
}
