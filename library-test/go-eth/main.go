package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"
	"sync"
	"time"

	"go-learning/library-test/go-eth/token"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	MinConfirmation = 5
	ToAddr          = "0xc40c7f7c427642cc7aabc8fe523e0905b040e6e2"
	GATContract     = "0xFD01302F6B6B5be5A71399Eb51c8282A98ed4551"
)

func main() {
	// client, err := ethclient.Dial("https://ropsten.infura.io")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// b, err := client.BlockByNumber(context.Background(), big.NewInt(4337125))
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println(b, len(b.Transactions()))

	// BlockNumber()
	// var wg sync.WaitGroup
	// wg.Add(1)
	// ListenForReceive()
	// GetBalance()
	// GetTxInfo()
	TestSendEth()
	// IsSubscribedTxPending()
	// DoubleSubscription()
	// go ListenToContract()

	// txHash := SendTokenTransaction()
	// cur := TransactionInBlock(txHash)
	// SubscribeNewBlock(txHash)
	// WaitConfirmation(cur)

	// wg.Wait()

	// SubEventLogs()
	// ReadEventLogs()
}

type LogTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
}

// LogApproval ..
type LogApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
}

func TestSendEth() {

	client, err := ethclient.Dial("https://ropsten.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("4B1ACD7E32ADF56BC9F9416D06672850A535BFC6821A637ACD849D03A1A014DD")
	if err != nil {
		log.Fatal(err)
	}

	fromAddress := common.HexToAddress("0x7A8a7870bd398240CE3523Ab7E76153ba34B7116")
	toAddress := common.HexToAddress("0xC40C7f7C427642CC7aABC8fe523e0905b040E6E2")
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000)                // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("gas price: ", gasPrice)
	fmt.Println("gas limit: ", gasLimit)
	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}

func GetTxInfo() {
	ctx := context.Background()
	client, err := ethclient.Dial("https://ropsten.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0xC40C7f7C427642CC7aABC8fe523e0905b040E6E2")

	amount := "1000000000000000000"
	a, _ := new(big.Int).SetString(amount, 10)
	data := ConstructTokenTxData(toAddress, a)

	gasLimit, err := client.EstimateGas(ctx, ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(gasPrice, gasLimit)
}

// ConstructTokenTxData constructs data for transfer token
func ConstructTokenTxData(toAddress common.Address, amount *big.Int) []byte {
	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4] // 0xa9059cbb

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	return data
}

func GetBalance() {
	client, err := ethclient.Dial("https://ropsten.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	account := common.HexToAddress("0xC40C7f7C427642CC7aABC8fe523e0905b040E6E2")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance)
}

func ReadEventLogs() {
	client, err := ethclient.Dial("https://ropsten.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	// 0x Protocol (ZRX) token address
	contractAddress := common.HexToAddress(GATContract)
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(4083412),
		ToBlock:   big.NewInt(4083489),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("logs", len(logs))

	// abi.Event{
	// 	Name: "Transfer",
	// }

	contractAbi, err := abi.JSON(strings.NewReader(string(token.TokenABI)))
	if err != nil {
		log.Fatal(err)
	}

	logTransferSig := []byte("Transfer(address,address,uint256)")
	LogApprovalSig := []byte("Approval(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	logApprovalSigHash := crypto.Keccak256Hash(LogApprovalSig)

	for _, vLog := range logs {
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vLog.Index)

		// block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(vLog.BlockNumber)))
		// if err != nil {
		// 	log.Println(err)
		// }

		// for _, tx := range block.Transactions() {
		// 	if tx.To() != nil && tx.To().Hex() == GATContract {
		// 		if tx.Value() != nil {
		// 			fmt.Printf("Amount: %s\n", tx.Value().String())
		// 		} else {
		// 			fmt.Printf("Amount: %v\n", tx.Value())
		// 		}
		// 	}
		// 	continue
		// }

		switch vLog.Topics[0].Hex() {
		case logTransferSigHash.Hex():
			fmt.Printf("Log Name: Transfer\n")

			var transferEvent LogTransfer

			// a := abi.ABI{Events: map[string]abi.Event{"Transfer": contractAbi.Events["Transfer"]}}

			err := contractAbi.Unpack(&transferEvent, "Transfer", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Data: %v\n", vLog.Data)

			transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
			transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Printf("From: %s\n", transferEvent.From.Hex())
			fmt.Printf("To: %s\n", transferEvent.To.Hex())
			fmt.Printf("Tokens: %s\n", transferEvent.Value.String())

		case logApprovalSigHash.Hex():
			fmt.Printf("Log Name: Approval\n")

			var approvalEvent LogApproval

			err := contractAbi.Unpack(&approvalEvent, "Approval", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			approvalEvent.TokenOwner = common.HexToAddress(vLog.Topics[1].Hex())
			approvalEvent.Spender = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Printf("Token Owner: %s\n", approvalEvent.TokenOwner.Hex())
			fmt.Printf("Spender: %s\n", approvalEvent.Spender.Hex())
			fmt.Printf("Tokens: %s\n", approvalEvent.Tokens.String())
		}

		fmt.Printf("\n\n")
	}
}

func SubEventLogs() {
	client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(GATContract)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(token.TokenABI)))
	if err != nil {
		log.Fatal(err)
	}

	logTransferSig := []byte("Transfer(address,address,uint256)")

	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			switch vLog.Topics[0].Hex() {
			case logTransferSigHash.Hex():
				fmt.Printf("Log Name: Transfer\n")

				var transferEvent LogTransfer

				err := contractAbi.Unpack(&transferEvent, "Transfer", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}

				transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
				transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

				fmt.Printf("From: %s\n", transferEvent.From.Hex())
				fmt.Printf("To: %s\n", transferEvent.To.Hex())
				fmt.Printf("Tokens: %s\n", transferEvent.Value.String())
			}
		}
	}
}

func DoubleSubscription() {
	client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		headers := make(chan *types.Header)
		sub, err := client.SubscribeNewHead(context.Background(), headers)
		if err != nil {
			log.Fatal(err)
		}

		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case header := <-headers:
				log.Println("1 #", header.Number.String())
			}
		}
	}()

	go func() {
		headers := make(chan *types.Header)
		sub, err := client.SubscribeNewHead(context.Background(), headers)
		if err != nil {
			log.Fatal(err)
		}

		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case header := <-headers:
				log.Println("2 #", header.Number.String())
			}
		}
	}()

	wg.Wait()
}

func IsSubscribedTxPending() {
	client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
	if err != nil {
		log.Fatal(err)
	}

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			// fmt.Printf("\n#%s\n", header.Number.String())

			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			for _, tx := range block.Transactions() {
				_, ispending, err := client.TransactionByHash(context.Background(), tx.Hash())
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("is pending:", ispending)
			}
		}
	}

}

func ListenForReceive() {
	var isFound bool

	client, err := ethclient.Dial("https://ropsten.infura.io")
	// client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
	if err != nil {
		log.Fatal(err)
	}

	for !isFound {
		header, err := client.HeaderByNumber(context.Background(), nil)
		if err != nil {
			log.Fatal("header error: ", err)
		}

		fmt.Printf("current header #%s\n", header.Number.String())

		block, err := client.BlockByNumber(context.Background(), header.Number)
		if err != nil {
			log.Println("block error: ", err)
			continue
		}

		fmt.Printf("latest block #%s\n", block.Number().String())

		for _, tx := range block.Transactions() {
			fmt.Println("tx hash:", tx.Hash().Hex())
			fmt.Println("tx to:", tx.To())
			if tx.To() != nil {
				fmt.Println("tx to:", tx.To().Hex())
			}

			chainID, err := client.NetworkID(context.Background())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("chainID: ", chainID)

			if tx.To() != nil && GATContract == tx.To().Hex() {
				log.Printf("receiving tx in block #%s\n", header.Number.String())

				msg, err := tx.AsMessage(types.NewEIP155Signer(chainID))
				if err != nil {
					fmt.Println(msg.To().Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
				}

				if msg.To().Hex() == ToAddr {
					isFound = true
				}

				_, ispending, err := client.TransactionByHash(context.Background(), tx.Hash())
				if err != nil {
					log.Fatal(err)
				}

				log.Println("is pending", ispending)

				if !ispending && IsTransactionSuccess(client, tx.Hash()) {
					log.Println("transaction success")
				}
				log.Fatal("Transaction error")
			}
		}
	}

}
func WaitConfirmation(current *big.Int) {
	var isConfirmed bool

	client, err := ethclient.Dial("https://ropsten.infura.io")
	// client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
	if err != nil {
		log.Fatal(err)
	}

	for !isConfirmed {
		time.Sleep(1 * time.Second)

		header, err := client.HeaderByNumber(context.Background(), nil)
		if err != nil {
			log.Fatal("header error: ", err)
		}
		log.Printf("current header #%s\n", header.Number.String())
		sub := big.NewInt(0)
		sub.Sub(header.Number, current)
		log.Println(sub)
		if sub.Cmp(big.NewInt(MinConfirmation-1)) >= 0 {
			isConfirmed = true
			log.Println("confirmed")
		}
	}
}

func TransactionInBlock(txHash string) *big.Int {
	log.Println()

	var isFound bool

	client, err := ethclient.Dial("https://ropsten.infura.io")
	// client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
	if err != nil {
		log.Fatal(err)
	}

	for !isFound {
		time.Sleep(1 * time.Second)

		header, err := client.HeaderByNumber(context.Background(), nil)
		if err != nil {
			log.Fatal("header error: ", err)
		}
		log.Printf("latest header #%s\n", header.Number.String())

		block, err := client.BlockByNumber(context.Background(), header.Number)
		if err != nil {
			log.Println("block error: ", err)
			continue
		}
		log.Printf("latest block #%s\n", block.Number().String())

		for _, tx := range block.Transactions() {
			if txHash == tx.Hash().Hex() {
				log.Printf("tx in block #%s\n", header.Number.String())
				isFound = true
				if IsTransactionSuccess(client, tx.Hash()) {
					return header.Number
				}

				log.Fatal("Transaction error")
			}
		}

	}

	log.Println("find")
	return big.NewInt(0)
}

func IsTransactionSuccess(client *ethclient.Client, txHash common.Hash) bool {
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatal("receipt error", err)
	}

	if receipt.Status == 1 {
		return true
	}
	return false

}

func BlockNumber() {
	client, err := ethclient.Dial("https://ropsten.infura.io")
	// client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
	if err != nil {
		log.Fatal(err)
	}

	// return the latest blk number
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	block, err := client.BlockByHash(context.Background(), header.Hash())
	if err != nil {
		log.Println(err)
	}

	log.Println(block, len(block.Transactions()))
}

// func GetTransaction() {
// 	client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// return the latest blk number
// 	client.TransactionByHash(context.Background(), txHash)

// 	client.TransactionReceipt()
// }

// func SubscribeContract() {
// 	client, err := ethclient.Dial("https://ropsten.infura.io")
// 	// client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	to := common.HexToAddress(GATContract)
// 	msg := ethereum.CallMsg{
// 		To: &to,
// 	}
// 	b, err := client.SubscribeFilterLogs(context.Background(), )
// 	if err != nil {

// 	}
// }

func SubscribeNewBlock() {
	log.Println("subscribe to new block")
	client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
	if err != nil {
		log.Fatal(err)
	}

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	var isFound bool
	for !isFound {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			log.Println()
			fmt.Println("#", header.Number.String()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f

			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			log.Println(block)
			chainID, err := client.NetworkID(context.Background())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("chainID: ", chainID)

			// fmt.Println(block.Hash().Hex())        // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			// fmt.Println(block.Number().Uint64())   // 3477413
			// fmt.Println(block.Time().Uint64())     // 1529525947
			// fmt.Println(block.Nonce())             // 130524141876765836
			// fmt.Println(len(block.Transactions())) // 7

			// for _, tx := range block.Transactions() {
			// 	fmt.Println("tx hash:", tx.Hash().Hex())
			// 	if txHash == tx.Hash().Hex() {
			// 		log.Printf("tx in block #%s\n", header.Number.String())
			// 		isFound = true
			// 		log.Println("is to equal to token contract: ", tx.To().Hex() == GATContract)
			// 		if IsTransactionSuccess(client, tx.Hash()) {
			// 			log.Println("Transaction success")
			// 			// return header.Number
			// 		}

			// 		log.Fatal("Transaction error")
			// 	}

			// if tx.To() != nil && tx.To().Hex() == GATContract {

			// 	log.Printf("receiving tx in block #%s\n", header.Number.String())

			// 	msg, err := tx.AsMessage(types.NewEIP155Signer(chainID))
			// 	if err != nil {
			// 		fmt.Println(msg.To().Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
			// 	}

			// 	if msg.To().Hex() == ToAddr {
			// 		isFound = true
			// 	}

			// 	_, ispending, err := client.TransactionByHash(context.Background(), tx.Hash())
			// 	if err != nil {
			// 		log.Fatal(err)
			// 	}

			// 	log.Println("is pending", ispending)

			// 	if !ispending && IsTransactionSuccess(client, tx.Hash()) {
			// 		log.Println("transaction success")
			// 	}
			// 	log.Fatal("Transaction error")

			// }
			// }
		}
	}
}

func ListenToContract() {
	client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(GATContract)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Listen to contract")

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Printf("%+v", vLog) // pointer to event log
		}
	}

}

func SendTokenTransaction() string {
	client, err := ethclient.Dial("https://ropsten.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	// privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	privateKey, err := crypto.HexToECDSA("4B1ACD7E32ADF56BC9F9416D06672850A535BFC6821A637ACD849D03A1A014DD")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("gas price:", gasPrice)

	toAddress := common.HexToAddress("0xc40c7f7c427642cc7aabc8fe523e0905b040e6e2")
	tokenAddress := common.HexToAddress("0xFD01302F6B6B5be5A71399Eb51c8282A98ed4551")

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	// fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d

	amount := new(big.Int)
	// amount.SetString("1000000000000000000", 10) // 1 tokens
	amount.SetString("1000000000000000000000000", 10) // 1 tokens
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	// fmt.Println(hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}

	gasLimit = gasLimit * 3

	log.Println("gas limit:", gasLimit)

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", signedTx.Hash().Hex()) // tx sent: 0xa56316b637a94c4cc0331c73ef26389d6c097506d581073f927275e7a6ece0bc

	return signedTx.Hash().Hex()
}
