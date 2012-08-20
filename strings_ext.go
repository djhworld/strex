/*
Package that adds a few extra features to the strings package. 

These functions are very high level and may not be very efficient. Most of the 
functions take inspiration from the Data.List package found in the Haskell programming 
language
*/
package strings_ext

import (
	"unicode/utf8"
)

//Head returns the first rune of s which must be non-empty
func Head(s string) rune {
	if s == "" {
		panic("empty list")
	}
	r, _ := utf8.DecodeRuneInString(s)
	return r
}

//Tail returns the the remainder of s minus the first rune of s, which must be non-empty
func Tail(s string) string {
	if s == "" {
		panic("empty list")
	}

	_, sz := utf8.DecodeRuneInString(s)
	return s[sz:]
}

//Removed. Recursive solution is not performant
//func Take(n int, s string) string {
//	if n <= 0 || s == "" {
//		return ""
//	}
//
//	x := string(Head(s))
//	xs := Tail(s)
//	return x + Take(n-1, xs)
//}

//Take returns the n rune prefix of s or s itself if n > len([]rune(s))
func Take(n int, s string) string {
	if n <= 0 || s == "" {
		return ""
	} else if n > utf8.RuneCountInString(s) {
		return s
	}

	//TODO: Deal with rune encoding errors
	byteLen := 0
	for i := 0; i < n; i++ {
		_, runeByteLen := utf8.DecodeRuneInString(s[byteLen:])
		byteLen += runeByteLen
	}

	return s[:byteLen]
}

//Removed. Recursive solution is not performant
//BenchmarkDrop	 5000000	       666 ns/op
//BenchmarkDrop	10000000	       283 ns/op - non recursive
//func Drop(n int, s string) string {
//	if n <= 0 || s == "" {
//		return s
//	}
//
//	xs := Tail(s)
//	return Drop(n-1, xs)
//}

//Drop returns the suffix of s after the first n runes, or "" if n > len([]rune(s))
func Drop(n int, s string) string {
	if n <= 0 || s == "" {
		return s
	} else if n > utf8.RuneCountInString(s) {
		return ""
	}

	//TODO: Deal with rune encoding errors
	byteLen := 0
	for i := 0; i < n; i++ {
		_, runeByteLen := utf8.DecodeRuneInString(s[byteLen:])
		byteLen += runeByteLen
	}

	return s[byteLen:]
}

// recursive solution is not very performant
// BenchmarkTakeWhile	  100000	     13097 ns/op
// BenchmarkTakeWhile	 1000000	      1839 ns/op
//func TakeWhile(p func(rune) bool, s string) string {
//	if s == "" {
//		return ""
//	}
//
//	x := Head(s)
//	xs := Tail(s)
//
//	if !p(x) {
//		return ""
//	}
//
//	return string(x) + TakeWhile(p, xs)
//}

//TakeWhile, applied to a predicate p and a string s, returns the longest 
// prefix (possibly empty) of s of elements that satisfy p
func TakeWhile(p func(rune) bool, s string) string {
	if s == "" {
		return ""
	}

	rs := []rune(s)
	takeCount := 0

	for i := 0; i < len(rs); i++ {

		if !p(rs[i]) {
			break
		}

		takeCount++
	}

	return Take(takeCount, s)
}

// not performant
//func DropWhile(p func(rune) bool, s string) string {
//	if s == "" {
//		return ""
//	}
//	x := Head(s)
//	xs := Tail(s)
//	if !p(x) {
//		return s
//	}
//	return DropWhile(p, xs)
//}

//DropWhile returns the suffix remaining after TakeWhile
func DropWhile(p func(rune) bool, s string) string {
	if s == "" {
		return ""
	}

	rs := []rune(s)
	dropCount := 0

	for i := 0; i < len(rs); i++ {

		if !p(rs[i]) {
			break
		}

		dropCount++
	}

	return Drop(dropCount, s)
}

//this is very inefficient
// BenchmarkReverse	   50000	     52679 ns/op
//func Reverse(s string) string {
//	if s == "" {
//		return ""
//	}
//	x := Head(s)
//	xs := Tail(s)
//	return Reverse(xs) + string(x)
//}

//Reverse returns the string s in reverse order
func Reverse(s string) string {
	if s == "" {
		return ""
	}

	rs := []rune(s)
	rss := make([]rune, len(rs), len(rs))

	o := 0
	for i := len(rs) - 1; i >= 0; i-- {
		rss[o] = rs[i]
		o++
	}

	return string(rss)
}

// Filter, applied to a predicate and a string, returns a string of characters 
// (runes) that satisfy the predicate
func Filter(p func(rune) bool, s string) string {
	if s == "" {
		return ""
	}

	x := Head(s)
	xs := Tail(s)

	if !p(x) {
		return Filter(p, xs)
	}

	return string(x) + Filter(p, xs)
}

//Span, applied to a predicate p and a string s, returns two strings where the 
//first string is longest prefix (possibly empty) of s of characters (runes) that 
//satisfy p and the second string is the remainder of the string 
func Span(p func(rune) bool, s string) (string, string) {
	return TakeWhile(p, s), DropWhile(p, s)
}

//Group takes a string and returns a slice of strings such 
//that the concatenation of the result is equal to the argument.
//Moreover, each sublist in the result contains only equal elements.
func Group(s string) []string {
	return GroupBy(func(a, b rune) bool { return a == b }, s)
}

//GroupBy is the non-overloaded version of Group.
func GroupBy(p func(rune, rune) bool, s string) []string {
	return reverseStringSlice(groupBy_(p, s))
}

// Distinct removes duplicate elements from a string. 
// In particular, it keeps only the first occurrence of each element. 
func Distinct(s string) string {
	if s == "" {
		return ""
	}

	x := Head(s)
	xs := Tail(s)

	return string(x) + Distinct(Filter(func(y rune) bool { return x != y }, xs))
}

//Last returns the last rune in a string s, which must be non-empty.
func Last(s string) rune {
	if s == "" {
		panic("empty list")
	}

	r, _ := utf8.DecodeLastRuneInString(s)
	return r
}

//Init returns all the elements of s except the last one. The string must 
//be non-empty.
func Init(s string) string {
	if s == "" {
		panic("empty list")
	}

	_, sz := utf8.DecodeRuneInString(s)
	c := utf8.RuneCountInString(s)
	return s[:(sz*c)-sz]
}

//IsEmpty tests whether the string s is empty
func IsEmpty(s string) bool {
	return s == ""
}

//All applied to a predicate p and a string s, determines if all elements of
//s satisfy p
func All(p func(rune) bool, s string) bool {
	if IsEmpty(s) {
		return true
	}

	var result bool = false

	runes := []rune(s)
	for _, v := range runes {
		result = p(v)
		if result == false {
			break
		}
	}

	return result
}

// Private methods 
func groupBy_(p func(rune, rune) bool, s string) []string {
	if s == "" {
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
