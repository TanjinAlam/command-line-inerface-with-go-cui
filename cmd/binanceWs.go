/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
    "github.com/adshao/go-binance/v2"
    "github.com/adshao/go-binance/v2/futures"
	"github.com/spf13/cobra"
	// "time"
	"context"
	"os"
	"github.com/joho/godotenv"
)

var updatedListenKey string


type WsHandler func(message []byte)

type WsDepthEvent struct {
	Event         string `json:"e"`
	Time          int64  `json:"E"`
	Symbol        string `json:"s"`
	LastUpdateID  int64  `json:"u"`
	FirstUpdateID int64  `json:"U"`
	Bids          []Bid  `json:"b"`
	Asks          []Ask  `json:"a"`
}

type orderDepth struct {
	Event         string `json: "depthUpdate"`
	Time          int64
	Symbol        string
	LastUpdateID  int64
	FirstUpdateID int64
	Bids          []Bid
	Asks          []Ask
}

type Bid struct {
	Price    string
	Quantity string
}

type Ask struct {
	Price    string
	Quantity string
}

// binanceWsCmd represents the binanceWs command
var binanceWsCmd = &cobra.Command{
	Use:   "binanceWs",
	Short: "It will show all the active order ",
	Long: `Along with active order it will filter the position of each order and trade details`,
	Run: func(cmd *cobra.Command, args []string) {
		// load .env file from given path
		// we keep it empty it will load .env from current directory
		err := godotenv.Load(".env")
		//need to handle error to show in the terminal
		if err != nil {
			fmt.Println("Error loading .env file")
		}
		API_KEY := os.Getenv("API_KEY")
        API_SECRET := os.Getenv("API_SECRET")
		//testnet true for api calling
		futures.UseTestnet = true
		BinanceClient := futures.NewClient(API_KEY, API_SECRET)
		openOrders, err := BinanceClient.NewListOpenOrdersService().Symbol("LTCUSDT").
		Do(context.Background())
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, o := range openOrders {
			fmt.Println(*o)
		}

		// callling Webscoket 
		// getOrderDept("LTCUSDT")
		// userOrderDetails()
	},

}

func init() {
	rootCmd.AddCommand(binanceWsCmd)
	
}

func getOpenPosition() {
	doneC, _, err := binance.WsUserDataServe(updatedListenKey, "AccountUpdate", errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}

// It will generate listending key after 60 and update the key
func generateListendingKey() {
	listenKey, err := futuresClient.NewStartUserStreamService().Do(context.Background())
		if err != nil {
			fmt.Println(err)
			return
		}
	updatedListenKey = listenKey
}

// function to get order depth
func getOrderDept(symbol string) {
	wsDepthHandler := func(event *binance.WsDepthEvent) {
		fmt.Println(event)
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}

	doneC, _, err := binance.WsDepthServe(symbol, wsDepthHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<- doneC
}


// function to get user order details
func userOrderDetails() {
	LISTENDING_KEY := os.Getenv("LISTENDING_KEY")
	wsHandler := func(event *binance.WsUserDataEvent) {
		fmt.Println(event)
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	
	doneC, _, err := binance.WsUserDataServe(LISTENDING_KEY, wsHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC

}



