package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type FullChain struct {
	Chain  []Block `json:"chain"`
	Length int     ` json:"length"`
}

type Blockchain struct {
	Chain              []Block
	CurrentTransaction Transaction
	Nodes              []string
}
type Block struct {
	Index        int
	Timestamp    int64
	Transaction  Transaction
	Proof        int
	PreviousHash string
}

type Transaction struct {
	Sender    string
	Recipient string
	Amount    int
}

var workLevel = 3

func (bc *Blockchain) NewBlock(proof int, previousHash ...string) Block {
	var pg string
	if len(previousHash) != 0 {
		pg = previousHash[0]
	} else {
		pg = bc.Hash(bc.Chain[len(bc.Chain)-1])
	}
	block := Block{
		Index:        (len(bc.Chain) + 1),
		Timestamp:    time.Now().Unix(),
		Transaction:  bc.CurrentTransaction,
		Proof:        proof,
		PreviousHash: pg,
	}
	bc.CurrentTransaction = Transaction{}
	bc.Chain = append(bc.Chain, block)
	return block
}

func (bc Blockchain) LastBlock() Block {
	return bc.Chain[len(bc.Chain)-1]
}

func (bc *Blockchain) NewTransaction(sender, recipient string, amount int) int {
	bc.CurrentTransaction = Transaction{
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
	}
	return bc.LastBlock().Index + 1
}

func (bc *Blockchain) Hash(block Block) string {
	blockString, err := json.Marshal(block)
	if err != nil {
		panic(err)
	}
	hashData := sha256.Sum256([]byte(blockString))
	return hex.EncodeToString(hashData[:])
}

func (bc *Blockchain) RegisterNode(address string) error {
	res, err := http.Get(address + "/chain")
	if err != nil || res.StatusCode != http.StatusOK {
		return err
	}
	for _, v := range bc.Nodes {
		if v == address {
			return errors.New("すでに登録されています")
		}
	}
	bc.Nodes = append(bc.Nodes, address)
	return nil
}
