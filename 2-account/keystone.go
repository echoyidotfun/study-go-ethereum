package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func CreateKs() {
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(account.Address.Hex())
}

func ImportKs() {
	file := "./tmp/.986469000Z--90730431ce39d072a5707139951f03a83779558b"
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := os.ReadFile(file)
	if err!= nil {
		log.Fatal(err)
	}
	password := "secret"
	account, err := ks.Import(jsonBytes, password, password)
	if err!= nil {
		log.Fatal(err)
	}
	fmt.Println(account.Address.Hex())

	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// CreateKs()
	ImportKs()
}