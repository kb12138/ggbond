package main

import (
	"context"
	"fmt"
	"github.com/blocto/solana-go-sdk/rpc"
	"github.com/gagliardetto/ws"
)

func main() {
	// 创建WebSocket客户端（连接到DevNet） // 导包有问题
	wsClient, err := ws.Connect(context.Background(), rpc.DevnetWS)
	if err != nil {
		panic("连接失败: " + err.Error())
	}

	// 订阅程序日志（示例：SPL Token程序）
	sub, err := wsClient.LogsSubscribeMatching(
		"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
		rpc.CommitmentConfirmed,
	)
	if err != nil {
		panic("订阅失败: " + err.Error())
	}
	defer sub.Unsubscribe()

	// 实时处理事件
	for {
		select {
		case log := <-sub.Channel():
			fmt.Printf("[程序日志] 账户: %s\n", log.Value.Pubkey)
			fmt.Printf("      日志内容: %+v\n", log.Value.Logs)
		case err := <-sub.Err():
			fmt.Printf("监听错误: %v\n", err)
			return
		}
	}
}
