package test

import (
	"base64/base64"
	"testing"
)

func TestTranslateBlockWithThreeBytes(t *testing.T) {
	bytes := []byte{
		0b10101010,
		0b01010101,
		0b10101010,
	}
	expect := []byte{
		0b00101010,
		0b00100101,
		0b00010110,
		0b00101010,
	}
	translated, _ := base64.TranslateBlock(bytes)

	for index := range expect {
		if translated[index] != expect[index] {
			t.Fatalf("Expected %06b, got %06b.", expect[index], translated[index])
		}
	}
}

func TestTranslateBlockWithTwoBytes(t *testing.T) {
	bytes := []byte{
		0b10101010,
		0b01010101,
	}
	expect := []byte{
		0b00101010,
		0b00100101,
		0b00010100,
	}
	translated, _ := base64.TranslateBlock(bytes)

	for index := range expect {
		if translated[index] != expect[index] {
			t.Fatalf("Expected %06b, got %06b.", expect[index], translated[index])
		}
	}
}

func TestTranslateBlockWithOneByte(t *testing.T) {
	bytes := []byte{
		0b10101010,
	}
	expect := []byte{
		0b00101010,
		0b00100000,
	}
	translated, _ := base64.TranslateBlock(bytes)

	for index := range expect {
		if translated[index] != expect[index] {
			t.Fatalf("Expected %06b, got %06b.", expect[index], translated[index])
		}
	}
}

func TestTranslateBlockWithUnsupportedLength(t *testing.T) {
	bytes := []byte{}
	_, err := base64.TranslateBlock(bytes)

	if err == nil {
		t.Fatalf("Expected error, got nil.")
	}
}

func TestDetranslateBlockWithFour(t *testing.T) {
	bytes := []byte{
		0b00101010,
		0b00100101,
		0b00010110,
		0b00101010,
	}
	expect := []byte{
		0b10101010,
		0b01010101,
		0b10101010,
	}
	detranslated, _ := base64.DetranslateBlock(bytes)

	for index := range expect {
		if detranslated[index] != expect[index] {
			t.Fatalf("Expected %06b, got %06b.", expect[index], detranslated[index])
		}
	}
}

func TestDetranslateBlockWithThreeBytes(t *testing.T) {
	bytes := []byte{
		0b00101010,
		0b00100101,
		0b00010100,
	}
	expect := []byte{
		0b10101010,
		0b01010101,
	}
	detranslated, _ := base64.DetranslateBlock(bytes)

	for index := range expect {
		if detranslated[index] != expect[index] {
			t.Fatalf("Expected %06b, got %06b.", expect[index], detranslated[index])
		}
	}
}

func TestDetranslateBlockWithTwoBytes(t *testing.T) {
	bytes := []byte{
		0b00101010,
		0b00100000,
	}
	expect := []byte{
		0b10101010,
	}
	detranslated, _ := base64.DetranslateBlock(bytes)

	for index := range expect {
		if detranslated[index] != expect[index] {
			t.Fatalf("Expected %06b, got %06b.", expect[index], detranslated[index])
		}
	}
}

func TestDetranslateBlockWithUnsupportedLength(t *testing.T) {
	bytes := []byte{0b0}
	_, err := base64.DetranslateBlock(bytes)

	if err == nil {
		t.Fatalf("Expected error, got nil.")
	}
}

func TestTranslate(t *testing.T) {
	bytes := []byte{
		0b10101010,
		0b01010101,
		0b10101010,
		0b10101010,
	}
	expect := []byte{
		0b00101010,
		0b00100101,
		0b00010110,
		0b00101010,
		0b00101010,
		0b00100000,
	}
	translated, _ := base64.Translate(bytes)

	for index := range expect {
		if translated[index] != expect[index] {
			t.Fatalf("[%d] Expected %06b, got %06b.", index, expect[index], translated[index])
		}
	}
}

func TestDetranslate(t *testing.T) {
	bytes := []byte{
		0b00101010,
		0b00100101,
		0b00010110,
		0b00101010,
		0b00101010,
		0b00100000,
	}
	expect := []byte{
		0b10101010,
		0b01010101,
		0b10101010,
		0b10101010,
	}
	detranslated, _ := base64.Detranslate(bytes)

	for index := range expect {
		if detranslated[index] != expect[index] {
			t.Fatalf("[%d] Expected %06b, got %06b.", index, expect[index], detranslated[index])
		}
	}
}

func TestEncode(t *testing.T) {
	bytes := []byte("Hello, World!")
	expect := []byte("SGVsbG8sIFdvcmxkIQ==")

	encoded, _ := base64.Encode(bytes)

	for index := range expect {
		if encoded[index] != expect[index] {
			t.Fatalf("[%d] Expected %x, got %x.", index, expect[index], encoded[index])
		}
	}
}

func TestDecode(t *testing.T) {
	bytes := []byte("SGVsbG8sIFdvcmxkIQ==")
	expect := []byte("Hello, World!")

	decoded, _ := base64.Decode(bytes)

	for index := range expect {
		if decoded[index] != expect[index] {
			t.Fatalf("[%d] Expected %x, got %x.", index, expect[index], decoded[index])
		}
	}
}
