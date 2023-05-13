package utils

import (
	"crypto/rand"
	"fmt"
	"unsafe"

	"github.com/google/uuid"
	"github.com/spf13/viper"
	"golang.org/x/crypto/argon2"
)

var alphabet = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func CustomUUID() uuid.UUID { // this function is specify all course  uuid start with ffc50000
	uuid := uuid.New()
	uuid[0] = 0xff
	uuid[1] = 0xc5
	uuid[2] = 0x00
	uuid[4] = 0x00

	return uuid
}

func RandStr() string {
	size := 50 // specifing the length of the string
	b := make([]byte, size)
	rand.Read(b)
	for i := 0; i < size; i++ {
		b[i] = alphabet[b[i]%byte(len(alphabet))]
	}
	return *(*string)(unsafe.Pointer(&b))
}

var salt = []byte(viper.GetString("SALTVALUE"))

func Hashes(password string) string {
	time := uint32(2)
	memory := uint32(64 * 1024)
	threads := uint8(4)
	keyLength := uint32(32)
	hash := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLength)
	return fmt.Sprintf("%x", hash)
}
