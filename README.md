# strings_ext

This is a small library that implements some of the functions from the ```Data.List``` package in Haskell that are only applicable to strings. The code is *probably* **very inefficient**, but I found it enjoyable as an intellectual excercise to port these useful functions over to Go.

## Install

	$ go get github.com/djhworld/strings_ext

## Rationale 

Because I enjoy programming in Haskell and want some of the standard library features found ```Data.List``` to be in Go

## But why only for strings? Why not use empty interfaces and make your code "generic" so it works with any datatype?

Personally, I don't see the use of the empty interface to be best practise when you want to guarantee type safety at runtime. I'm a strong advocate for compile time type checking and I feel it would not be appropriate to use ```interface{}``` as a half-way house to generics.

As soon as the Go team implement generics/parametric polymorphism into the language I'll almost certainly drop this library and create a more "generic" solution akin to ```Data.List```. For those looking to go down the ```interface{}``` route now, please see the [Seq](https://github.com/zot/seq/blob/release/seq.go) library by Bill Burdick

## Why no Map?

See [strings.Map](http://golang.org/pkg/strings/#Map) for the default implementation, although this version is NOT like Haskell's ```map``` in the sense that you can only input and output a string, no other type.


## Examples
	Head("hello") // returns 'h'
	Tail("hello") // returns "ello"
	Take(2, "golang") // returns "go"
	Drop(2, "golang") // returns "lang"
	TakeWhile(func(a rune) bool { return a == 'a' }, "aaaaAbbbbbccccc") // returns "aaaa"
	DropWhile(func(a rune) bool { return a == ' ' }, "        Hello World") // returns "Hello World"
	Reverse("golang") // returns "gnalog"
	Filter(func(a rune) bool { return !strings.ContainsRune("!.,?:;-'\"", a) }, "hello!?") // returns "hello"
	Span(func(a rune) bool { return a != ' ' }, "hello world") // returns ("hello"," world")
	Group("aaabbbccd") // returns []string{ "aaa", "bbb", "cc", "d" }
	Distinct("google") // returns "gole"
