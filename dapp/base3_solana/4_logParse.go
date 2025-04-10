package main

import (
	"encoding/json"
	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/program/token"
)

// 代币转账事件结构
type TokenTransfer struct {
	Source      common.PublicKey `json:"source"`
	Destination common.PublicKey `json:"destination"`
	Amount      uint64           `json:"amount"`
}

func parseTokenTransfer(logs []string) (*TokenTransfer, error) {
	// 查找转账指令日志
	for _, l := range logs {
		if token.IsTransferInstructionLog(l) {
			// 示例日志格式："Program log: Instruction: Transfer,
			// Amount=100000000"
			return &TokenTransfer{
				Source:      common.PublicKeyFromString(logs[2]), // 账户位置固定
				Destination: common.PublicKeyFromString(logs[3]),
				Amount:      extractAmountFromLog(l),
			}, nil
		}
	}
	return nil, errors.New("未找到转账事件")
}

// 从日志字符串提取金额
func extractAmountFromLog(log string) uint64 {
	var amount struct {
		Value uint64 `json:"Amount"`
	}
	json.Unmarshal([]byte(log[15:]), &amount) // 跳过"Program log: "前缀
	return amount.Value
}
