package main

import "fmt"

func main() {
	prices := [][]int{
		{7, 1, 5, 3, 6, 4},
		{7, 6, 4, 3, 1},
	}
	for _, p := range prices {
		profit := maxTheStock(p)
		fmt.Printf("profit for %v: %d\n", p, profit)
	}
}

func maxTheStock(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	var (
		minPrice  = prices[0]
		maxProfit = 0
	)
	for _, val := range prices {
		if val < minPrice {
			minPrice = val
		}
		if currProfit := val - minPrice; currProfit > maxProfit {
			maxProfit = currProfit
		}
	}
	return maxProfit
}
