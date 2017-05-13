![logo][duck-overview]

# [Introduction]() ![logo][logo-xs]

[Duck]() provides **short syntax** to manage your **C++ projects**.

This is still *development phase*, but the tool is **already ready to be used**.

## Table of contents

* [Presentation](#presentation)
  * [What is Duck ?]()
  * [Features](#features)
  * [Command usage](#command-usage)
* [Getting started](#getting-started)
  * [7 seconds start](#7-seconds-start)
  * [Quick start](#quick-start)
  * [Duck architecture](#duck-architecture)
  * [Project organization](#project-organization)
  * [Deploy the Duck Architecture](#deploy-the-duck-architecture--)
  * [How is my code organized ?](#how-is-my-code-organized-)
  * [Compiling your project](#compiling-your-project)
  * [Project builds](#project-builds)
  * [A little bit about versioning](#a-little-bit-about-versioning)
  * [Run](#run)
  * [Unit tests](#unit-tests)
  * [Logs](#logs)
  * [Let's configure](#lets-configure-)
* [Author](#author)
* [License](#license)

# [Presentation](#presentation)
## [Features](#features)

  - Huge **[time saving](#)** (see here [start coding a new project in 7 seconds](#7-seconds-start))

  - Insanely **[simple configuration](#)** ([more informations](#configure))

  - Simple **[class interface](#)** (create, delete, list)

  - Automatic **[code organization](#)** (Duck architecture) [more informations](#duck-architecture)

  - **[Junk code](#)** managing

  - **[Versioning](#)** (tag-version) *example : beta-0.1.4 release-1.2*

  - (incomming) **[Unit tests automation](#)** (dependencies management, class-test, specialised tests)

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

# [Getting started]() ![logo][logo-xs]

There are many ways to start a new development project. One of them is to use Duck to create all needed files and folders to your project.

Duck always automates the maximum he can, asking something only if he cannot predict what will be the answer. However, you can still configure it to make it behave in the way you would exactly want it to.

## [7 seconds start]()

*The following steps are recommended if you want to create a **new project**.*

<h4>Since **[Duck]() is built to save your time**</h4>, we though that providing a way to do all this stuff in **one command** would be a good thing.

![logo][duck-fast]

`duck deploy ProjectName`. Yeah, as simple. And it works **from anywhere** on an os with duck installed.

You will then be asked a *name* for your *project*. Hit Enter if *'ProjectName'* is the name you want. 

## [Quick Start]()

*The following steps are recommended if you want to **use duck in an existing project** (or different kind).*

This tutorial is the "classic tutorial". Since I implemented `duck deploy projectName` (see [7 seconds start](#7-seconds-start)), I'm wondering if this is still useful (and the answer is probably no).
You can still use it to understand how duck works, but the 7 seconds method is much better.

### 1. **Create a folder for your project** if you don't have one already. (and place yourself in)

![logo][gif-mkdir]

### 2. Let's initialize a **new [duck]() project with `duck init`**

![logo][gif-init]

### 3. Now we will **deploy *[Duck Architecture]()*** using `duck deploy`

![logo][gif-deploy]

### 4. **Let's code !**

You can now begin coding. Yes, it's as simple. Let's recap :

```
mkdir MyProject
cd MyProject
duck init
duck deploy
```

## [Duck architecture]()

**Duck** is based on the *complete* and *easy-to-install* **Duck architecture**.

If you currently aren't using any [architecture/managing tool](http://github.com/snwfdhmp/duck) for your project, you may want to **deploy [Duck architecture](#how-does-duck-work--)** to ***organize*** your **source code**, **ressources**, **builds**, **unit tests**, **classes**, or whatever.

## [Project organization]()

Name | Description
--- | ---
`.duck/` | duck's project **preferences**
`backups/` | **backups** dir (use `duck tar`)
`build/` | your project **builds** (versioned)
`config/` | **configuration** dir
`doc/` | **documentation**
`junk/` | **trash**/junk code
`logs/` | **logs** for compilation, unit test, versioning, etc 
`src/` | The **source code**, organized in classes. This dir also includes unit tests builds (`src/tests/`)

# [Deploy the Duck architecture]() ![logo][logo-xs]

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


# [Compiling your project]()

You can **compile your project** with `duck compile [tag]`.

Duck will **automatically compile** your `main.cpp` and all files in `src/classes/*` folders with extension `.class.cpp` *(by default)*.

You can **compile quickly** with tag *dev* by running `duck -c`

***Warning!* Duck do not provides external library linking for now**. You can still modify the `compile()` function in `duck.sh` if needed.

# [Project builds]()

**All your builds** compiled with `duck compile` are **stored** in `build/` as *tag*-*version*.*suffix*

Examples : `dev-0.1.0` `beta-0.2.7` `release-2.4.1`

## [A little bit about versioning]()

Versioning is **managed automatically** by duck.

You can **show last compiled version** by running `duck pv`.

You can **set next compiled version** by running `duck pv set <version>` (example : `duck pv set 1.5`).

You can **choose the compile tag** when using `duck compile [tag]`. Default is `dev (example `duck compile beta`)

# [Run](#run)

This section is not available yet.

# [Unit tests](#unit-tests)

This section is not available yet.

# [Logs](#logs)

This section is not available yet.


# [Let's configure]() ![logo][logo-xs]

Duck is **easily configurable** using the `duck.conf`.

This file is located in the `.duck` directory in your project folder (run `duck init` to create it if it doesn't exists).

Name | Description | Variable | Default Value | Notes
--- | --- | --- | --- | ---
Duck location | Duck's **storage path** | `$duck_conf_dir` | `/etc/.duck` (Mac) | examples : `/home/user/.duck`, `/var/.duck`, `/usr/share/.duck`
Project path | Your **project's root folder**. | `$project_root` | user's actual directory when 'duck init' is called
Project name | Your **project name** is used in logs, class creation, unit testing. | `$project_root` | project root folder name when 'duck init' is called
*Folder* name | **Rename Duck architecture's folders**. ***WARNING!*** This option was written for special case uses, even if you don't have to, you should prefer default configuration. |`$name_dir` | `name` | Available for names `build, backups, config, doc, junk, logs, src, classes, tests`

---
---

# [Author]()

[snwfdhmp](http://github.com/snwfdhmp) - Visit my other repos, leave a star if you like one ;)

# [License]()

You can find duck's license [here](https://www.github.com/snwfdhmp/duck/raw/master/LICENSE.txt)

[duck-fast]: http://media.giphy.com/media/3og0IBl9zPCFv1tOb6/giphy.gif "duck deploy MyProject"
[duck-overview]: https://www.github.com/snwfdhmp/duck/raw/master/ressources/img/duck-overview.png "Duck overview"
[gif-mkdir]:http://media.giphy.com/media/xUPGcKV7oHQQQSsueA/giphy.gif "mkdir MyProject;cd MyProject"
[gif-init]:http://media.giphy.com/media/l0Iye57FiXyRsKPew/giphy.gif "duck init"
[gif-deploy]:http://media.giphy.com/media/xUA7baQKlb6BjgLpKg/giphy.gif "duck deploy"
[logo-xs]: https://www.github.com/snwfdhmp/duck/raw/master/ressources/img/logo-xs.png "Logo"
[logo-lg]: https://www.github.com/snwfdhmp/duck/raw/master/ressources/img/logo-lg.png "Logo"
