# testig

Go Test Helpers

## What?

Helper functions for Go testing.  May they help you thoroughly test your
Go code!

The name `testig` is a pun in two languages, as well as an experiment in
misspelling.

## Why?

I find myself with shared testing code across projects, which I do not wish
to have scattered about like that.

Note that some of this stuff may eventually be merged into other packages
(see below); or may be marked *deprecated as dumb* if I find something much
better.

Also please note that *test coverage is crap here* because testing the tester
is kinda hard.

## Who?

(c) 2016 Kevin Frost; BSD license (cf. the `LICENSE` file).

This package leans *heavily* on the [assert][1] package by Mat Ryer and Tyler
Bunnel (et al.), without which I probably wouldn't have even stuck with Go
long enough to care about this. Thanks [guys][2]!

[1]: https://godoc.org/github.com/stretchr/testify/assert
[2]: https://github.com/stretchr/testify/graphs/contributors
