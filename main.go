package main

import (
	"flag"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/ryomak/blockchain-with-go/blockchain"
	"github.com/ryomak/blockchain-with-go/controller"
	"github.com/ryomak/blockchain-with-go/middleware"
)

var port = flag.String("p", "5000", "port option")

var bc *blockchain.Blockchain
var nodeIdentifire string

func init() {
	flag.Parse()
	bc = new(blockchain.Blockchain)
	bc.NewBlock(100, "1")
}

func main() {
	e := echo.New()
	nodeIdentifire = strings.Replace(uuid.New().String(), "-", "", -1)

	e.Use(mw.Logger())
	e.Use(middleware.InsertBlockchainMiddleware(bc))
	e.Use(middleware.InsertIdentMiddleware(nodeIdentifire))
	e.GET("/mine", controller.MineController)
	e.POST("/transactions/new", controller.NewTransactionController)
	e.GET("/chain", controller.ChainController)

	e.POST("/nodes/register", controller.RegisterNodeController)
	e.GET("/nodes/resolve", controller.ConsensusController)

	e.Start(":" + *port)
}
