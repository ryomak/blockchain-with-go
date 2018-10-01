package middleware

import (
"fmt"
	"github.com/labstack/echo"
	"github.com/ryomak/blockchain-with-go/blockchain"
)

func InsertBlockchainMiddleware(bc *blockchain.Blockchain) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			c.Set("BLOCKCHAIN", bc)
			return nil
		})
	}
}

func InsertIdentMiddleware(nodeIdent string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			c.Set("IDENT", &nodeIdent)
			return nil
		})
	}
}

func GetBlockchain(c echo.Context) *blockchain.Blockchain {
	fmt.Println(c.Get("BLOCKCHAIN").(*blockchain.Blockchain))
	return c.Get("BLOCKCHAIN").(*blockchain.Blockchain)
}

func GetIdent(c echo.Context) *string {
	return c.Get("IDENT").(*string)
}
