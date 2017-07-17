# [Duck]() ![logo][logo-xs]

## Introduction

[Duck]() is a **developer tool** which brings **abstraction** to the **terminal**.

### Compiling (C++)

`g++ -o ./binary1.4.3 main.cpp Logger.cpp AnotherFile.cpp MyClass.cpp -lm -Wall`

becomes

`@ build`

### Adding package (Go)

```
mkdir myNewPackage
touch myNewPackage/myNewPackage.go
```

becomes

```
[@]() pack myNewPackage
```

### Commit (git)

```
git add *
git commit -a -m "My message"
git push origin master
```

becomes :

```
[@]() gcp "My message"
```

## Dependencies

| name | installation process |
| --- | --- |
| curl | `sudo apt-get install curl` or `brew install curl` |
| go | [official tutorial](https://golang.org/doc/install) |


## One-line Installation

Installing [duck]() is a very **easy** step.

Once you have `curl` and `go` installed, just run :

`curl https://raw.githubusercontent.com/snwfdhmp/duck/master/INSTALL.sh | bash`

This will download the installation script and execute it. Ensure to have sudo permissions.

## Usage

`duck <action> [args]`


## Getting started

Create a new directory for your projects

```
$ mkdir my-project
$ cd my-project
```
Init a duck repo in this directory

```
$ duck init
Name: tictactoe
Lang: go
Main: game.go
```

Install the packages you want (see the official repo [here](https://github.com/snwfdhmp/duck-core))

```
$ duck install snwfdhmp/std
$ duck install snwfdhmp/go
$ duck install snwfdhmp/cpp
$ duck install snwfdhmp/junk
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

<code>.duck/pkg/go.pkg</code>

```json
{
	"Lings":[
		{
			"Label":"compile",
			"Description":"compile project",
			"Commands":["go build -o $name $main"],
			"Aliases":["build"]
		},
		{
			"Label":"create",
			"Description":"create necessary files for a new package",
			"Commands":[
				"mkdir $1",
				"touch $1/$1.go"
				]
		}
	]
}
```


## Contributing

Duck is open to contributors, feel free to open issues as well.

## Author

- [snwfdhmp](http://github.com/snwfdhmp) (I'm currently the only one on this project.)



[logo-xs]: https://www.github.com/snwfdhmp/duck/raw/master/ressources/img/logo-xs.png "Logo"