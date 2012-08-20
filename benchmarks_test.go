package strings_ext

import "testing"

var inputStr string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func BenchmarkHead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Head(inputStr)
	}
}

func BenchmarkTake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Take(12, inputStr)
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

func BenchmarkTakeWhile(b *testing.B) {
	var isLower func(rune)bool = func(r rune)bool { return r >= 97 && r <= 122 }

	for i := 0; i < b.N; i++ {
		TakeWhile(isLower, inputStr)
	}
}

func BenchmarkDropWhile(b *testing.B) {
	var isLower func(rune)bool = func(r rune)bool { return r >= 97 && r <= 122 }

	for i := 0; i < b.N; i++ {
		DropWhile(isLower, inputStr)
	}
}
