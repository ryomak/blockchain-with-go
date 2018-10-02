package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/ryomak/blockchain-with-go/blockchain"
	"github.com/ryomak/blockchain-with-go/middleware"
)

func MineController(c echo.Context) error {
	bc := middleware.GetBlockchain(c)
	lastBlock := bc.LastBlock()
	lastProof := lastBlock.Proof
	proof := bc.ProofOfWork(lastProof, blockchain.WORKLEVEL)
	nodeIdentifire := middleware.GetIdent(c)
	bc.NewTransaction("0", *nodeIdentifire, 1)
	block := bc.NewBlock(proof)
	return c.JSON(http.StatusOK, block)
}

func NewTransactionController(c echo.Context) error {
	type params struct {
		Sender    string `json:"sender"`
		Recipient string `json:"recipient"`
		Amount    int    `json:"amount"`
	}
	p := new(params)
	bc := middleware.GetBlockchain(c)
	if err := c.Bind(&p); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "bind error",
		})
	}
	index := bc.NewTransaction(p.Sender, p.Recipient, p.Amount)
	if index == 0{
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "cannot added Transaction to Block",
		})
	}
	return c.JSON(http.StatusCreated, map[string]string{
		"message": "added Transaction to Block" + strconv.Itoa(index),
	})
}

func ChainController(c echo.Context) error {
	bc := middleware.GetBlockchain(c)
	var f blockchain.FullChain
	f.Chain = bc.Chain
	f.Length = len(bc.Chain)
	return c.JSON(http.StatusOK, f)
}

func RegisterNodeController(c echo.Context) error {
	type params struct {
		Nodes []string `json:"nodes"`
	}
	p := new(params)
	bc := middleware.GetBlockchain(c)
	if err := c.Bind(&p); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "bind error",
		})
	}
	for _, node := range p.Nodes {
		if err := bc.RegisterNode(node);err != nil{
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": err.Error(),
			})
		}
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "register nodes",
		"nodes":   bc.Nodes,
	})
}

func ConsensusController(c echo.Context) error {
	type response struct {
		Message string             `json:"message"`
		Chain   []blockchain.Block `json:"chain"`
	}
	bc := middleware.GetBlockchain(c)
	res := new(response)
	if bc.ResolveConflicts() {
		res.Message = "updated chain"
	} else {
		res.Message = "comfirmed chain"
	}
	res.Chain = bc.Chain
	return c.JSON(http.StatusOK, res)
}

func GetNodes(c echo.Context)error{
	bc := middleware.GetBlockchain(c)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"nodes":   bc.Nodes,
	})
}

func GetAmout(c echo.Context)error{
	bc := middleware.GetBlockchain(c)
	nodeIdent := c.QueryParam("node")
	if nodeIdent ==""{
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message":   "nodeパラメータを選択してください",
		})
	}
	return c.JSON(http.StatusOK,map[string]interface{}{
		"node":nodeIdent,
		"amount":bc.GetAmount(nodeIdent),
	})
}
