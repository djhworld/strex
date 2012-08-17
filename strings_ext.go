/*
Package that adds a few extra features to the strings package. 

These functions are very high level and may not be very efficient. Most of the 
functions take inspiration from the Data.List package found in the Haskell programming 
language
*/
package strings_ext

const EMPTY_STR string = ""

//Extract the first element of a string, which must be non-empty.
func Head(s string) rune {
	if len(s) == 0 {
		panic("empty list")
	}

	runes := []rune(s)
	return runes[0]
}

//Extract the remaining string after the head of a string, which must be non-empty.
func Tail(s string) string {
	if len(s) == 0 {
		panic("empty list")
	}

	runes := []rune(s)
	xs := runes[1:]
	return string(xs)
}

//Take N, applied to a string S, returns the prefix of S of length N, or S itself if N > len(S)
func Take(i int, s string) string {
	if i <= 0 || len(s) == 0 {
		return EMPTY_STR
	}

	x := string(Head(s))
	xs := Tail(s)
	return x + Take(i-1, xs)
}

//Drop(N,S) returns the suffix of S after the first N elements, or "" if N > len(S)
func Drop(i int, s string) string {
	if i <= 0 {
		return s
	} else if len(s) == 0 {
		return s
	}

	xs := Tail(s)
	return Drop(i-1, xs)
}

//TakeWhile, applied to a predicate P and a string S, returns the longest prefix (possibly empty) of S of elements that satisfy P
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

//DropWhile(P,S) returns the suffix remaining after TakeWhile(P,S)
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

//Reverse(S) returns the string S in reverse order
func Reverse(s string) string {
	if len(s) == 0 {
		return EMPTY_STR
	}
	x := Head(s)
	xs := Tail(s)
	return Reverse(xs) + string(x)
}

// Filter, applied to a predicate and a string, returns a string of characters (runes) that satisfy the predicate
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

//Span, applied to a predicate P and a string S, returns two strings where the 
//first string is longest prefix (possibly empty) of S of characters (runes) that 
//satisfy P and the second string is the remainder of the string 
func Span(p func(rune) bool, s string) (string, string) {
	return TakeWhile(p, s), DropWhile(p, s)
}

//The Group function takes a string and returns a slice of strings such that the concatenation of the result is equal to the argument.
//Moreover, each sublist in the result contains only equal elements.
func Group(s string) []string {
	return GroupBy(func(a, b rune) bool { return a == b }, s)
}

//The GroupBy function is the non-overloaded version of Group.
func GroupBy(p func(rune, rune) bool, s string) []string {
	return reverseStringSlice(groupBy_(p, s))
}

// The Distinct function removes duplicate elements from a string. In particular, it keeps only the first occurrence of each element. 
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
