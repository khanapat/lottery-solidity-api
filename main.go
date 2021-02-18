package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"krungthai.com/khanapat/lottery/ether/feature"
)

func main() {
	endpoint := "https://kovan.infura.io/v3/9657539221eb40a79ce550650f0530a3"

	client, err := ethclient.Dial(endpoint)
	if err != nil {
		log.Fatal(err)
	}

	account := common.HexToAddress("0xc083EB69aa7215f4AFa7a22dcbfCC1a33999371C")
	checkEthBalance(*client, account)

	lotteryAddress := common.HexToAddress("0x41Aead3398676100B6C47172865279d70ce9A347")
	lotteryContract, err := feature.NewFeature(lotteryAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	checkPrizePool(*client, *lotteryContract, account)
	checkINLBalance(*client, *lotteryContract, account)
	checkVote(*client, *lotteryContract, account)
}

func checkEthBalance(client ethclient.Client, account common.Address) {
	balanceEth, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ETH Balance: %d Type: %T\n", balanceEth, balanceEth)
}

func checkPrizePool(client ethclient.Client, lotteryContract feature.Feature, account common.Address) {
	callOpt := bind.CallOpts{
		From: account,
	}
	prizePool, err := lotteryContract.CheckTotal(&callOpt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total Prize: %d Type: %T\n", prizePool, prizePool)
}

func checkINLBalance(client ethclient.Client, lotteryContract feature.Feature, account common.Address) {
	callOpt := bind.CallOpts{
		From: account,
	}
	balanceINL, err := lotteryContract.BalanceOf(&callOpt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("INL Balance: %d Type: %T\n", balanceINL, balanceINL)
}

func checkVote(client ethclient.Client, lotteryContract feature.Feature, account common.Address) {
	callOpt := bind.CallOpts{
		From: account,
	}
	address, vote, err := lotteryContract.CheckVoteByAddress(&callOpt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Address: %#x Vote: %v\n", address, vote)
}

// func voteTo(client ethclient.Client, voteContract vote.Vote, account common.Address, privateKey ecdsa.PrivateKey) {
// 	gas, err := client.SuggestGasPrice(context.Background())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	opt := bind.NewKeyedTransactor(&privateKey)
// 	opt.GasPrice = gas
// 	opt.GasLimit = 300000
// 	transaction, err := voteContract.VoteTo(opt, big.NewInt(1))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("https://kovan.etherscan.io/tx/%s\n", transaction.Hash().Hex())
// }
