package main

import "testing"

func TestBufDemo(t *testing.T) {
	bufDemo("我是一个好人")
	bufDemo("\x68\xE2\x82")
	bufDemo("\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98")
}

func TestRangeString(t *testing.T) {
	rangeString("我是一个好人")
}

func TestPrintfString(t *testing.T) {
	printfString()
}

var rs = []rune{'H', 'e', 'l', 'l', 'o', ' ', '世', '界'}

func BenchmarkFirst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runesToUTF8(rs)
	}
}

func BenchmarkSecond(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runesToUTF8Manual(rs)
	}
}

func BenchmarkThird(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runesToUTF8Manual2(rs)
	}
}
