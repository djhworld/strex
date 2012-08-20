package strings_ext

import "testing"

var inputStr string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func BenchmarkHead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Head(inputStr)
	}
}

func BenchmarkTail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Tail(inputStr)
	}
}

func BenchmarkInit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Init(inputStr)
	}
}

func BenchmarkTake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Take(12, inputStr)
	}
}

func BenchmarkLast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Last(inputStr)
	}
}

func BenchmarkDrop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Drop(12, inputStr)
	}
}

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse(inputStr)
	}
}

func BenchmarkDistinct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Distinct(inputStr)
	}
}

func BenchmarkGroup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Group(inputStr)
	}
}

func BenchmarkFilter(b *testing.B) {
	var isLower func(rune) bool = func(r rune) bool { return r >= 97 && r <= 122 }
	for i := 0; i < b.N; i++ {
		Filter(isLower, inputStr)
	}
}

func BenchmarkAll(b *testing.B) {
	var isLower func(rune) bool = func(r rune) bool { return r >= 97 && r <= 122 }
	for i := 0; i < b.N; i++ {
		All(isLower, inputStr)
	}
}

func BenchmarkTakeWhile(b *testing.B) {
	var isLower func(rune) bool = func(r rune) bool { return r >= 97 && r <= 122 }

	for i := 0; i < b.N; i++ {
		TakeWhile(isLower, inputStr)
	}
}

func BenchmarkDropWhile(b *testing.B) {
	var isLower func(rune) bool = func(r rune) bool { return r >= 97 && r <= 122 }

	for i := 0; i < b.N; i++ {
		DropWhile(isLower, inputStr)
	}
}

func BenchmarkSpan(b *testing.B) {
	var isLower func(rune) bool = func(r rune) bool { return r >= 97 && r <= 122 }

	for i := 0; i < b.N; i++ {
		Span(isLower, inputStr)
	}
}

