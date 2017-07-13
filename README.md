# Duck ![logo][logo-xs]

## Note to readers

This README isn't complete yet as duck is still in early development phase. Please note that duck is still useable.

We apologize for the lack of informations on this readme.

## Installation

We do not provide any installation script yet. However, you can compile duck :

```bash
go build duck.go
```

## Usage

Assuming duck has been compiled to your current working directory, you can use it by running:

```bash
./duck <action> args
```

## Make a duckling

Ducklings are duck's most interesting part.
They are automated commands configured by yourself to fit your development environment perfectly.

Even if you'd normally build a duckling yourself, we provide some ducklings examples in <code>.duck/go.duck</code> (Schemes' elements)

<code>.duck/go.duck</code>

```shell
{
	"Schemes":[
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
		},
		{
			"Label":"run",
			"Description":"run binaries",
			"Commands":["$path/$name"],
			"Aliases":["r"]
		},
		{
			"Label":"commit",
			"Description":"add all files to git and commit $1",
			"Commands":[
				"git add *",
				"git commit -a -m '$1'"
			],
			"Aliases":["qc"]
		}
	]
}
```


## Contributing

Duck is open to contributors, feel free to open issues as well.

## Author

- [snwfdhmp](http://github.com/snwfdhmp) (I'm currently the only one on this project.)



[logo-xs]: https://www.github.com/snwfdhmp/duck/raw/master/ressources/img/logo-xs.png "Logo"