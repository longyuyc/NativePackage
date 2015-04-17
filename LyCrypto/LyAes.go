/*
高级加密标准（英语：Advanced Encryption Standard，缩写：AES），
在密码学中又称Rijndael加密法，是美国联邦政府采用的一种区块加密标准。
这个标准用来替代原先的DES，已经被多方分析且广为全世界所使用。
经过五年的甄选流程，高级加密标准由美国国家标准与技术研究院（NIST）
于2001年11月26日发布于FIPS PUB 197，并在2002年5月26日成为有效的标准。
2006年，高级加密标准已然成为对称密钥加密中最流行的算法之一。
*/
package LyCrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"sync"
)

//key
const aesTable = "abcdefghijklmnopkrstuvwsyz012345"

var (
	block cipher.Block
	mutex sync.Mutex
)

func TestAes() {
	fmt.Println("TestAes")
}

func init() {
	fmt.Println("lyAes init")
	mutex.Lock()
	defer mutex.Unlock()
	if block != nil {
		return
	}

	cblock, err := aes.NewCipher([]byte(aesTable))
	if err != nil {
		panic("aes.NewCipher:" + err.Error())
	}
	block = cblock
}
