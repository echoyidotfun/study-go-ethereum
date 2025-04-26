package main

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"study-go-ethereum/cfg"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main () {
	// 使用正则表达式匹配以太坊地址
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	fmt.Printf("is valid: %v\n", re.MatchString("0x323b5d4c32345ced77393b3530b1eed0f346429d")) // is valid: true
    fmt.Printf("is valid: %v\n", re.MatchString("0xZYXb5d4c32345ced77393b3530b1eed0f346429d")) // is valid: false

	config, err := cfg.LoadConfig("../config.json")
	if err!= nil {
		log.Fatalf("加载配置文件失败: %v", err)
	}

	client, err := ethclient.Dial(config.Ethereum.RPCURL)
	if err!= nil {
		log.Fatalf("连接以太坊节点失败: %v", err)
	}

	// odos arb
	address := common.HexToAddress("0xa669e7a0d4b3e4fa48af2de86bd4cd7126be4e13")
	bytecode, err := client.CodeAt(context.Background(), address, nil)

	if err!= nil {
		log.Fatalf("获取合约字节码失败: %v", err)
	}

	isContract := len(bytecode) > 0
	fmt.Printf("is contract: %v\n", isContract)

}