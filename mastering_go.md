# Mastering Go

## Additional Info

- link to book <https://www.packtpub.com/product/mastering-go-third-edition/9781801079310>
- link to Go code 2nd - <https://github.com/PacktPublishing/Mastering-Go-Second-Edition>
- link to Go code, 3rd - <https://github.com/mactsouk/mastering-Go-3rd>
- Mihalis Tsoukalos blog - <https://www.mtsoukalos.eu/post/>
- Golang source - <https://github.com/golang/go>

## Standard Library

- `fmt`
- `os`
  - `os.Args()`
  - `os.Stat()`
    - `os.Stat.Mode()`
    - `os.Stat.IsRegular()`
  - `os.Getenv`
  - `os.TempDir()`
  - `os.OpenFile()`
- `path/filepath`
  - `filepath.SplitList()`
  - `filepath.Join()`
  - `path.Base()`
- `strconv`
  - `strconv.Atoi()`
- `time`:
  - `time.Sleep()`
  - `time.Second()`
  - `time.Microsecond()`
- `log/syslog`:
  - `log.SetOutput()`
  - `syslog.New()`
  - `log.Print()`
  - `log.Fatal()`
  - `log.Panic()`
  - `log.SetFlags()` - possible to change flags during program execution

## Table of Contents

01. A quick introduction to Go
01.01. Introducing Go
01.02. `Hello World!`
01.03. Running Go code
01.04. Important characteristics of Go
01.05. Developing the which utility in Go
01.06. Logging Information
01.07. Overview of Go generics
01.08. Developing a basic phone book application
02. Basic Go Data Types

## 1. A quick introduction to Go

### 01.01. Introducing Go

Go is a general-purpose programming language, primary used for writing system tools, command line utilities, web services and software that works over networks. Go can help you develop the following kinds of applications:

- professional web services
- networking tools and services, such as K8s and Istio
- Backend systems
- System utilities
- command line utilities, like docker and hugo
- applications that exchange data in JSON format
- applications that process data from relational databases, NoSQL databases and other popular storage systems
- compilers and interpreters for programming languages you design
- database systems, like CockroachDB,  and key/value stores, like etcd

Go  syntax looks like C whereas the package concept was inspired by Modula-2. Go is a programming language, tools and standard library.

Go doc allow you to see the documentation of existing Go functions and packages w/o needing an internet connection:

- `go doc fmt`
- `go doc fmt.Println`

### 01.02. `Hello World!`

```go
package main 

import "fmt"

func main() {
    fmt.Println("Hello World!")
}
```

Each Go source code begins w/ a package declaration. `package main` has a special meaning in Go. To create an executable application use a main() function. Go considers this the entry point to the application and begins the execution of the application with the code found int the main() function of the main package.

Two characteristics make Go code an autonomous source that can generate and executable binary:

- the name if the package, which should be main and
- the presence of the main() function.

Go programs are organized in packages - even the smallest Go program should be delivered as a package.

### 01.03. Running Go code

There are 2 ways t execute Go code:

1. as a compiled language, using `go build .` where go build will look for a main package in the current directory.
2. as a scripting language, using `go run`, which builds the package, creates a temp. executable file, executes that file and deletes it once it is done.

- `go fmt FILE.go` - to format go code automatically

Coding blocks in Go are embedded in curly braces even if they contain just a single statement or no statement at all.

### 01.04. Important characteristics of Go

#### Defining and using variables

- If not initial value is given to a variable, the Go compiler will automatically initialize that variable to the zero value of its data type.
- `:=` short assignment statement, is frequently used in Go, especially for getting the return values from functions and for loops with the range keyword. Short assignment statement cannot be used outside of a function environment.
- You will rarely see the use of `var` in go the `var` keyword is mostly used for declaring global or local variables w/o an initial value, because every statement that exists outside of the code of a function must begin w/ a keyword such as `func` or `var`.

Make Global variables stand out by either beginning them with uppercase letter or using all capital letters.

#### Controlling program flow

- `common_patterns.go`
- `switch.go`

<!-- TODO check for keywords in Go -->

#### Working w/ command-line arguments

By default, command line arguments are stored in the `os.Args` slice.

The first command-line argument stored in `os.Args` slice is always the full path and name  of the executable.

- `go run Args.go fileToProcess -n 10 /tmp "A phrase"`
- `go run min_max.go z 1 2 3 -100 100 100.0 two three`

#### Understanding the Go concurrency model

The Go concurrency model is implemented using `goroutines` and `channels`:

- A `goroutine` is the smallest executable Go entity. In order to create a new goroutine you have to use the go keyword followed by a predefined or  anonymous function.
- A `channel` in Go is a mechanism that allows goroutines to communicate and exchange data.

It is easy to create goroutine, but there are other difficulties when dealing w/ concurrent programming, including goroutine synchronization and sharing data b/w goroutines - this is a Go mechanism for avoiding side effects, when running goroutines. As `main()` runs as goroutine as well you do not want `main()` to finish before other goroutines of the program, because when `main()` exits, the entire program along with any goroutines that have not yet finished will terminate.

Although goroutines do not share any variables, they can share memory. The good thing is that there are various techniques for the `main()` function to wait for goroutines to exchange data through channels for using shared memory.

- `go run goroutine.go` - running this multiple times will generate different output each time. This is because goroutines are initialized in random order and start running in random order. The Go Scheduler is responsible for the execution of goroutines, similarly as OS scheduler is responsible for the execution of OS threads.

### 01.05. Developing the which utility in Go

A good way of learning a new programming language is by trying to implement simple versions of traditional UNIX utilities.

In our implementation:

- first part is about reading the input argument
- second part is about reading the PATH variable, splitting it and iterating over the directories of the PATH variable
- the 3rd part is about looking for the desired binary file, and determining whether that can be found, whether it is a regular file and whether it is an executable file.

To execute the program run:

- `go run which.go which`

The `os` package is used to interact w/ underlying operating system and `path/filepath` is used for working with the contents of the PATH variable that is read as a long string, depending on the number of directories it contains.

- `filepath/SplitList()` offers a portable way of separating a list of paths.
- `filepath/Join()` is used for concatenating the different parts of a path using OS-specific separator - this makes `filepath/Join()` work in all supported OS.

### 01.06. Logging Information

Generally speaking, using a log file to write some information used to be considered a better practice than writing the same output on screen for 2 reasons:

1. because the output does not get lost as it is stored on a file
2/ you can search and process log files using UNIX tools, like `grep`, `awk` and `sed`.

However this is not true anymore.

As we usually run our services via system, program should log to `stdout` so `systemd` can put logging data in the journal. Additionally in cloud native applications, we are encouraged to simply log to `stderr` and let container system redirect the stderr stream to the desired destination.

In Go the `log` package sends log messages to `stderr`. Part of `log` package is the `log/syslog` package, which allows you to send log messages to the syslog server of your machine. Although by default log writes to stderr, the use of log.SetOutput() modifies that behaviour.

In order to write to system logs call the `syslog.New()` function. Writing to the main system log file is as easy as calling `syslog.New()` with the `syslog.LOG_SYSLOG` option. After that you need to tell your Go program that all logging information goes to the new logger - this is implemented with the call to the `log.SetOutput()` function.

- `go run log_syslog.go`

After the call to `log.SetOutput()`, all logging information goes to the syslog logger variable, that sends it to syslog.LOG_SYSLOG.

#### log.Fatal() and log.Panic()

- The log.Fatal() function is used when something erroneous has happened and you just want to exit your app as soon as possible after reporting that bad situation. The call to log.Fatal() terminates a Go program at the point where log.Fatal() was called after printing an error message. Additionally it returns back a non-zero exit code, which in UNIX indicates an error.

- There are situations, when a program is about to fail for good and you want it to have as much information about the failure as possible - log.Panic() implies that something really unexpected and unknown has happened. Analogous to log.Fatal(), log.Panic() prints a custom message and immediately terminates a Go program. Output of Log.Panic() includes additional low-level information.

Parts tat make log.Panic() and log.Fatal():

- `log.Panic()` = `log.Print()` + `panic()`
- `log.Fatal()` = `log.Print()` + `os.Exit()`

- `go run log_fatal.go`

#### Writing to a custom log file

Especially on applications and services that are deployed to prod you just nee to write your logging data in a log file of your choice. There can be many reasons for this, including:

- writing debugging data w/o messing w/ the system log files,
keeping your own logging data separate from system logs in order to transfer it or store it in a database or software like Elasticsearch.

The path of the log file that is used is hardcoded into the code using a global variable named LOGFILE. For the purpose of this chapter, that log file resides inside the `/tmp` directory, which is not the usual place for storing data.

- `go run log_custom.go`

#### Printing line nr in log file

The desired functionality is implemented with the use of `log.Lshortfile` in the parameter of `log.New()` or `SetFlags()`. The `log.Lshortfile` flag adds the filename as well as the line number of the Go statement that printed the log entry in the log entry itself.

- `log.Llongfile` adds the full path of the Go file.

- `go run log_custom_line_nrs.go`

### 01.07. Overview of Go generics

The main idea behind generics in Go is not having to write special code for supporting multiple data types when performing the same task. For example, you can see Go support for multiple data types in functions such as `fmt.Println()` using interfaces and reflection.

- `go run generics.go`

- generics variable `[T any]` specified after function name and before function parameter s.

### 01.08. Developing a basic phone book application

command line utility that searches a slice of structures that is statically defined (hardcoded) in the Go code. The utility offers support for two commands:

- `search` - searches for given surname and return its full record
- `list` - lists all available records, respectively.
- `go run phoneBook.go`

## 02. Basic Go Data Types
