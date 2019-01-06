# Advent of Code 2017 in Go

These are my soluions for [Advent of Code 2017](https://adventofcode.com/2017/) written in Go.

I've solved them in Coffeescript one year ago. But since I'm learning Go at the moment, and have just finished writing solutions for puzzles from this year (2018), I've decided to go with year 2017 puzzles as well, refactoring and cleaning up solutions as much as possible.

Goals:

* implemented automatic puzzle input retrieval and result submission;
* implement generic code, that can come in handy during programming contests;
* same as for year 2018: learn & document quirks & tricks of Go, which are new to me.

## "Go gotchas"

Go is low-level language with built-in concurrency and garbage collection, designed as a highly efficient C++ or Java competitor. To achieve high speed (both in compilation and execution), some surprising design decisions were made. It takes time to learn them.

For list of quirks found previously, refer to [README of my Go solutions to year 2018](https://github.com/metalim/metalim.adventofcode.2018.go/blob/master/README.md#go-gotchas).

* ### To create custom iterators/generators, channels come in handy

  ```go
  myString := "a string with Unicode ⅀😃😉🎄❤"
  fmt.Println("\niterate over string, 1 Unicode rune at time. Note byteIndex is not continuous.")
  for byteIndex, aRune := range myString {
    fmt.Println(byteIndex, string(aRune))
  }

  mySlice := []string{"hello", "world"}
  fmt.Println("\niterate over slice")
  for index, value := range mySlice {
    fmt.Println(index, value)
  }

  myMap := map[string]int{"a": 15, "c": 19}
  fmt.Println("\niterate over map")
  for key, value := range myMap {
    fmt.Println(key, value)
  }

  myChan := make(chan string)
  go func() {
    myChan <- "Hello"
    myChan <- "World!"
    // generate values on demand. Then close the channel to end iteration.
    close(myChan)
  }()

  fmt.Println("\niterate over channel")
  for aString := range myChan {
    fmt.Println(aString)
  }
  ```

* ### inheritance is hard

  For example you want 2d grid base class with common methods, and implementation details in derived classes. In C++ you just create base class with virtual methods for impl.
  In Go you create base interface with method signatures, but can't add common methods. Instead you write functions, that recieve interface, and work as external functions.

## Puzzle inputs

Inputs are automatically retrieved from Advent of Code, provided you put at least one `<session-name>.cookie` into `inputs/` folder. To get the cookie, refer to website properties in your browser, after logging in into Advent of Code website.

## Log

Check out [LOG.md](LOG.md) for specifics of each task.
