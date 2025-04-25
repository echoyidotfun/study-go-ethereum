package main

import (
	"fmt"
	"log"

	"study-go-ethereum/cfg"

	"github.com/ethereum/go-ethereum/ethclient"
)



func main() {
	config, err := cfg.LoadConfig("../config.json")
	if err != nil {
		log.Fatalf("加载配置文件失败: %v", err)
	}

	client, err := ethclient.Dial(config.Ethereum.RPCURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("已成功连接")
	_ = client
	
}