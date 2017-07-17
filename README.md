![head][head]

## Introduction

[Duck]()![logo][logo-xs] is a **developer tool** which brings **abstraction** to the **terminal**.

### Compiling (C++)

```
$ g++ -o ./binary1.4.3 main.cpp Logger.cpp AnotherFile.cpp MyClass.cpp -lm -Wall
```

becomes

```
$ @ build
```

### Adding package (Go)

```
$ mkdir myNewPackage
$ touch myNewPackage/myNewPackage.go
```

becomes

```
$ @ pack myNewPackage
```

### Commit (git)

```
$ git add *
$ git commit -a -m "My message"
$ git push origin master
```

becomes :

```
$ @ gcp "My message"
```

## Dependencies

| name | installation process |
| --- | --- |
| curl | `sudo apt-get install curl` or `brew install curl` |
| go | [official tutorial](https://golang.org/doc/install) |


## One-line Installation

Installing [duck]() is a very **easy** step.

Once you have `curl` and `go` installed, just run :

```bash
$ curl https://raw.githubusercontent.com/snwfdhmp/duck/master/INSTALL.sh | bash
```

This will download the installation script and execute it. Ensure to have sudo permissions.

## Usage

duck is available under the alias `@` to speed up the command-writing process.

`@ <action> [args]`

| command | description | example |
| --- | --- | --- |
| init | add duck to your project | `@ init` |
| install | install a package | `@ install snwfdhmp/std` |
| lings | view loaded lings | `@ lings` |
| exec | execute your project last compiled binaries | `@ exec` |

To see the list of all available commands, type `@ help`

## Getting started

Create a new directory for your projects

```
$ mkdir my-project
$ cd my-project
```
Init a duck repo in this directory

```
$ @ init
Name: tictactoe
Lang: go
Main: game.go
```

Install the packages you want (see the official repo [here](https://github.com/snwfdhmp/duck-core))

```
$ @ install snwfdhmp/std
$ @ install snwfdhmp/go
$ @ install snwfdhmp/cpp
$ @ install snwfdhmp/junk
```

## Make a ling

Lings are duck's most interesting part.
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

<code>.duck/pkg/go.pkg</code>

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

Duck is open to contributors, feel free to open issues as well.

## Author

- [snwfdhmp](http://github.com/snwfdhmp) (I'm currently the only one on this project.)



[logo-xs]: https://www.github.com/snwfdhmp/duck/raw/master/ressources/img/logo-xs.png "Logo"
[head]: https://www.github.com/snwfdhmp/duck/raw/master/ressources/img/head.png "Head"
