package global

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"github.com/warthecatalyst/douyin/config"
	"strconv"
)

func aesCtrCrypt(plainText []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	iv := bytes.Repeat([]byte("1"), block.BlockSize())
	stream := cipher.NewCTR(block, iv)
	dst := make([]byte, len(plainText))
	stream.XORKeyStream(dst, plainText)

	return dst, nil
}

func CreateToken(userId int64, username, password string) (string, error) {
	// TODO add salt
	token, err := aesCtrCrypt([]byte(strconv.FormatInt(userId, 10)), []byte(config.TokenEncryptKey))
	if err != nil {
		DyLogger.Errorf("aesCtrCrypt error: %s", err)
		return "", err
	}
	return string(token), nil
}

func GetUserIdFromToken(token string) (int64, error) {
	code, err := aesCtrCrypt([]byte(token), []byte(config.TokenEncryptKey))
	if err != nil {
		DyLogger.Errorf("aesCtrCrypt error: %s", err)
		return -1, err
	}

	userId, err := strconv.Atoi(string(code))
	if err != nil {
		DyLogger.Errorf("strconv.Atoi error: %s", err)
		return -1, err
	}

	return int64(userId), nil
}
