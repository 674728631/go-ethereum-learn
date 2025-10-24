package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"math/big"
)

// 查询收据
func main() {

	url := "https://eth-sepolia.g.alchemy.com/v2/ZcQ5LqZBiwEi0ydLIqlm9"

	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	// 通过区块哈希获取手机 0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5
	receiptByHash, err := client.BlockReceipts(context.Background(),
		rpc.BlockNumberOrHashWithHash(common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5"), false))
	if err != nil {
		log.Fatal(err)
	}
	blockNumber := big.NewInt(5671744)
	receiptByNum, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Uint64())))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(receiptByHash[0].TxHash.Hex() == receiptByNum[0].TxHash.Hex())
}
