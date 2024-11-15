package utils

import (
	gonanoid "github.com/matoous/go-nanoid"
)

func GenerateNanoIdWithLength(length int) string {
	nanoId, _ := gonanoid.Generate("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", length)
	return nanoId
}
