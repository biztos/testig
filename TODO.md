# testig TODO

(Arguably at least.)

## Accept a Regexp *or* a string for AssertPanicsRegexp

This is generally a good idea (cf. `testify/assert`) but it's not a priority,
as I personally never use that form.

## A utility to rig up placeholder test functions

* Read all .go files in the given directory (recurse? probably not)
* First make sure each one is "go fmt" safe, otherwise fail.
* For every method, set up a placeholder test function.
* OPTION: assume `assert` is in use.
* If the method returns an error, create two functions: _Success and _Failure
* Allow override of default style which is `Test_FuncName_Detail`
* Create `_test.go` files if needed, otherwise append to them.
* Don't re-create anything: if `Test_FuncName` or `TestFuncName` exists, pass.

## Recorders for various reader/writer interfaces

* The goal being to minimize interaction with the file system.
* Check what I did for kisipar.

## Consider vendoring in "assert"

Just on principle.  I would hope to always be compatible with their latest
version, since I use it everywhere, but considering what's happened lately
with a lot of NodeJS `npm` dependencies, one ought not to assume.

**What about the license?** That is, what *should* I do differently if I'm
redistributing all that code in `vendor`?  Anything?  (Pretty sure I have this
covered correctly in `frostedmd` and will do the same in `kisipar`.)