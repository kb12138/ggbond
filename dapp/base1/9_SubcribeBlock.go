package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"go_test/dapp/base1/config"
	"log"
)

func main9() {
	client, err := ethclient.Dial(config.Infura_Wss)
	if err != nil {
		log.Fatal(err)
	}

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)

	//client.SubscribeFilterLogs()
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("(block.Hash().Hex(): ", block.Hash().Hex())               // SHA-3 Keccak-256哈希（66字符），区块唯一指纹 | 示例值截断：0xbc10defa...（用于链上定位/防篡改验证，任何数据变更都会改变哈希）
			fmt.Println("(block.Number().Uint64(): ", block.Number().Uint64())     // 区块高度（从0开始计数），当前为第3477414个区块 | 主网高度>1.8亿（2024年），大数值建议用block.Number().String()防溢出
			fmt.Println("(block.Time(): ", block.Time())                           // UNIX时间戳（秒）→ 北京时间：2018-06-21 11:39:07（time.Unix(1529525947, 0).Format()）| 矿工打包时的本地时间，允许±N秒误差
			fmt.Println("(block.Nonce(): ", block.Nonce())                         // PoW随机数（8字节），仅ETH合并前有效 | 当前值：130524141876765836（ETH合并后返回0，旧链如ETHW仍用于工作量证明）
			fmt.Println("(len(block.Transactions()): ", len(block.Transactions())) // 区块内交易数：7笔 | 受区块Gas限制（现约3000万Gas），每笔交易平均21000Gas时最多约1428笔
		}
	}
}
