# [Duck]() ![logo][logo-xs]


# Introduction

[Duck]() provides **short syntax** to shortcut common actions on your **C++ projects**.

This is still early *development phase*, but the tool is **already ready to be used**.

![logo][logo-xs]

# [Table of contents]()

* [Introduction](#introduction)
* [Presentation](#presentation)
  * [Features](#features)
  * [Command usage](#command-usage)
* [How does Duck work ?](#how-does-duck-work--)
  * [Duck architecture](#duck-architecture)
    * [What ?](#what--)
    * [Why ?](#why--)
  * [Code](#code)
  * [Compile](#compile)
  * [Builds](#builds)
  * [Versioning](#versioning)
  * [Run](#run)
  * [Unit tests](#unit-tests)
  * [Logs](#logs)

## [Presentation](#presentation)
## [Features](#features)

  - **Easy to configure** [more informations](#configure)

  - **Class management** (create, delete, list)

  - **Code organization** (Duck architecture) [more informations](#duck-architecture)

  - **Junk code management**

  - **Versioning** (tag-version) *example : beta-0.1.4 release-1.2*

  - (incomming) **Unit tests automation** (dependencies management, class-test, specialised tests)

## [Command usage]()

```
usage: duck <options> 

Available options :

name             usage                      description
----             -----                      -----------
init             duck init <name> [opt]     initialize a new duck project
deploy           duck deploy [opt]          deploy duck architecture
config           duck config <action> [opt] project configuration tools
class            duck class <action> [opt]  tools for classes
compile          duck compile [env]         run project compiler
run              duck run [target-version]  run project (no arg->last version)
project-version  duck pv [show|set|inc]     configuration for project version
tar              duck tar                   backup 'src/' dir into a tarball
quick-commit     duck qc [custom-msg]       alias git add *, commit, push
doc              duck doc [command]         shows command's help message
help             duck help                  shows this message

```

# [How does Duck work ?]() ![logo][logo-xs]

## [Introduction]()

**Duck** is based on the *complete* and *easy-to-install* **Duck architecture**.

If you currently aren't using any [architecture/managing tool](http://github.com/snwfdhmp/duck) for your project, you may want to deploy [**Duck architecture**](#how-does-duck-work--) to organize your source code, ressources, builds, unit tests, classes, or whatever.

## [Project root directory]()

Name | Description
--- | ---
`.duck` | project preferences
`build` | app builds
`config` | configuration files
`doc` | documentation
`junk` | trash
`logs` | logs for compilation, unit test, versioning, etc 
`src` | source code

## Deploy

You can *deploy* **Duck architecture** by running `duck deploy`

# Builds

**All your builds** compiled with `duck compile` are **stored** in `build/` as *tag*-*version*.*suffix*

Examples : `dev-0.1.0` `beta-0.2.7` `release-2.4.1`

# Versioning

Versioning is **managed automatically** by duck.

You can **show last compiled version** by running `duck pv`.

You can **set next compiled version** by running `duck pv set <version>` (example : `duck pv set 1.5`).

You can **choose the compile tag** when using `duck compile [tag]`. Default is `dev (example `duck compile beta`)


## Code

Your code is organized that way :

- `src/`

  - Your `main.cpp` : main file

  - `classes/` : contains all your classes organized in their *className* dir as

      <code>*className*.class.cpp</code>

        - class method implementation

      <code>*className*.class.h</code>

        - class implementation and method declaration

      <code>*className*.test.cpp</code>

        - class unit test

      <code>*className*.test.dependencies</code>

        - '*className*.test.cpp' compilation dependencies

      + *optional* `someOtherUnitTests.test.cpp` always with their `someOtherUnitTests.test.dependencies`

  - `config/` : contains 2 files

    `macros.h` -> preprocessor macros

    `constants.h` -> preprocessor constants


# Configure ![logo][logo-xs]

Name | Description | Variable | Default Value | Notes
--- | --- | --- | --- | ---
Duck location | Duck's storage path | `$duck_conf_dir` | `/etc/.duck` (Mac) | examples : `/home/user/.duck`, `/var/.duck`, `/usr/share/.duck`
Project path | Your project's root folder. | `$project_root` | user's actual directory when 'duck init' is called
Project name | Used in logs, class creation, unit testing. | `$project_root` | project root folder name when 'duck init' is called
*Folder* name | Rename Duck architecture's folders. ***WARNING!*** This option was written for special case uses, even if you don't have to, you should prefer default configuration. |`$name_dir` | `name` | Available for names `build, backups, config, doc, junk, logs, src, classes, tests`

[logo-xs]: https://www.github.com/snwfdhmp/duck/raw/master/ressources/img/logo-xs.png "Logo"
[logo-sm]: https://www.github.com/snwfdhmp/duck/raw/master/ressources/img/logo-sm.png "Logo"