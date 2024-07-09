# Ascii-art Project

## Overview

This is a Golang  program that renders input text strings into ASCII art representations. Transform your simple text into an exciting visual that stands out in any plain text environment.

## Objectives

- Receive input strings and convert them into graphic representations using ASCII characters.
- Support varied input including alphanumerics, spaces, special characters, and line breaks.
- Offer multiple ASCII art fonts such as shadow, standard, and thinkertoy.

## Technical Specifications

- **Language**: The entire project is implemented in Go (Golang).
- **Coding Standards**: Follow Go's best practices for a clean, maintainable codebase.
- **Unit Testing**: Includes a test to ensure the reliability and stability of the application.

## Banner Format

Ascii-art characters adhere to the following specifications:

- Fixed height of 8 lines.
- New line separation for each character representation.

## Usage
For you to use our project your have to clone the repository from gitea 

```shell
git clone https://learn.zone01kisumu.ke/git/aadero/ascii-art/fs.git
```

once you have cloned it you go to the terminal and enter to the directory where the program is contained.

```shell
cd ascii-art-fs
```
To run Ascii-art, use the Go command-line interface with the following syntax:

```shell
go run . "Your desired string"
```

## Example

if we run 
```go
go run . "Hello\n" | cat -e
```
on the terminal we expect an ascii art to be printed for hello and new line i.e
```

 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
$
```

if you use characters which are out of range in ascii between 32 and 126 it print's out an error message displaying it's not an ascii character
```shell
go run . "さようなら" 
```
the output will be: "Not an ascii character"


## Banner file switch
If you look at the resources folder it contains banner files that can be used to provide the various types of formats in which you want to print the word. We have handled a set of banner files including 


* standard
``` 
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
```

 * thinkertoy
```
                 $
o        o o     $
|        | |     $
O--o o-o | | o-o $
|  | |-' | | | | $
o  o o-o o o o-o $
                 $
                 $
```
* shadow
```
                                 $
_|                _| _|          $
_|_|_|     _|_|   _| _|   _|_|   $
_|    _| _|_|_|_| _| _| _|    _| $
_|    _| _|       _| _| _|    _| $
_|    _|   _|_|_| _| _|   _|_|   $
                                 $
                                 $
```
* ac(our bannerfile)
```
                                 $
AC                AC AC          $
ACACAC     ACAC   AC AC   ACAC   $
AC    AC ACACACAC AC AC AC    AC $
AC    AC AC       AC AC AC    AC $
AC    AC   ACACA  AC AC   ACAC   $
                                 $
                                 $
```

Here is how you run the command;
```bash
go run . "Desired string" bannerfile
```
Example:
```bash
go run . "Test run" thinkertoy
```
Output:
```bash
                                      
o-O-o          o                      
  |            |                      
  |   o-o o-o -o-       o-o o  o o-o  
  |   |-'  \   |        |   |  | |  | 
  o   o-o o-o  o        o   o--o o  o 
                                      
                                      
```

## Contributions

<div style="display: flex; justify-content: space-around; align-items: center;">
<div style ="text-align: center; margin: 10px;">
</div>

### [Clifford Otieno](https://learn.zone01kisumu.ke/git/cliffootieno)
<img src="https://learn.zone01kisumu.ke/git/avatars/7c3793c3fac1a5908d1646d153555890?size=870" width="200">

### [Arnold Odero](https://learn.zone01kisumu.ke/git/aadero)
<img src="https://learn.zone01kisumu.ke/git/avatars/3b0994024734dea36638192cb212b8f1?size=870" width="200">

### [Wilfred Njuguna](https://learn.zone01kisumu.ke/git/wnjuguna)
<img src="https://learn.zone01kisumu.ke/git/avatars/c9b7b96426b4781d5a16fef462551fb5?size=870" width="200">

