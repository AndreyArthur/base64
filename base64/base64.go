package base64

import "errors"

var encoding = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '+', '/'}
var decoding = [256]byte{'A': 0, 'B': 1, 'C': 2, 'D': 3, 'E': 4, 'F': 5, 'G': 6, 'H': 7, 'I': 8, 'J': 9, 'K': 10, 'L': 11, 'M': 12, 'N': 13, 'O': 14, 'P': 15, 'Q': 16, 'R': 17, 'S': 18, 'T': 19, 'U': 20, 'V': 21, 'W': 22, 'X': 23, 'Y': 24, 'Z': 25, 'a': 26, 'b': 27, 'c': 28, 'd': 29, 'e': 30, 'f': 31, 'g': 32, 'h': 33, 'i': 34, 'j': 35, 'k': 36, 'l': 37, 'm': 38, 'n': 39, 'o': 40, 'p': 41, 'q': 42, 'r': 43, 's': 44, 't': 45, 'u': 46, 'v': 47, 'w': 48, 'x': 49, 'y': 50, 'z': 51, '0': 52, '1': 53, '2': 54, '3': 55, '4': 56, '5': 57, '6': 58, '7': 59, '8': 60, '9': 61, '+': 62, '/': 63}

func TranslateBlock(block []byte) ([]byte, error) {
	if len(block) < 1 || len(block) > 3 {
		return nil, errors.New("Unsupported length!")
	}

	bytes := [3]byte{}
	copy(bytes[:], block)

	translated := [4]byte{}

	translated[0] = bytes[0] >> 2
	translated[1] = ((bytes[0] << 4) | (bytes[1] >> 4)) & 0x3f
	translated[2] = ((bytes[1] << 2) | (bytes[2] >> 6)) & 0x3f
	translated[3] = bytes[2] & 0x3f

	var end int
	switch len(block) {
	case 1:
		end = 2
	case 2:
		end = 3
	case 3:
		end = 4
	}

	return translated[:end], nil
}

func DetranslateBlock(block []byte) ([]byte, error) {
	if len(block) < 2 || len(block) > 4 {
		return nil, errors.New("Unsupported length!")
	}

	bytes := [4]byte{}
	copy(bytes[:], block)

	detranslated := [3]byte{}

	detranslated[0] = bytes[0]<<2 | bytes[1]>>4
	detranslated[1] = bytes[1]<<4 | bytes[2]>>2
	detranslated[2] = bytes[2]<<6 | bytes[3]

	var end int
	switch len(block) {
	case 2:
		end = 1
	case 3:
		end = 2
	case 4:
		end = 3
	}

	return detranslated[:end], nil
}

func Translate(bytes []byte) ([]byte, error) {
	translated := []byte{}

	for i := 0; i < len(bytes); i += 3 {
		var block []byte
		if i+3 > len(bytes) {
			block = bytes[i:]
		} else {
			block = bytes[i : i+3]
		}

		translatedBlock, err := TranslateBlock(block)
		if err != nil {
			return nil, err
		}

		translated = append(translated, translatedBlock...)
	}

	return translated, nil
}

func Detranslate(bytes []byte) ([]byte, error) {
	detranslated := []byte{}

	for i := 0; i < len(bytes); i += 4 {
		var block []byte
		if i+4 > len(bytes) {
			block = bytes[i:]
		} else {
			block = bytes[i : i+4]
		}

		detranslatedBlock, err := DetranslateBlock(block)
		if err != nil {
			return nil, err
		}

		detranslated = append(detranslated, detranslatedBlock...)
	}

	return detranslated, nil
}

func Encode(bytes []byte) ([]byte, error) {
	translated, err := Translate(bytes)
	if err != nil {
		return nil, err
	}

	encoded := []byte{}
	for i := 0; i < len(translated); i++ {
		encoded = append(encoded, encoding[translated[i]])
	}

	padding := 4 - len(encoded)%4
	for i := 0; i < padding; i++ {
		encoded = append(encoded, '=')
	}

	return encoded, nil
}

func Decode(bytes []byte) ([]byte, error) {
	padded := 0
	for i := len(bytes) - 1; i >= 0; i-- {
		if bytes[i] == '=' {
			padded++
		} else {
			break
		}
	}

	decoded := []byte{}
	for i := 0; i < len(bytes)-padded; i++ {
		decoded = append(decoded, decoding[bytes[i]])
	}

	detranslated, err := Detranslate(decoded)
	if err != nil {
		return nil, err
	}

	return detranslated, nil
}

func EncodeString(text string) (string, error) {
	encoded, err := Encode([]byte(text))
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func DecodeString(text string) (string, error) {
	decoded, err := Decode([]byte(text))
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}
