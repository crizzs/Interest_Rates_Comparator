# Interest Rates Comparator

## Overview
Through this console-based application, you are able to compare Interest Rates Data in Singapore (Fixed Deposits and Saving Deposits Rates) between Banks and Finance Companies across different financial period. 

## Consideration on App Design

This app is designed in consideration of `Single Responsibility Principle`, `Code Reusability` , `Test-Driven Development(TDD)`, `Performance` and `Portability`. 

`Single Responsibility Principle` : Borrows idea from Object-Oriented Design, classes have been tasked with its duty (Display-Controller-Models). One thing to note, in Golang, there is no concept of classes but structs which allows for polymorphism. 

`Code Reusability`: Using the nature of Go, package `masapi` could be redeployed for different use with minimal code changes. (Eg. Becoming API and etc.)

`Test-Driven Development(TDD)`:Develops based on the idea of building the unit tests first. Test coverage designed to hit at least the ballpark 75%-80%.

`Performance`: Memory is only allocated based on the data available. Minimize the use of complex structure and try to stay within O(n) performance. 

`Portability`: This app could be deployed onto multiple platforms in complied form (executable).

## Requirements

- [Golang 1.9 & Up](https://golang.org/)
- [Golang 1.5 and Below - Requires GCC-toolchain for Cross-Complier](https://golang.org/)
- [TDD Assertion dependency](https://github.com/stretchr/testify)

## Installation

1. To install Interest Rates Comparator use `go get`:

    `go get github.com/crizzs/Interest_Rates_Comparator`

2. Go into your Golang Workspace and search for the `Interest_Rates_Comparator` folder

    Part 1: Open your command prompt or terminal   
     
    Part 2: Navigate into the folder using the  command 
    
3. Use Makefile to download deps, test and build your executable(s). Type `make` into your terminal.
4. Executable will be created. This file is named `main` 
5. For windows, run `main.exe`. For Linux/Mac, run `./main` in terminal.

## How to use the app?

1. You will be greeted by an interactive screen terminal.
2. Set a `Start Date(Jan-2017)` and `End Date(Dec-2017)` for financial period
3. Walla! You can input 1 to 6 to navigate different functions
4. If you are still unsure??? View the `screenshots` folder

## Screenshots

Screenshots are stored inside `screenshots` folder.
