// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/fatih/color"
	"github.com/snwfdhmp/duck/cmd"
	"github.com/snwfdhmp/duck/pkg/data"
	"github.com/snwfdhmp/duck/pkg/pkg"
	"github.com/snwfdhmp/duck/pkg/projects"
	"github.com/spf13/afero"
)

var (
	fs           afero.Fs = afero.NewOsFs()
	paths        []string
	duckCommands []pkg.Command
)

func main() {
	var err error
	var paths []string

	sources := map[string]bool{
		projects.PackagesPath: (!projects.HasErrors && projects.IsInside),
		data.PackagesPath:     !data.HasErrors,
	}

	for target, condition := range sources {
		if !condition {
			continue
		}

		paths = append(paths, target)
	}

	duckCommands, err := pkg.ReadDirs(paths)
	if err != nil {
		color.Red(err.Error())
	}

	pkg.CreateCobraCommands(cmd.RootCmd, duckCommands)

	cmd.Execute()
}
