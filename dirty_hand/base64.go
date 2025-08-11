package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

// Base64Encode 将字节数据编码为 Base64 字符串
func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Base64Decode 将 Base64 字符串解码为原始字节数据
func Base64Decode(s string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, fmt.Errorf("base64 decode failed: %w", err)
	}
	return decoded, nil
}

func main() {
	// 编码示例
	originalData := []byte("2355199564")
	encoded := Base64Encode(originalData)
	fmt.Printf("Original Data:   %s\n", originalData)
	fmt.Printf("Base64 Encoded: %s\n", encoded)

	// 解码示例
	decoded, err := Base64Decode(encoded)
	if err != nil {
		log.Fatalf("Decoding failed: %v", err)
	}
	fmt.Printf("Decoded Data:    %s\n", decoded)

	// 处理填充符和非标准字符
	encodedWithPadding := "SGVsbG8sIFdvcmxkIQ=="
	decodedWithPadding, err := Base64Decode(encodedWithPadding)
	if err != nil {
		log.Fatalf("Decoding failed: %v", err)
	}
	fmt.Printf("Decoded Padded: %s\n", decodedWithPadding)

	rst, err := Base64Decode("Mjg1MjE5OTcwOA==")
	fmt.Printf("%s\n", rst)
}
