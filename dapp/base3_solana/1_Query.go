package main

import (
	"context"
	"fmt"
	"github.com/blocto/solana-go-sdk/client"
)

func main() {

	query()
}

func query() {
	//client0 := client.NewClient("https://api.devnet.solana.com")
	//BZgyU42TchAjZSajPp1xz9X3Y1lFj_y7

	//c := client.NewClient(rpc.MainnetRPCEndpoint)
	c := client.NewClient("https://solana-mainnet.g.alchemy.com/v2/BZgyU42TchAjZSajPp1xz9X3Y1lFj_y7")
	//c := client.NewClient("https://solana-devnet.g.alchemy.com/v2/BZgyU42TchAjZSajPp1xz9X3Y1lFj_y7")
	version, _ := c.GetVersion(context.Background())
	fmt.Println("Solana node version:", version.SolanaCore)

	//block, _ := c.GetBlock(context.Background(), 12345)
	//
	//fmt.Println("block.Blockhash: ", block.Blockhash)
	//fmt.Println("block.BlockHeight: ", block.BlockHeight)

	// 获取最新区块
	recentBlock, err := c.GetBlock(context.Background(), 0) // 0表示最新区块
	if err != nil {
		panic("查询失败: " + err.Error())
	}

	fmt.Printf("区块高度: %d\n", recentBlock.BlockHeight)
	fmt.Printf("交易数量: %d\n", len(recentBlock.Transactions))
}
