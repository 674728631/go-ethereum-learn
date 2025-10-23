package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"log"
)

// 创建新钱包
func main() {

	// 生成随机私钥
	//privateKey, err := crypto.GenerateKey()
	//if err != nil {
	//	log.Fatal(err)
	//}

	// 私钥的 Hex 字符串，HexToECDSA 方法恢复私钥
	privateKey, err := crypto.HexToECDSA("bd408c77966a926bf803539341de67d0bc9220709d666a14b1406837d055bb9f")
	if err != nil {
		log.Fatal(err)
	}
	// 转换为字节
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("privateKeyBytes hex: ", hexutil.Encode(privateKeyBytes)[2:]) // 去掉'0x'

	// 公钥是从私钥派生
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("from pubKey:", hexutil.Encode(publicKeyBytes)[4:]) // 去掉'0x04'

	// 公共地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("address: ", address) // 0x19BaF43f1a88E9E8fB512E2D275ccA2BD5aE9DFC
	// 公共地址其实就是公钥的 Keccak-256 哈希
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println("full:", hexutil.Encode(hash.Sum(nil)[:]))    // 原长32位，截去12位，保留后20位
	fmt.Println("hash: ", hexutil.Encode(hash.Sum(nil)[12:])) // 原长32位，截去12位，保留后20位
}
