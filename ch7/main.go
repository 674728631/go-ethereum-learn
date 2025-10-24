package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

// 查询账户余额
func main() {

	// 链接客户端
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/ZcQ5LqZBiwEi0ydLIqlm9")
	if err != nil {
		log.Fatal(err)
	}

	// 获取余额
	address := common.HexToAddress("0xabDABC95f5b0B14b0F53E7eC9020f847a17f6D2C")
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("balance: ", balance)

	// 处理余额
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(1e18))
	fmt.Println("ethValue: ", ethValue)
}
