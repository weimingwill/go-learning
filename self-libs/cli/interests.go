package cli

import (
	"log"

	"distributed-load-testing/common/jsonparser"
	"roshar/services/users-svc/pkg/users"
)

const interestsKey = "interests"

// CreateInterests creats intersts table from
func (cli CLI) CreateInterests(filename string) (err error) {
	names, err := cli.getInterestNames(filename)
	if err != nil {
		return err
	}

	interests, err := cli.createIntersts(names)
	if err != nil {
		return err
	}

	log.Println("create interests successfully.")
	cli.printIntersts(interests)

	return nil
}

// OverwriteInterests deletes all rows in intersts table and
// insert new interests from filename
func (cli CLI) OverwriteInterests(filename string) error {
	names, err := cli.getInterestNames(filename)
	if err != nil {
		return err
	}

	err = cli.Users.DeleteInterests()
	if err != nil {
		return err
	}

	interests, err := cli.createIntersts(names)
	if err != nil {
		return err
	}

	log.Println("overwrite interests successfully.")
	cli.printIntersts(interests)

	return nil
}

func (cli *CLI) getInterestNames(filename string) ([]string, error) {
	getter, err := jsonparser.NewGetter(filename)
	if err != nil {
		return nil, err
	}

	names, err := getter.StringArray(interestsKey)
	if err != nil {
		return nil, err
	}
	return names, err
}

func (cli *CLI) createIntersts(names []string) ([]users.Interest, error) {
	interests, err := cli.Users.CreateInterests(names)
	if err != nil {
		return nil, err
	}
	return interests, nil
}

func (cli *CLI) printIntersts(interests []users.Interest) {
	for _, interest := range interests {
		log.Printf("%+v", interest)
	}
}
