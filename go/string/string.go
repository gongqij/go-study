package main

import (
	"bufio"
	"fmt"
	"strings"
	"unicode/utf8"
)

func rangeString(str string) {
	strArray := []rune(str)
	for index, runeValue := range strArray {
		fmt.Printf("%s starts at byte position %d\n", string(runeValue), index)
	}
	for i, w := 0, 0; i < len(str); i += w {
		runeValue, width := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}
}

func bufDemo(str string) {
	position := 0
	reader := bufio.NewReader(strings.NewReader(str))
	for {
		r, size, err := reader.ReadRune()
		if err != nil {
			break
		}
		fmt.Println(string(r), size, err)
		fmt.Printf("%#U starts at byte position %d\n", r, position)
		position = position + size
	}
}

func printfString() {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	fmt.Println("Println:")
	fmt.Println(sample)

	fmt.Println("Byte loop:")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Printf("\n")

	fmt.Printf("%x\n", sample)

	fmt.Printf("% x\n", sample)

	fmt.Printf("%q\n", sample)

	fmt.Printf("%+q\n", sample)
}

func runesToUTF8(rs []rune) []byte {
	return []byte(string(rs))
}

func runesToUTF8Manual(rs []rune) []byte {
	bs := make([]byte, len(rs)*utf8.UTFMax)

	count := 0
	for _, r := range rs {
		count += utf8.EncodeRune(bs[count:], r)
	}

	return bs[:count]
}

func runesToUTF8Manual2(rs []rune) []byte {
	size := 0
	for _, r := range rs {
		size += utf8.RuneLen(r)
	}

	bs := make([]byte, size)

	count := 0
	for _, r := range rs {
		count += utf8.EncodeRune(bs[count:], r)
	}

	return bs
}
