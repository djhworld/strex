#strex

This is a small library that implements some of the functions from the `Data.List` package in Haskell that are only applicable to strings. 

##Install

	$ go get github.com/djhworld/strings_ext

##Rationale 

Because I enjoy programming in Haskell and want some of the standard library features found `Data.List` to be in Go, plus I found it an enjoyable intellectual exercise too. 

## But why only for strings? Why not use empty interfaces and make your code "generic" so it works with any datatype?

Personally, I don't see the use of the empty interface to be best practise when you want to guarantee type safety at runtime. I'm a strong advocate for compile time type checking and I feel it would not be appropriate to use `interface{}` as a half-way house to generics.

As soon as the Go team implement generics/parametric polymorphism into the language I'll almost certainly drop this library and create a more "generic" solution akin to `Data.List`. For those looking to go down the `interface{}` route now, please see the [Seq](https://github.com/zot/seq/blob/release/seq.go) library by Bill Burdick

##Why no Map?

See [strings.Map](http://golang.org/pkg/strings/#Map) for the default implementation, although this version is NOT like Haskell's `map` in the sense that you can only input and output a string, no other type.


##Documentation
[http://go.pkgdoc.org/github.com/djhworld/strex](http://go.pkgdoc.org/github.com/djhworld/strex)

# With thanks

To the people on reddit [who helped me with this code.](http://www.reddit.com/r/golang/comments/yfurz/made_a_small_library_critique_welcome/) Originally, I implemented all the functions in a Haskell like style with plenty of recursion and other inefficiencies that don't work so well in a non tail-call optimised language. With the help from the users such as ["rogpeppe"](https://github.com/rogpeppe) I was able to tighten up the package a lot with much faster code.  
