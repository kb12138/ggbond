package main

import (
	"context"
	"fmt"
	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/rpc"
	"github.com/blocto/solana-go-sdk/system"
	"github.com/blocto/solana-go-sdk/types"
)

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	// 生成发送方密钥对（实际项目应从钱包加载）
	payer := types.NewAccount()

	// 获取最新区块哈希
	recentBlockhash, _ := c.GetLatestBlockhash(context.Background())

	// 创建转账指令（向随机地址转账0.01 SOL）
	instruction := system.NewTransferInstruction(
		1_000_000, // 0.01 SOL (lamport单位)
		payer.PublicKey,
		types.NewAccount().PublicKey,
	)

	// 构建交易
	tx, _ := types.NewTransaction(
		types.NewTransactionParam{
			Instructions:    []types.Instruction{instruction},
			Signers:         []types.Account{payer},
			RecentBlockHash: recentBlockhash.Blockhash,
			FeePayer:        payer.PublicKey,
		},
	)

	// 发送交易
	sig, _ := c.SendTransactionWithConfig(
		context.Background(),
		tx,
		rpc.SendTransactionConfig{
			Encoding: rpc.TransactionEncodingBase64,
		},
	)

	fmt.Printf("交易已提交，哈希: %s\n", sig)
}
