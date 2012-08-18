package strings_ext

import (
	"fmt"
	"strings"
)

func ExampleHead() {
	//Haskell type signature (polymorphic): -
	//    head :: [a] -> a

	var str string = "golang"
	fmt.Println(string(Head(str))) //Head returns a rune

	//Output: g
}

func ExampleTail() {

	// Haskell type signature (polymorphic): -
	//    tail :: [a] -> [a]

	var str string = "golang"
	fmt.Println(Tail(str))

	//Output: olang
}

func ExampleLast() {
	//Haskell type signature (polymorphic): -
	//    last :: [a] -> a

	var str string = "google"
	fmt.Println(string(Last(str))) //Last returns a rune

	//Output: e
}

func ExampleInit() {
	//Haskell type signature (polymorphic): -
	//    init :: [a] -> [a]

	var str string = "google"
	fmt.Println(string(Init(str)))

	//Output: googl
}

func ExampleTake() {
	//Haskell type signature (polymorphic): -
	//    take :: Int -> [a] -> [a]

	var str string = "golang"
	fmt.Println(Take(2, str))

	//Output: go
}

func ExampleDrop() {
	// Haskell type signature (polymorphic): -
	//     drop :: Int -> [a] -> [a]

	var str string = "golang"
	fmt.Println(Drop(2, str))

	//Output: lang
}

func ExampleReverse() {
	//Haskell type signature (polymorphic): -
	//    reverse :: [a] -> [a]

	var str string = "golang"
	fmt.Println(Reverse(str))

	//Output: gnalog
}

func ExampleDistinct() {
	//Haskell type signature (polymorphic) - Haskell calls this function 'nub': -
	//    nub :: Eq a => [a] -> [a]

	var str string = "aaabbbcccdddeeefff"
	fmt.Println(Distinct(str))

	//Output: abcdef
}

func ExampleGroup() {
	//Haskell type signature (polymorphic): -
	//    group :: Eq a => [a] -> [[a]]

	var input string = "aaabbbccd"
	fmt.Println(Group(input))

	//Output: [aaa bbb cc d]
}

func ExampleGroupBy() {
	//Haskell type signature (polymorphic): -
	//    groupBy :: (a -> a -> Bool) -> [a] -> [[a]]

	var isDigit func(rune) bool = func(a rune) bool {
		return strings.ContainsRune("0123456789", a)
	}

	var input string = "02/08/2010"
	fmt.Println(GroupBy(func(a, b rune) bool { return (isDigit(a)) == (isDigit(b)) }, input))

	//Ouput: [02 / 08 / 2010]
}

func ExampleFilter() {
	//Haskell type signature (polymorphic): -
	//    filter :: (a -> Bool) -> [a] -> [a]

	var isNotPunctuation func(rune) bool = func(a rune) bool {
		return !strings.ContainsRune("!.,?:;-'\"", a)
	}

	fmt.Println(Filter(isNotPunctuation, "he said \"hello there!\"")) //strips all punctuation

	//Output: he said hello there
}

func ExampleDropWhile() {
	// Haskell type signature (polymorphic): -
	//		dropWhile :: (a -> Bool) -> [a] -> [a]

	var input string = "        Hello World"
	fmt.Println(DropWhile(func(a rune) bool { return a == ' ' }, input))

	//Output: Hello World
}

func ExampleTakeWhile() {
	//Haskell type signature (polymorphic): -
	//    takeWhile :: (a -> Bool) -> [a] -> [a]

	var input string = "aaaaAbbbbbccccc"
	fmt.Println(TakeWhile(func(a rune) bool { return a == 'a' }, input))

	//Output: aaaa
}

func ExampleSpan() {
	//Haskell type signature (polymorphic): -
	//    span :: (a -> Bool) -> [a] -> ([a], [a])

	var input string = "hello world"
	fmt.Println(Span(func(a rune) bool { return a != ' ' }, input))

	//Output: hello  world
}

func ExampleIsEmpty() {
	//Haskell type signature (polymorphic): -
	//    null :: [a] -> Bool

	var input string = ""
	fmt.Println(IsEmpty(input))

	//Output: true
}

func ExampleAll() {
	var isLowercase func(rune) bool = func(r rune) bool {
		return r >= 97 && r <= 122
	}

	var input string = "Google"
	fmt.Println(All(isLowercase, input))

	//Output: false
}
