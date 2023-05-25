package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

var keyCodes = [20]string{
	"53BE83A26FF44A74", "8E945A15C14F42C5", "3D7A9D2E8807435B", "8D6D52083887A315",
	"13B67DF295E34CA7", "A4AC79E5C6759C71", "937446ACF6034772", "BC9A4F39F8E42537",
	"87B602E5C6DC46EB", "8CFFA24A180E5857", "4F250ED69B984DD1", "A5AAA09CF51AB6F2",
	"4E9C1854CA914188", "B87EF5F56597598D", "8471C480A05B446B", "BF8AA2850A16A0E7",
	"163361E317894E04", "B7293B8FE96728DD", "7170DE5CA95D4524", "A890C9DE46B0FCE4",
}

func AesEncryptAnother(origData []byte, keyIndex int) ([]byte, error) {

	key := []byte(keyCodes[keyIndex])

	block, _ := aes.NewCipher(key)

	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecryptFile(cryptedByte []byte) ([]byte, error) {
	if len(cryptedByte) < 16 {
		return cryptedByte, nil
	}
	var headByte []byte
	var dataByte []byte
	for index, date := range cryptedByte {
		if index < 16 {
			headByte = append(headByte, date)
		} else {
			dataByte = append(dataByte, date)
		}
	}

	if headByte[0] != 0xFF || headByte[1] != 0xFF {
		return cryptedByte, nil
	}
	if headByte[2] != 1 {
		return cryptedByte, nil
	}
	if headByte[3] < 0 || headByte[3] >= 20 {
		return cryptedByte, nil
	}
	keyIndex := headByte[3]

	key := []byte(keyCodes[keyIndex])
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(dataByte))
	//var origData []byte
	blockMode.CryptBlocks(origData, dataByte)

	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
