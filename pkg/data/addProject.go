package data

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type addProjectOptions struct {
	Name        string
	Path        string
	Force       bool
	Interactive bool
}

func AddProject(name, path string, force, interactive bool) error {
	return addProjectToData(addProjectOptions{
		Name:        name,
		Path:        path,
		Force:       force,
		Interactive: interactive,
	})
}

func addProjectToData(opts addProjectOptions) error {
	if Projects.Haskey(opts.Name) && !opts.Force && !opts.Interactive {
		return errors.New("project already exists")
	}

	reader := bufio.NewReader(os.Stdin)

	for Projects.Haskey(opts.Name) && !opts.Force {
		if Projects.Key(opts.Name).Value() == opts.Path {
			opts.Force = true
			break
		}
		fmt.Print("A project named '" + opts.Name + "' already exists. Enter another name (or ENTER to overwrite) : ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return err
		}

		input = input[:len(input)-1] // pop the '\n'

		if input == "" {
			opts.Force = true
			break
		} else {
			opts.Name = input
		}
	}

	if opts.Force {
		p, err := Projects.GetKey(opts.Name)
		if err != nil {
			return err
		}
		p.SetValue(opts.Path)

	} else {
		_, err := Projects.NewKey(opts.Name, opts.Path)
		if err != nil {
			return err
		}
	}

	err := File.SaveTo(Path)
	return err
}
