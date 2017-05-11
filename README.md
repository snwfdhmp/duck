# [Duck]() ![logo][logo-xs]

![logo][logo-lg]

# [Introduction]()

[Duck]() provides **short syntax** to shortcut common actions on your **C++ projects**.

This is still early *development phase*, but the tool is **already ready to be used**.

# Table of contents

* [Introduction](#)
* [Presentation](#presentation)
  * [Features](#features)
  * [Command usage](#command-usage)
* [How does it works ?](#how-does-duck-work--)
  * [Duck architecture](#duck-architecture)
  * [Project organization ?](#what--)
  * [Deploy the Duck Architecture](#deploy-the-duck-architecture)
  * [How is my code organized ?](#how-is-my-code-organized--)
  * [Compile](#compile)
  * [Builds](#builds)
  * [A little bit about versioning](#a-little-bit-about-versioning)
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

# [How does it works ?]() ![logo][logo-xs]

## [Duck architecture]()

**Duck** is based on the *complete* and *easy-to-install* **Duck architecture**.

If you currently aren't using any [architecture/managing tool](http://github.com/snwfdhmp/duck) for your project, you may want to deploy [**Duck architecture**](#how-does-duck-work--) to organize your source code, ressources, builds, unit tests, classes, or whatever.

## [Project organization]()

Name | Description
--- | ---
`.duck` | project preferences
`build` | app builds
`config` | configuration files
`doc` | documentation
`junk` | trash
`logs` | logs for compilation, unit test, versioning, etc 
`src` | source code

# [Deploy the Duck architecture]()

You can simply *deploy* **Duck architecture** by running `duck deploy`


# [How is my code organized ?]()

Your source code in placed in the `src/` directory (in the root project dir).

Let's see what he contains :

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



# [Builds]()

**All your builds** compiled with `duck compile` are **stored** in `build/` as *tag*-*version*.*suffix*

Examples : `dev-0.1.0` `beta-0.2.7` `release-2.4.1`

## [A little bit about versioning]()

Versioning is **managed automatically** by duck.

You can **show last compiled version** by running `duck pv`.

You can **set next compiled version** by running `duck pv set <version>` (example : `duck pv set 1.5`).

You can **choose the compile tag** when using `duck compile [tag]`. Default is `dev (example `duck compile beta`)


# [Let's configure]() ![logo][logo-xs]

Duck can be easily configured using the `duck.conf` file located in the `.duck` at the root of all projects managed by duck (run `duck init` to create it if it doesn't exists.

Name | Description | Variable | Default Value | Notes
--- | --- | --- | --- | ---
Duck location | Duck's storage path | `$duck_conf_dir` | `/etc/.duck` (Mac) | examples : `/home/user/.duck`, `/var/.duck`, `/usr/share/.duck`
Project path | Your project's root folder. | `$project_root` | user's actual directory when 'duck init' is called
Project name | Used in logs, class creation, unit testing. | `$project_root` | project root folder name when 'duck init' is called
*Folder* name | Rename Duck architecture's folders. ***WARNING!*** This option was written for special case uses, even if you don't have to, you should prefer default configuration. |`$name_dir` | `name` | Available for names `build, backups, config, doc, junk, logs, src, classes, tests`

[logo-xs]: https://www.github.com/snwfdhmp/duck/raw/master/ressources/img/logo-xs.png "Logo"
[logo-lg]: https://www.github.com/snwfdhmp/duck/raw/master/ressources/img/logo-lg.png "Logo"