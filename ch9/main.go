package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

// 订阅区块
func main() {
	client, err := ethclient.Dial("wss://eth-sepolia.g.alchemy.com/v2/ZcQ5LqZBiwEi0ydLIqlm9")
	if err != nil {
		log.Fatal(err)
	}

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex()) // 0xab56fe9f7381b4dc33afbeeaaedf5c9484cf641edbdc6318d402151e7c6a2afd
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(block.Hash().Hex())        // 0xab56fe9f7381b4dc33afbeeaaedf5c9484cf641edbdc6318d402151e7c6a2afd
			fmt.Println(block.Number().Uint64())   // 9473464
			fmt.Println(block.Time())              //
			fmt.Println(block.Nonce())             // 0
			fmt.Println(len(block.Transactions())) // 158
		}
	}
}
