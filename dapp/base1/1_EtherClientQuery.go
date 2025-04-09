package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"go_test/dapp/base1/config"
	"log"
	"math/big"
)

func main1() {
	var client, err = ethclient.Dial(config.Test_Url)
	//var client, err = ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(header.Number.String())

	blockNumber := big.NewInt(8083522)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Block Number (Uint64):", block.Number().Uint64())                 // 打印区块的编号，为无符号 64 位整数
	fmt.Println("Block Timestamp:", block.Time())                                  // 打印区块的时间戳，代表该区块生成的时间
	fmt.Println("Block Difficulty (Uint64):", block.Difficulty().Uint64())         // 打印区块的难度值，为无符号 64 位整数
	fmt.Println("Block Hash (Hex):", block.Hash().Hex())                           // 打印区块的哈希值，以十六进制字符串形式表示
	fmt.Println("Number of Transactions in the Block:", len(block.Transactions())) // 打印该区块中包含的交易数量

	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count) // 70

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for _, tx := range block.Transactions() {
		//if tx.Hash().Hex() != "0x5dc8971c6a954eab3c4b9fa854f365de36f04a9333c695620642d8da4d70621a" {
		//	continue
		//}
		fmt.Println("Transaction Hash (Hex):", tx.Hash().Hex())                // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca5
		fmt.Println("Transaction Value (String):", tx.Value().String())        // 100000000000000000
		fmt.Println("Transaction Gas:", tx.Gas())                              // 21000
		fmt.Println("Transaction Gas Price (Uint64):", tx.GasPrice().Uint64()) // 100000000000
		fmt.Println("Transaction Nonce:", tx.Nonce())                          // 245132
		fmt.Println("Transaction Data:", tx.Data())                            // []
		TryParseData(tx.Data())
		fmt.Println("Transaction To Address (Hex):", tx.To().Hex()) // 0x8F9aFd209339088Ced7Bc0f57Fe08566ADda3587

		if err != nil {
			log.Fatal(err)
		}
		if sender, err := types.Sender(types.NewEIP155Signer(chainID), tx); err == nil {
			fmt.Println("sender： ", sender.Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
		}
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("receipt.Status: ", receipt.Status) // 1
		fmt.Println("receipt.Logs", receipt.Logs)       // ...
	}

}

// TryParseData 尝试以不同方式解析字节切片并打印结果
func TryParseData(data []byte) {
	if len(data) == 0 {
		return
	}
	// 1. 尝试作为字符串解析
	strData := string(data)
	fmt.Printf("尝试作为字符串解析: %s\n", strData)

	// 2. 尝试以十六进制形式解析
	hexData := hex.EncodeToString(data)
	fmt.Printf("尝试以十六进制形式解析: %s\n", hexData)

	// 3. 尝试作为 JSON 数据解析
	var jsonResult interface{}
	err := json.Unmarshal(data, &jsonResult)
	if err == nil {
		fmt.Printf("尝试作为 JSON 数据解析: %+v\n", jsonResult)
	} else {
		fmt.Printf("尝试作为 JSON 数据解析失败: %v\n", err)
	}

	// 4. 打印每个字节的十进制值
	fmt.Print("每个字节的十进制值: ")
	for _, b := range data {
		fmt.Printf("%d ", b)
	}
	fmt.Println()
}
