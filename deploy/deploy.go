package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// 读取合约二进制文件和ABI文件
	contractBin, err := ioutil.ReadFile("build/MyContract.bin")
	if err != nil {
		log.Fatalf("Failed to read contract binary: %v", err)
	}

	contractABI, err := ioutil.ReadFile("build/MyContract.abi")
	if err != nil {
		log.Fatalf("Failed to read contract ABI: %v", err)
	}

	// 连接到以太坊节点
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/YOUR-INFURA-KEY")
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum network: %v", err)
	}

	// 创建新账户
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalf("Failed to generate private key: %v", err)
	}

	// 构建部署的auth对象
	auth := bind.NewKeyedTransactor(privateKey)

	// 部署合约
	address, tx, _, err := bind.DeployContract(auth, bind.ContractBackend(client), contractABI, contractBin, nil)
	if err != nil {
		log.Fatalf("Failed to deploy contract: %v", err)
	}

	fmt.Printf("Contract deployed to: %s\n", address.Hex())
	fmt.Printf("Transaction hash: %s\n", tx.Hash().Hex())
}
