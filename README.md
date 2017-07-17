# [Duck]() ![logo][logo-xs]

## Note to readers

This README isn't complete yet as duck is still in early development phase. Please note that duck is still useable.

We apologize for the lack of informations on this readme.

## Dependencies

| name | installation process |
| --- | --- |
| curl | `sudo apt-get install curl` or `brew install curl` |
| go | [official tutorial](https://golang.org/doc/install) |


## One-line Installation

When you have `curl` and `go` installed, you can install duck easily by executing the following line into a shell.

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
| This | Will execute this |
| --- | --- |
| ```
@ pack MyPackage
``` |
```
mkdir MyPackage
touch MyPackage/MyPackage.go
```
|

We provide some ducklings examples in <code>.duckling/</code>

<code>.duckling/go.duckling</code>

```json
{
	"Ducklings":[
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