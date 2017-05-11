# Duck

## Presentation

Duck provides short syntax to shortcut common actions on your C++ projects.

This is still early development phase, but the tool is already ready to be used.

## Features

  - **Easy to configure** [more informations](#configure)

  - **Class management** (create, delete, list)

  - **Code organization** (Duck architecture) [more informations](#duck-architecture)

  - **Junk code management**

  - **Versioning** (tag-version) *example : beta-0.1.4 release-1.2*

  - (incomming) **Unit tests automation** (dependencies management, class-test, specialised tests)



## Usage

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

# Duck architecture

  You can deploy Duck's architecture by running `duck deploy`

  Name | Description
  --- | ---
  `build` | app builds
  `config` | configuration files
  `doc` | documentation
  `junk` | trash
  `logs` | logs for compilation, unit test, versioning, etc 
  `src` | source code

## App builds

All your builds compiled with `duck compile` are stored in `build/` as *tag*-*version*.*suffix*

Examples : `dev-0.1.0` `beta-0.2.7` `release-2.4.1`

### Version

Versioning is managed automatically by duck.

You can see last compiled version by running `duck pv`.

You can set next compiled version by running `duck pv set <version>` (example : `duck pv set 1.5`).

### Tag

The tag is choosen when using `duck compile [tag]` (example `duck compile beta`)


## Source code

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


# Configure

Name | Description | Variable | Default Value | Notes
--- | --- | --- | --- | ---
Duck location | Duck's storage path | `$duck_conf_dir` | `/etc/.duck` (Mac) | examples : `/home/user/.duck`, `/var/.duck`, `/usr/share/.duck`
Project path | Your project's root folder. | `$project_root` | user's actual directory when 'duck init' is called
Project name | Used in logs, class creation, unit testing. | `$project_root` | project root folder name when 'duck init' is called
*Folder* name | Rename Duck architecture's folders. ***WARNING!*** This option was written for special case uses, even if you don't have to, you should prefer default configuration. |`$name_dir` | `name` | Available for names `build, backups, config, doc, junk, logs, src, classes, tests`
