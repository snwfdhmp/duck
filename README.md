# Message from the developer

As duck is being rewritten, this documentation might contain incorrect informations. I will get it up to date asap.

<!-- ![head][head] -->

## Introduction

[Duck]()![logo][logo-xs] is a **developer tool** which brings **abstraction** to the **terminal**.

## Examples usage

### C++ : compile


#### before
```
$ g++ -o ./binary1.4.3 main.cpp Logger.cpp AnotherFile.cpp MyClass.cpp -lm -Wall
```

#### after

```
$ @ build
```

### Go : create packages

#### before

```
$ mkdir myNewPackage
$ touch myNewPackage/myNewPackage.go
```

#### after

```
$ @ pack myNewPackage
```

### git: add, commit, push

#### before

```
$ git add *
$ git commit -a -m "My message"
$ git push origin master
```

#### after

```
$ @ gcp "My message"
```

## Dependencies

| name | installation process |
| --- | --- |
| curl | `apt install curl` or `brew install curl` or whatever |
| go | [official tutorial](https://golang.org/doc/install) |


## Installation

- Download the latest version of duck [here](https://github.com/snwfdhmp/duck/releases)

- Move it into /usr/local/bin and name it `duck`

- Run `ln -s /usr/local/bin/duck /usr/local/bin/@` to add `@` support

- Download the `project.conf` in ressources/duck.conf (on the repo) and put it in /etc/duck/duck.conf

- Start using duck !

## One-line Installation

> Currently not available

Installing [duck]() from sources is a very **easy** step.

Once you have `curl` and `go` installed, just run :

> WARNING : Due to recent changes, this script is being rewritten. You can still install duck but not with the script.

```bash
$ curl https://raw.githubusercontent.com/snwfdhmp/duck/master/INSTALL.sh | bash
```

This will download the installation script and execute it. Ensure to have sudo permissions.

## Manual Installation

This will come later.

## Usage

duck is available under the alias `@` to speed up the command-writing process.

> if `@` is not available for you, run `ln -s $(which duck) /usr/local/bin/@`

usage: `@ <action> [args]`

| command | description |
| --- | --- |
| `@ init` | add duck to your project |  |
| `@ install pkg` | download and install package pkg |
| `@ lings` | view loaded lings |
| `@ exec` | run your project |
| `@ repo-list` | print a list of installed repositories |
| `@ mkdir` | mkdir a directory if it doesn't exist |
| `@ buid` | build your project |

To see the list of all available commands, type `@ help`

## Getting started

### Create a new directory for your projects

```
$ mkdir my-project
$ cd my-project
```
### Init a duck repo in this directory

```
$ @ init
Name: tictactoe
Lang: go
Main: game.go
```

### Install the packages you want
> see the official repo [here](https://github.com/snwfdhmp/duck-core) to discover packages

```
$ @ install snwfdhmp/std
$ @ install snwfdhmp/go
$ @ install snwfdhmp/cpp
$ @ install snwfdhmp/junk
```

## Make a ling

*Lings* are duck's **most interesting part**.
They are custom commands you build to avoid repeating commands.

Examples :

| This | Will execute |
| --- | --- |
| `@ pack MyPackage` | `mkdir MyPackage && touch MyPackage/MyPackage.go` |
| `@ gcp "My message"` | `git add * && git commit -a -m "My message" && git push` |

### Tags

You can use different $tags in a ling

| Tag | Description | Example |
| --- | --- | --- |
| *$main* | Your project's main file | `main.go` |
| *$name* | Your project's name | `myAwesomeProject` |
| *$path* | Path to your project | `/home/snwfdhmp/my-project` |
| *$1*, *$2*, ..., *$9* | Commands arguments (like in shell) | `@ create toto` => `mkdir toto && touch toto/toto.go` |

Example lings using tags :

| This | Will execute |
| --- | --- |
| `@ build` | `go build -o $path/$name` |
| `@ junk fileToThrow.txt` | `mv fileToThrow.txt $path/.junk` |

### Sample package

Packages contain lings

You can build packages to import/export lings.

Create a file in YOURPROJECT/.duck/YOURNAME/PKGNAME.pkg

<code>.duck/pkg/snwfdhmp/go.pkg</code>

```json
{
	"Dependencies":[],
	"Lings":[
		{
			"Label":"build",
			"Description":"compile project",
			"Commands":["go build -o $name $main"],
			"Aliases":["b"]
		},
		{
			"Label":"pack",
			"Description":"create a new package",
			"Commands":[
				"mkdir $1",
				"touch $1/$1.go"
				],
			"Aliases":["p"]
		},
		{
			"Label":"run",
			"Description":"go run your project",
			"Commands":[
				"go run $main"
				],
			"Aliases":["r"]
		}
	]
}
```


## Contributing

Duck is currently closed to contributions.

However, please feel free to open issues.

## Author

- [snwfdhmp](http://github.com/snwfdhmp) (I'm currently the only one on this project.)

## Thanks to

- [GoReleaser](https://github.com/goreleaser/goreleaser) used for releases management


[logo-xs]: ressources/logo-xs.png "Logo"
[head]: https://www.github.com/snwfdhmp/duck/raw/master/doc/ressources/img/head.png "Head"
