# Duck

## Presentation

Duck provides short syntax to shortcut common actions on your C++ projects.

This is still early development phase, but the tool is already ready to be used.

## Features

  - **Easy to configure** [more informations](#configure)

  - **Class management** (create, delete, list)

  - **Code organization** (Duck architecture) [more informations](#duck-architecture)

  - **Junk code management**

  - **Versioning** (prefix-tag) *example : beta-0.1.4 release-1.2*

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

## Duck architecture

`duck deploy`

  Deploy duck architecture in project dir :

  `build`
  
    - all your app builds
  
  `config`
  
    - some configuration variables
  
  `doc`
  
    - documentations
  
  `junk`
  
    - for junk code
  
  `logs`
  
    - logs for compilation, unit test, versioning, etc
  
  `src`
  
    - contains your source code
  
  `compile`
  
    - compilation file
  
  `runtest`
  
    - unit test central
  
  `run`
  
    - run the last build of your project

## Organization

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

- **Duck storage location**

  Used to store default configurations for your projects.

  variable : `$duck_conf_dir`

  default value : `/etc/.duck`

  other value example : `/home/user/.duck`, `/var/.duck`, `/usr/share/.duck`

- **Project root**

  Used to perform actions on your project. Having an invalid root path will disable many Duck features.

  variable : `$project_root`

  default_value : `actual directory when 'duck init' is called`

- **Project name**

  Used in logs, class creation, unit testing.

  variable : `$project_root`

  default_value : `actual directory name when 'duck init' is called`

- ***Folder* name**
  
  Rename Duck architecture's folders.

  **This option was written for special case uses, even if you don't have to, you should prefer default configuration.**

  variable : `$name_dir`

  default value : `name`

  Availaible for names : `build | backups | config | doc | junk | logs | src | classes | tests`
