package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

// 查询交易
func main() {

	url := "https://eth-sepolia.g.alchemy.com/v2/ZcQ5LqZBiwEi0ydLIqlm9"

	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	blockNumber := big.NewInt(5671744)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("block hash:", block.Hash().Hex()) // 0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 从block块中获取交易信息
	for _, tx := range block.Transactions() {
		fmt.Println("tx hash: ", tx.Hash().Hex()) // 交易哈希
		// 获取交易的 sender
		sender, err := types.Sender(types.NewEIP155Signer(chainID), tx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("sender:", sender.Hex())

		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("receipt:", receipt)
		break
	}

	// 根据块哈希获取交易
	blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5")
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("count:", count)
	tx, err := client.TransactionInBlock(context.Background(), blockHash, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx hash:", tx.Hash().Hex()) // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5

	// 通过交易hash获取交易
	tx, pending, err := client.TransactionByHash(context.Background(), common.HexToHash("0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("tx pending: ", pending)
	fmt.Println("tx hash:", tx.Hash().Hex())
}
