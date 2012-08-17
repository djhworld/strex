/*
Package that adds a few extra features to the strings package. 

These functions are very high level and may not be very efficient. Most of the functions take inspiration from the Data.List package found in the Haskell programming language

Examples: -
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
*/
package strings_ext

const EMPTY_STR string = ""

/*
Extract the first element of a string, which must be non-empty.

Haskell type signature (polymorphic): -
	head :: [a] -> a
*/
func Head(s string) rune {
	if len(s) == 0 {
		panic("empty list")
	}

	runes := []rune(s)
	return runes[0]
}

/*
Extract the remaining string after the head of a string, which must be non-empty.

Haskell type signature (polymorphic): -
	tail :: [a] -> [a]
*/
func Tail(s string) string {
	if len(s) == 0 {
		panic("empty list")
	}

	runes := []rune(s)
	xs := runes[1:]
	return string(xs)
}

/*
Take N, applied to a string S, returns the prefix of S of length N, or S itself if N > len(S)

Haskell type signature (polymorphic): -
	take :: Int -> [a] -> [a]
*/
func Take(i int, s string) string {
	if i <= 0 || len(s) == 0 {
		return EMPTY_STR
	}

	x := string(Head(s))
	xs := Tail(s)
	return x + Take(i-1, xs)
}

/*
Drop(N,S) returns the suffix of S after the first N elements, or "" if N > len(S)

Haskell type signature (polymorphic): -
	drop :: Int -> [a] -> [a]
*/
func Drop(i int, s string) string {
	if i <= 0 {
		return s
	} else if len(s) == 0 {
		return s
	}

	xs := Tail(s)
	return Drop(i-1, xs)
}

/* 
TakeWhile, applied to a predicate P and a string S, returns the longest prefix (possibly empty) of S of elements that satisfy P

Haskell type signature (polymorphic): -
    takeWhile :: (a -> Bool) -> [a] -> [a]
*/
func TakeWhile(p func(rune) bool, s string) string {
	if len(s) == 0 {
		return EMPTY_STR
	}

	x := Head(s)
	xs := Tail(s)

	if !p(x) {
		return EMPTY_STR
	}

	return string(x) + TakeWhile(p, xs)
}

/* 
DropWhile(P,S) returns the suffix remaining after TakeWhile(P,S)

Haskell type signature (polymorphic): -
    dropWhile :: (a -> Bool) -> [a] -> [a]
*/
func DropWhile(p func(rune) bool, s string) string {
	if len(s) == 0 {
		return EMPTY_STR
	}
	x := Head(s)
	xs := Tail(s)
	if !p(x) {
		return s
	}
	return DropWhile(p, xs)
}

/*
Reverse(S) returns the string S in reverse order

Haskell type signature (polymorphic): -
	reverse :: [a] -> [a]
*/
func Reverse(s string) string {
	if len(s) == 0 {
		return EMPTY_STR
	}
	x := Head(s)
	xs := Tail(s)
	return Reverse(xs) + string(x)
}

/*
Filter, applied to a predicate and a string, returns a string of characters (runes) that satisfy the predicate

Haskell type signature (polymorphic): -
	filter :: (a -> Bool) -> [a] -> [a]
*/
func Filter(p func(rune) bool, s string) string {
	if len(s) == 0 {
		return EMPTY_STR
	}

	x := Head(s)
	xs := Tail(s)

	if !p(x) {
		return Filter(p, xs)
	}

	return string(x) + Filter(p, xs)
}

/*
Span, applied to a predicate P and a string S, returns two strings where the first string is longest prefix (possibly empty) of S of characters (runes) that satisfy P and the second string is the remainder of the string 

Haskell type signature (polymorphic): -
    span :: (a -> Bool) -> [a] -> ([a], [a])
*/
func Span(p func(rune) bool, s string) (string, string) {
	return TakeWhile(p, s), DropWhile(p, s)
}

/*
The Group function takes a string and returns a slice of strings such that the concatenation of the result is equal to the argument. Moreover, each sublist in the result contains only equal elements.

Haskell type signature (polymorphic): -
	group :: Eq a => [a] -> [[a]]
*/
func Group(s string) []string {
	return GroupBy(func(a, b rune) bool { return a == b }, s)
}

/*
The GroupBy function is the non-overloaded version of Group.

Haskell type signature (polymorphic): -
	groupBy :: (a -> a -> Bool) -> [a] -> [[a]]

Example: -
	var isDigit func(rune) bool = func(a rune) bool {
		return strings.ContainsRune("0123456789", a)
	}

	var input string = "02/08/2010"
	GroupBy(func(a, b rune) bool { return (isDigit(a)) == (isDigit(b)) }, input) // returns []string{"02", "/", "08", "/", "2010"}
*/
func GroupBy(p func(rune, rune) bool, s string) []string {
	return reverseStringSlice(groupBy_(p, s))
}

/*
The Distinct function removes duplicate elements from a string. In particular, it keeps only the first occurrence of each element. 

Haskell type signature (polymorphic) - Haskell calls this function 'nub': -
	nub :: Eq a => [a] -> [a]
*/
func Distinct(s string) string {
	if len(s) == 0 {
		return EMPTY_STR
	}

	x := Head(s)
	xs := Tail(s)

	return string(x) + Distinct(Filter(func(y rune) bool { return x != y }, xs))
}

// Private methods 
func groupBy_(p func(rune, rune) bool, s string) []string {
	if len(s) == 0 {
		return []string{}
	}
	x := Head(s)
	xs := Tail(s)
	ys, zs := Span(func(a rune) bool { return p(a, x) }, xs)
	return append(groupBy_(p, zs), string(x)+ys)
}

func reverseStringSlice(s []string) []string {
	if len(s) == 0 {
		return []string{}
	}

	x := s[0]
	xs := s[1:]
	return append(reverseStringSlice(xs), x)
}
