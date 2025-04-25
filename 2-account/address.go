package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"study-go-ethereum/cfg"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	config, err := cfg.LoadConfig("../config.json")
	if err != nil {
		log.Fatalf("加载配置文件失败: %v", err)
	}
	client, err := ethclient.Dial(config.Ethereum.RPCURL)
	if err!= nil {
		log.Fatalf("连接以太坊节点失败: %v", err)
	}

	address := common.HexToAddress("0x9A01CE1eCbb36E4f3892c94F6243a4E27E5Ce3DE")
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("获取余额失败: %v", err)
	}

	fmt.Println("账户余额: ", balance)

	// 获取当前区块高度
	blockNumber, _ := client.BlockNumber(context.Background())
	blockNumberBig := new(big.Int).SetUint64(blockNumber)
	balanceAt, err := client.BalanceAt(context.Background(), address, blockNumberBig)
	if err != nil {
		log.Fatalf("获取指定区块的余额失败: %v", err)
	}
	fmt.Printf("区块高度: %d 的余额: %d\n", blockNumber, balanceAt)

	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())

	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println("scaled账户余额: ", ethValue)

	// 什么是pending余额？
	// pending余额是指当前账户的余额，包括当前未被确认的交易。
	pendingBalance, err := client.PendingBalanceAt(context.Background(), address)
	if err!= nil {
		log.Fatalf("获取pending余额失败: %v", err)
	}	
	fmt.Println("pending余额: ", pendingBalance)

}