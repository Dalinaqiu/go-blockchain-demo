package main

import "fmt"

type Transaction struct {
	Sender    string
	Recipient string
	Amount    int
}

type Block struct {
	Index        int
	Timestamp    string
	Transactions []Transaction
	PrevHash     string
	Hash         string
}

type Blockchain struct {
	Chain []Block
}

func main() {
	fmt.Println("Some Data Constructure")
}
