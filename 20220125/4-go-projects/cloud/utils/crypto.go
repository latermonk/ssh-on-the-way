package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

// 计算字符串sha256 hex值
func Sha256Hex(text string) string {
	hasher := sha256.New()
	hasher.Write([]byte(text))
	return strings.ToLower(hex.EncodeToString(hasher.Sum(nil)))
}

// 计算字符串hmac-sha256值
func HS256(text, key string) string {
	hasher := hmac.New(sha256.New, []byte(key))
	hasher.Write([]byte(text))
	return string(hasher.Sum(nil))
}

// 计算字符串hmac-sha256 hex值
func HS256Hex(text, key string) string {
	return strings.ToLower(hex.EncodeToString([]byte(HS256(text, key))))
}
