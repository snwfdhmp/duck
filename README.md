# Duck

## Presentation

Duck is what makes your project go from this

[IMAGE junk code]

to this

[IMAGE organized repo]

## Commands

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

  - `classes/` *: contains all your classes organized as :*

    <code>*className*/</code>

      <code>*className*.class.cpp</code>

        - *: class method implementation*

      <code>*className*.class.h</code>

        - *class implementation and method declaration*

      <code>*className*.test.cpp</code>

        - *class unit test*

      <code>*className*.test.dependencies</code> *: className*.test.cpp *compilation dependencies*

      - *optional* `someOtherUnitTests.test.cpp` always with their `someOtherUnitTests.test.dependencies`

  `config/` : contains 2 files

    `macros.h`

      - preprocessor macros

    `constants.h` -> preprocessor constants