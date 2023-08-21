package main

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func encode(utf32 []rune) []byte {
	var utf8 []byte
	for _, codepoint := range utf32 {
		if codepoint <= 0x7F {
			utf8 = append(utf8, byte(codepoint))
		} else if codepoint <= 0x7FF {
			utf8 = append(utf8, byte(0xC0|(codepoint>>6)))
			utf8 = append(utf8, byte(0x80|(codepoint&0x3F)))
		} else if codepoint <= 0xFFFF {
			utf8 = append(utf8, byte(0xE0|(codepoint>>12)))
			utf8 = append(utf8, byte(0x80|((codepoint>>6)&0x3F)))
			utf8 = append(utf8, byte(0x80|(codepoint&0x3F)))
		} else {
			utf8 = append(utf8, byte(0xF0|(codepoint>>18)))
			utf8 = append(utf8, byte(0x80|((codepoint>>12)&0x3F)))
			utf8 = append(utf8, byte(0x80|((codepoint>>6)&0x3F)))
			utf8 = append(utf8, byte(0x80|(codepoint&0x3F)))
		}
	}
	return utf8
}

func decode(utf8 []byte) []rune {
	var utf32 []rune
	for i := 0; i < len(utf8); {
		if utf8[i]&0x80 == 0 {
			utf32 = append(utf32, rune(utf8[i]))
			i++
		} else if utf8[i]&0xE0 == 0xC0 {
			utf32 = append(utf32, rune(((int(utf8[i])&0x1F)<<6)|(int(utf8[i+1])&0x3F)))
			i += 2
		} else if utf8[i]&0xF0 == 0xE0 {
			utf32 = append(utf32, rune(((int(utf8[i])&0x0F)<<12)|((int(utf8[i+1])&0x3F)<<6)|(int(utf8[i+2])&0x3F)))
			i += 3
		} else {
			utf32 = append(utf32, rune(((int(utf8[i])&0x07)<<18)|((int(utf8[i+1])&0x3F)<<12)|((int(utf8[i+2])&0x3F)<<6)|(int(utf8[i+3])&0x3F)))
			i += 4
		}
	}
	return utf32
}

func TestEncode(t *testing.T) {
	utf32 := []rune{0x0041, 0x30A2, 0x1F60A}
	expected := []byte{0x41, 0xE3, 0x82, 0xA2, 0xF0, 0x9F, 0x98, 0x8A}
	result := encode(utf32)
	if !bytes.Equal(result, expected) {
		t.Errorf("encode(%v) = %v, expected %v", utf32, result, expected)
	}
}

func TestDecode(t *testing.T) {
	utf8 := []byte{0x41, 0xE3, 0x82, 0xA2, 0xF0, 0x9F, 0x98, 0x8A}
	expected := []rune{0x0041, 0x30A2, 0x1F60A}
	result := decode(utf8)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("decode(%v) = %v, expected %v", utf8, result, expected)
	}
}

func main() {
	fmt.Println("Running tests for encode and decode functions:")
	testing.Main(func(pat, str string) (bool, error) {
		return true, nil
	}, []testing.InternalTest{
		{Name: "TestEncode", F: TestEncode},
		{Name: "TestDecode", F: TestDecode},
	}, nil, nil)
}
