package strings_ext

import (
	"github.com/bmizerany/assert"
	"strings"
	"testing"
	"unicode/utf8"
)

func FailWithLog(t *testing.T, log string) {
	t.Log(log)
	t.Fail()
}

// --------------------- HEAD ------------------------
func TestHead(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			FailWithLog(t, "Exception should NOT have been thrown!")
		} else {
			t.Log("No exception caught, we're OK!\n")
		}
	}()

	var input string = "hello"
	var expected rune = 'h'
	var actual rune = Head(input)

	assert.Equal(t, expected, actual)
}

func TestHeadWithEmpty(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Log("Exception was thrown successfully\n")
		} else {
			FailWithLog(t, "No exception was thrown!")
		}
	}()

	//should throw panic for empty
	Head("")
	t.Fail()
}

// --------------------- TAIL ------------------------
func TestTail(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			FailWithLog(t, "Exception should NOT have been thrown!")
		} else {
			t.Log("No exception caught, we're OK!\n")
		}
	}()

	var input string = "hello"
	var expected string = "ello"
	var actual string = Tail(input)

	assert.Equal(t, actual, expected)
}

func TestTailWithOneChar(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			FailWithLog(t, "Exception should NOT have been thrown!")
		} else {
			t.Log("No exception caught, we're OK!\n")
		}
	}()

	var input string = "h"
	var expected string = ""
	var actual string = Tail(input)

	assert.Equal(t, actual, expected)
}

func TestTailWithEmpty(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Log("Exception was thrown successfully\n")
		} else {
			FailWithLog(t, "No exception was thrown!")
		}
	}()

	//should throw panic for empty
	Tail("")
	t.Fail()
}

// --------------------- TAKE ------------------------
func TestTake(t *testing.T) {
	var input string = "testing"
	var expected string = "test"
	var actual string = Take(4, input)

	assert.Equal(t, actual, expected)
}

func TestTakeWithEmpty(t *testing.T) {
	var input string = ""
	var expected string = ""
	var actual string = Take(10, input)

	assert.Equal(t, actual, expected)
}

func TestTakeZero(t *testing.T) {
	var input string = "testing"
	var expected string = ""
	var actual string = Take(0, input)

	assert.Equal(t, actual, expected)
}

func TestTakeBelowZero(t *testing.T) {
	var input string = "testing"
	var expected string = ""
	var actual string = Take(-1, input)

	assert.Equal(t, actual, expected)
}

func TestTakeWithOneChar(t *testing.T) {
	var input string = "t"
	var expected string = "t"
	var actual string = Take(1, input)

	assert.Equal(t, actual, expected)
}

func TestTakeWithMoreThanStringLength(t *testing.T) {
	var input string = "test"
	var expected string = "test"
	var actual string = Take(500, input)

	assert.Equal(t, actual, expected)
}

func TestTakeWithStringLength(t *testing.T) {
	var input string = "test"
	var expected string = "test"
	var actual string = Take(utf8.RuneCountInString(input), input)

	assert.Equal(t, actual, expected)
}

// --------------------- DROP ------------------------
func TestDrop(t *testing.T) {
	var input string = "abcdef"
	var expected string = "def"
	var actual string = Drop(3, input)

	assert.Equal(t, actual, expected)
}

func TestDropWithEmpty(t *testing.T) {
	var input string = ""
	var expected string = ""
	var actual string = Drop(3, input)

	assert.Equal(t, actual, expected)
}

func TestDropWithOneChar(t *testing.T) {
	var input string = "h"
	var expected string = ""
	var actual string = Drop(1, input)

	assert.Equal(t, actual, expected)
}

func TestDropWithStringLength(t *testing.T) {
	var input string = "hello world"
	var expected string = ""
	var actual string = Drop(utf8.RuneCountInString(input), input)

	assert.Equal(t, actual, expected)
}

func TestDropWithMoreThanStringLength(t *testing.T) {
	var input string = "hello world"
	var expected string = ""
	var actual string = Drop(500, input)

	assert.Equal(t, actual, expected)
}

func TestDropWithZero(t *testing.T) {
	var input string = "hello"
	var expected string = "hello"
	var actual string = Drop(0, input)

	assert.Equal(t, actual, expected)
}

func TestDropWithBelowZero(t *testing.T) {
	var input string = "hello"
	var expected string = "hello"
	var actual string = Drop(-18, input)

	assert.Equal(t, actual, expected)
}

// --------------------- REVERSE ------------------------
func TestReverse(t *testing.T) {
	var input string = "testing"
	var expected string = "gnitset"
	var actual string = Reverse(input)

	assert.Equal(t, actual, expected)
}

func TestReverseWithEmpty(t *testing.T) {
	var input string = ""
	var expected string = ""
	var actual string = Reverse(input)

	assert.Equal(t, actual, expected)
}

func TestReverseWithOneChar(t *testing.T) {
	var input string = "a"
	var expected string = "a"
	var actual string = Reverse(input)

	assert.Equal(t, actual, expected)
}

func TestReverseWithPalindrome(t *testing.T) {
	var input string = "level"
	var expected string = "level"
	var actual string = Reverse(input)

	assert.Equal(t, actual, expected)
}

// --------------------- SPAN ------------------------
func TestSpan(t *testing.T) {
	var equalsA func(rune) bool = func(a rune) bool { return a == 'a' }
	var input string = "aaabbb"
	expected1, expected2 := "aaa", "bbb"
	actual1, actual2 := Span(equalsA, input)

	assert.Equal(t, actual1, expected1)
	assert.Equal(t, actual2, expected2)
}

func TestSpanWithEmpty(t *testing.T) {
	var equalsA func(rune) bool = func(a rune) bool { return a == 'a' }
	var input string = ""
	expected1, expected2 := "", ""
	actual1, actual2 := Span(equalsA, input)

	assert.Equal(t, actual1, expected1)
	assert.Equal(t, actual2, expected2)
}

func TestSpanWithOneChar(t *testing.T) {
	var equalsA func(rune) bool = func(a rune) bool { return a == 'a' }
	var input string = "a"
	expected1, expected2 := "a", ""
	actual1, actual2 := Span(equalsA, input)

	assert.Equal(t, actual1, expected1)
	assert.Equal(t, actual2, expected2)
}

func TestSpanWithNoMatches(t *testing.T) {
	var equalsA func(rune) bool = func(a rune) bool { return a == 'a' }
	var input string = "bbbccc"
	expected1, expected2 := "", "bbbccc"
	actual1, actual2 := Span(equalsA, input)

	assert.Equal(t, actual1, expected1)
	assert.Equal(t, actual2, expected2)
}

// --------------------- FILTER ------------------------
func TestFilter(t *testing.T) {
	var isNotPunctuation func(rune) bool = func(a rune) bool { return !strings.ContainsRune("!.,?:;-'\"", a) }
	var input string = "To be, or not to be. That is the question. Or is it?"
	var expected string = "To be or not to be That is the question Or is it"

	var actual string = Filter(isNotPunctuation, input)
	assert.Equal(t, actual, expected)
}

func TestFilterWithEmpty(t *testing.T) {
	var isNotPunctuation func(rune) bool = func(a rune) bool { return !strings.ContainsRune("!.,?:;-'\"", a) }
	var input string = ""
	var expected string = ""

	var actual string = Filter(isNotPunctuation, input)
	assert.Equal(t, actual, expected)
}

func TestFilterWithNoMatches(t *testing.T) {
	var equalsA func(rune) bool = func(a rune) bool { return a == 'a' }
	var input string = "Nothing found"
	var expected string = ""

	var actual string = Filter(equalsA, input)
	assert.Equal(t, actual, expected)
}

// --------------------- GROUP ------------------------
func TestGroup(t *testing.T) {
	var input string = "voodoo"
	var expected []string = []string{"v", "oo", "d", "oo"}

	var actual []string = Group(input)
	assert.Equal(t, actual, expected)
}

func TestGroupWithEmpty(t *testing.T) {
	var input string = ""
	var expected []string = []string{}

	var actual []string = Group(input)
	assert.Equal(t, actual, expected)
}

func TestGroupWithOneChar(t *testing.T) {
	var input string = "v"
	var expected []string = []string{"v"}

	var actual []string = Group(input)
	assert.Equal(t, actual, expected)
}

// --------------------- GROUP BY ------------------------
func TestGroupBy(t *testing.T) {
	var isDigit func(rune) bool = func(a rune) bool {
		return strings.ContainsRune("0123456789", a)
	}

	var input string = "02/08/2010"
	var expected []string = []string{"02", "/", "08", "/", "2010"}

	var actual []string = GroupBy(func(a, b rune) bool { return (isDigit(a)) == (isDigit(b)) }, input)
	assert.Equal(t, actual, expected)
}

func TestGroupByWithEmpty(t *testing.T) {
	var isDigit func(rune) bool = func(a rune) bool {
		return strings.ContainsRune("0123456789", a)
	}

	var input string = ""
	var expected []string = []string{}

	var actual []string = GroupBy(func(a, b rune) bool { return (isDigit(a)) == (isDigit(b)) }, input)
	assert.Equal(t, actual, expected)
}

func TestGroupByWithOneChar(t *testing.T) {
	var isDigit func(rune) bool = func(a rune) bool {
		return strings.ContainsRune("0123456789", a)
	}

	var input string = "0"
	var expected []string = []string{"0"}

	var actual []string = GroupBy(func(a, b rune) bool { return (isDigit(a)) == (isDigit(b)) }, input)
	assert.Equal(t, actual, expected)
}

func TestGroupByWithOneGroup(t *testing.T) {
	var isDigit func(rune) bool = func(a rune) bool {
		return strings.ContainsRune("0123456789", a)
	}

	var input string = "999"
	var expected []string = []string{"999"}

	var actual []string = GroupBy(func(a, b rune) bool { return (isDigit(a)) == (isDigit(b)) }, input)
	assert.Equal(t, actual, expected)
}

func TestGroupByWithNoGroups(t *testing.T) {
	var isDigit func(rune) bool = func(a rune) bool {
		return strings.ContainsRune("0123456789", a)
	}

	var input string = "hello world"
	var expected []string = []string{"hello world"}

	var actual []string = GroupBy(func(a, b rune) bool { return (isDigit(a)) == (isDigit(b)) }, input)
	assert.Equal(t, actual, expected)
}

// --------------------- DISTINCT ------------------------
func TestDistinct(t *testing.T) {
	var input string = "GOOGLE"
	var expected string = "GOLE"
	var actual string = Distinct(input)
	assert.Equal(t, actual, expected)
}

func TestDistinctWithEmpty(t *testing.T) {
	var input string = ""
	var expected string = ""
	var actual string = Distinct(input)
	assert.Equal(t, actual, expected)
}

func TestDistinctWithAllDistinct(t *testing.T) {
	var input string = "great"
	var expected string = "great"
	var actual string = Distinct(input)
	assert.Equal(t, actual, expected)
}

// --------------------- LAST ------------------------
func TestLast(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			FailWithLog(t, "Exception should NOT have been thrown!")
		} else {
			t.Log("No exception caught, we're OK!\n")
		}
	}()

	var input string = "hello"
	var expected rune = 'o'
	var actual rune = Last(input)

	assert.Equal(t, expected, actual)
}

func TestLastWithEmpty(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Log("Exception was thrown successfully\n")
		} else {
			FailWithLog(t, "No exception was thrown!")
		}
	}()

	//should throw panic for empty
	Last("")
	t.Fail()
}

// --------------------- INIT ------------------------
func TestInit(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			FailWithLog(t, "Exception should NOT have been thrown!")
		} else {
			t.Log("No exception caught, we're OK!\n")
		}
	}()

	var input string = "hello"
	var expected string = "hell"
	var actual string = Init(input)

	assert.Equal(t, actual, expected)
}

func TestInitWithEmpty(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Log("Exception was thrown successfully\n")
		} else {
			FailWithLog(t, "No exception was thrown!")
		}
	}()

	//should throw panic for empty
	Init("")
	t.Fail()
}

func TestInitWithOneChar(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			FailWithLog(t, "Exception should NOT have been thrown!")
		} else {
			t.Log("No exception caught, we're OK!\n")
		}
	}()

	var input string = "h"
	var expected string = ""
	var actual string = Init(input)

	assert.Equal(t, actual, expected)
}

// --------------------- ISEMPTY ------------------------
func TestIsEmptyTrue(t *testing.T) {
	var input string = ""
	var expected bool = true
	var actual = IsEmpty(input)

	assert.Equal(t, actual, expected)
}

func TestIsEmptyFalse(t *testing.T) {
	var input string = "daniel"
	var expected bool = false
	var actual = IsEmpty(input)

	assert.Equal(t, actual, expected)
}

// --------------------- ALL ------------------------
func TestAll(t *testing.T) {
	var isLowercase func(rune) bool = func(r rune) bool {
		return r >= 97 && r <= 122
	}

	var input string = "aaa"
	var expected bool = true
	var actual = All(isLowercase, input)

	assert.Equal(t, actual, expected)
}

func TestAllWithEmpty(t *testing.T) {
	var isLowercase func(rune) bool = func(r rune) bool {
		return r >= 97 && r <= 122
	}

	var input string = ""
	var expected bool = true
	var actual = All(isLowercase, input)

	assert.Equal(t, actual, expected)
}


func TestAllFalse(t *testing.T) {
	var isLowercase func(rune) bool = func(r rune) bool {
		return r >= 97 && r <= 122
	}

	var input string = "aaaA"
	var expected bool = false
	var actual = All(isLowercase, input)

	assert.Equal(t, actual, expected)
}

// --------------------- TAKEWHILE ------------------------
func TestTakeWhile(t *testing.T) {
	var isA func(rune)bool = func(r rune)bool { return r == 'a' }
	var input string = "aaaaaaaaaaaaaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbbbbbbbbbbbb"
	var expected string = "aaaaaaaaaaaaaaaaaaaaaaaaaaa"
	var actual string = TakeWhile(isA, input)

	assert.Equal(t, actual, expected)
}

func TestTakeWhileWithEmpty(t *testing.T) {
	var isA func(rune)bool = func(r rune)bool { return r == 'a' }
	var input string = ""
	var expected string = ""
	var actual string = TakeWhile(isA, input)

	assert.Equal(t, actual, expected)
}

func TestTakeWhileWithOneChar(t *testing.T) {
	var isA func(rune)bool = func(r rune)bool { return r == 'a' }
	var input string = "a"
	var expected string = "a"
	var actual string = TakeWhile(isA, input)

	assert.Equal(t, actual, expected)
}

// --------------------- DROPWHILE ------------------------
func TestDropWhile(t *testing.T) {
	var isA func(rune)bool = func(r rune)bool { return r == 'a' }
	var input string = "aaaaaaaaaaaaaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbbbbbbbbbbbb"
	var expected string = "bbbbbbbbbbbbbbbbbbbbbbbbbbbbb"
	var actual string = DropWhile(isA, input)

	assert.Equal(t, actual, expected)
}

func TestDropWhileWithEmpty(t *testing.T) {
	var isA func(rune)bool = func(r rune)bool { return r == 'a' }
	var input string = ""
	var expected string = ""
	var actual string = DropWhile(isA, input)

	assert.Equal(t, actual, expected)
}

func TestDropWhileWithOneChar(t *testing.T) {
	var isA func(rune)bool = func(r rune)bool { return r == 'a' }
	var input string = "b"
	var expected string = "b"
	var actual string = DropWhile(isA, input)

	assert.Equal(t, actual, expected)
}
