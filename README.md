# Duck

## Presentation

Duck provides short syntax to shortcut common actions on your C++ projects.

This is still early development phase, but the tool is already ready to be used.

## Features

  - **Easy to configure** [more informations *(incomming)*](#)

  - **Class management** (create, delete, list)

  - **Code organization** (Duck architecture) [more informations](#Duck-architecture)

  - **Junk code management**

  - **Versioning** (version tag+prefix) *example : beta-0.1.4 release-1.2*

  - (incomming) **Unit tests automation** (dependencies management, class-test, specialised tests)

## Configuration

# Paths:

Duck configuration folder
variable : $duck_conf_dir
default : /etc/.duck

Project root
variable $project_root
default : '.' when user call 'duck init'

## Usage

```
usage: duck <options> 

Available options :

name             usage                      description

---------------  -------------------------  ---------------------------------

class            duck class <action> [opt]  tools for classes

compile          duck compile [env]         project compiler

run              duck run [target-version]  run project (no arg->last version)

project-version  duck pv [show|set|inc]     configuration for project version

tar              duck tar                   backup 'src/' dir into a tarball

quick-commit     duck qc [custom-msg]       alias git add *, commit, push

help             duck help                  shows this message

doc              duck doc [command]         shows command's help message

```

## Duck architecture

`duck create *project-name*`

  Creates folders for your project including :

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

##Organization

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
