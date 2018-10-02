package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func ValidProof(lastProof, proof int, level int) bool {
	guess := []byte(strconv.Itoa(lastProof) + strconv.Itoa(proof))
	hashData := sha256.Sum256(guess)
	guessHash := hex.EncodeToString(hashData[:])
	return guessHash[:level] == strings.Repeat("0", level)
}

func (bc *Blockchain) ValidChain(chain []Block) bool {
	lastBlock := chain[0]
	currentIndex := 1
	for currentIndex < len(chain) {
		block := chain[currentIndex]
		if block.PreviousHash != bc.Hash(lastBlock) {
			return false
		}
		if !ValidProof(lastBlock.Proof, block.Proof, WORKLEVEL) {
			guess := []byte(strconv.Itoa(lastBlock.Proof) + strconv.Itoa(block.Proof))
			hashData := sha256.Sum256(guess)
			guessHash := hex.EncodeToString(hashData[:])
			return false
		}
		lastBlock = block
		currentIndex++
	}
	return true
}

func (bc Blockchain) ProofOfWork(lastProof, level int) int {
	proof := 0
	for !ValidProof(lastProof, proof, level) {
		proof++
	}
	return proof
}

func (bc *Blockchain) ResolveConflicts() bool {
	neighbours := bc.Nodes
	newChain := []Block{}

	maxLength := len(bc.Chain)
	for _, node := range neighbours {
		res, err := http.Get(node + "/chain")
		if err != nil {
			log.Println(err.Error())
			continue
		}
		if res.StatusCode != http.StatusOK {
			log.Println(err.Error())
			continue
		}
		var fullChain FullChain
		if err := json.NewDecoder(res.Body).Decode(&fullChain); err != nil {
			log.Println(err.Error())
			continue
		}
		log.Println(fullChain.Length, ":", maxLength, ":", bc.ValidChain(fullChain.Chain))
		if fullChain.Length > maxLength && bc.ValidChain(fullChain.Chain) {
			maxLength = fullChain.Length
			newChain = fullChain.Chain
		}
	}
	if len(newChain) != 0 {
		bc.Chain = newChain
		return true
	}
	return false
}
