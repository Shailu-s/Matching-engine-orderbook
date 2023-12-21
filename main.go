package main

import (
	"fmt"

	"github.com/Shailu-s/Matching-engine-orderbook/orderbook"

	"github.com/shopspring/decimal"
)

func main() {
	// Create a new order book
	ob := orderbook.NewOrderBook()

	// Example: Adding a buy (bid) order
	buyOrderID := "buy123"
	buyQuantity := decimal.NewFromFloat(10.0)
	buyPrice := decimal.NewFromFloat(100.0)
	buySide := orderbook.Buy

	// Process the buy order
	doneBuy, partialBuy, partialQtyBuy, errBuy := ob.ProcessLimitOrder(buySide, buyOrderID, buyQuantity, buyPrice)
	if errBuy != nil {
		fmt.Println("Error placing buy order:", errBuy)
		return
	}

	// Example: Adding a sell (ask) order
	sellOrderID := "sell456"
	sellQuantity := decimal.NewFromFloat(5.0)
	sellPrice := decimal.NewFromFloat(95.0)
	sellSide := orderbook.Sell

	// Process the sell order
	doneSell, partialSell, partialQtySell, errSell := ob.ProcessLimitOrder(sellSide, sellOrderID, sellQuantity, sellPrice)
	if errSell != nil {
		fmt.Println("Error placing sell order:", errSell)
		return
	}

	// Print the results
	fmt.Println("Buy Order Results:")
	fmt.Println("Completed Orders:", doneBuy)
	fmt.Println("Partial Order:", partialBuy)
	fmt.Println("Partial Quantity Processed:", partialQtyBuy)

	fmt.Println("\nSell Order Results:")
	fmt.Println("Completed Orders:", doneSell)
	fmt.Println("Partial Order:", partialSell)
	fmt.Println("Partial Quantity Processed:", partialQtySell)

	// Print the updated order book depth
	asks, bids := ob.Depth()
	fmt.Println("\nOrder Book Depth:")
	fmt.Println("Asks (Sell Orders):", asks)
	fmt.Println("Bids (Buy Orders):", bids)
}
