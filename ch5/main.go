package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

// eth转账
func main() {

	// 链接客户端
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/ZcQ5LqZBiwEi0ydLIqlm9")
	if err != nil {
		log.Fatal(err)
	}
	// 加载私钥
	privateKey, err := crypto.HexToECDSA("bd408c77966a926bf803539341de67d0bc9220709d666a14b1406837d055bb9f")
	if err != nil {
		log.Fatal(err)
	}
	// 从私钥得到公钥
	publicKey := privateKey.Public()
	// 类型断言，得到公钥的ECDSA格式
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	// ecdsa转换为地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//fromAddress := common.HexToAddress("0x71C063bf0235591235029d1C90bEFA69df2CC612")

	// 读取账户交易随机数
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	// 要转移的eth数量,wei
	value := big.NewInt(10000000000000000)
	// gas 上限
	gasLimit := uint64(21000)
	// gas 价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// to address
	toAddress := common.HexToAddress("0xabDABC95f5b0B14b0F53E7eC9020f847a17f6D2C")

	// 生成未签名交易
	//tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddress,
		Value:    value,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     nil,
	})
	// 使用发件人的私钥对事务进行签名
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx send: ", signedTx.Hash().Hex()) // 0xdc642e65675f43b7a9c01945f42c059ed98574c405d7ed2840c82d597eee7920
}
