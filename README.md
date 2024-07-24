# Base64 Implementation in Go

## Overview

This repository contains a custom implementation of the Base64 encoding algorithm in Go. The purpose of this project is to provide an educational example of how Base64 encoding works internally.

**Warning:** This code is intended for educational purposes only. It is not optimized for production use. For any production applications, you should use the standard library's `encoding/base64` package, which is well-tested and optimized.

## Getting Started

### Prerequisites

To run this project, you will need to have Go installed on your machine. You can download and install Go from the [official website](https://golang.org/dl/).

### Installation

Clone the repository to your local machine:

```sh
git clone https://github.com/AndreyArthur/base64
cd base64
```

### Usage

You can run the provided examples to see how the custom Base64 implementation works.

### Example

Here's a quick example of how to use the custom Base64 encoder and decoder:

```go
package main

import (
    "base64/base64"
    "fmt"
)

func main() {
    input := "Hello, World!"
    encoded, err := base64.EncodeString(input)
    if err != nil {
        fmt.Println("Error encoding:", err)
        return
    }
    fmt.Println("Encoded:", encoded)

    decoded, err := base64.DecodeString(encoded)
    if err != nil {
        fmt.Println("Error decoding:", err)
        return
    }
    fmt.Println("Decoded:", decoded)
}
```

## Custom Implementation

The custom Base64 implementation includes the following functions:

- `TranslateBlock(block []byte) ([]byte, error)`: Translates a block of up to 3 bytes into a Base64 block.
- `DetranslateBlock(block []byte) ([]byte, error)`: Detranslates a Base64 block back into a byte block.
- `Translate(bytes []byte) ([]byte, error)`: Translates a byte slice into Base64 blocks.
- `Detranslate(bytes []byte) ([]byte, error)`: Detranslates a slice of Base64 blocks back into bytes.
- `Encode(bytes []byte) ([]byte, error)`: Encodes a byte slice into a Base64 encoded byte slice.
- `Decode(bytes []byte) ([]byte, error)`: Decodes a Base64 encoded byte slice back into bytes.
- `EncodeString(text string) (string, error)`: Encodes a string into a Base64 encoded string.
- `DecodeString(text string) (string, error)`: Decodes a Base64 encoded string back into a string.

## Testing

To run tests for the custom implementation, use the following command:

```sh
go test ./test/...
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
