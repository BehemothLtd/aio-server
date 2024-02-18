package cryption

// Reference: https://www.kiteworks.com/risk-compliance-glossary/aes-256-encryption/
import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

// Generates a random IV for AES encryption.
func generateIV(blockSize int) ([]byte, error) {
	iv := make([]byte, blockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	return iv, nil
}

// Convert string into exact [32]byte but return []byte for Encryption use
// of Secure AES-256-CBC Encryption and Decryption
func StringTo32Bytes(input string) []byte {
	// Hash the input string using SHA-256. This will make sure output of [32]bytes
	sum := sha256.Sum256([]byte(input))
	sumStr := string(sum[:])
	return []byte(sumStr)
}

// Encrypts plaintext using AES-256-CBC with a given key and returns the result as hex.
// The first block of the ciphertext is the IV.
func Encrypt(plaintext string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Pad plaintext to the block size with PKCS#7 padding.
	padding := block.BlockSize() - len(plaintext)%block.BlockSize()
	paddedText := append([]byte(plaintext), bytes.Repeat([]byte{byte(padding)}, padding)...)

	// Generate a random IV.
	iv, err := generateIV(block.BlockSize())
	if err != nil {
		return "", err
	}

	// Encrypt the plaintext.
	ciphertext := make([]byte, len(paddedText))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, paddedText)

	// Prefix the IV to the ciphertext.
	return hex.EncodeToString(append(iv, ciphertext...)), nil
}

// Decrypts a hex-encoded ciphertext using AES-256-CBC and the given key.
func Decrypt(encryptedHex string, key []byte) (string, error) {
	ciphertext, err := hex.DecodeString(encryptedHex)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < block.BlockSize() {
		return "", fmt.Errorf("ciphertext too short")
	}

	// The IV is the first block of the ciphertext.
	iv := ciphertext[:block.BlockSize()]
	ciphertext = ciphertext[block.BlockSize():]

	// Decrypt the ciphertext.
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	// Remove padding.
	padding := ciphertext[len(ciphertext)-1]
	if int(padding) > block.BlockSize() || int(padding) == 0 {
		return "", fmt.Errorf("invalid padding")
	}
	unpadded := ciphertext[:len(ciphertext)-int(padding)]

	return string(unpadded), nil
}
