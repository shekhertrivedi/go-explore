package questions

import (
	"fmt"
	"strconv"
)

//Sample code to run this

//prices := []int32{2, 3, 1, 2, 4, 2}
//prices := []int32{5, 1, 3, 4, 6, 2}
//prices := []int32{6, 5, 1, 3, 4, 6, 2}
//prices := []int32{1, 3, 3, 2, 5}
//finalPrice1(prices)

func finalPrice1(prices []int32) {

	indexOfItemsWithFullPrice := make([]int, len(prices), len(prices))
	discount := make([]int32, len(prices), len(prices))

	i := 0
	for i < len(prices) {
		discount[i] = 100
		i = i + 1
	}

	for i := 1; i < len(prices); i++ {
		for j := 0; j < i; j++ {
			if prices[j] >= prices[i] {
				if discount[j] == 100 {
					discount[j] = prices[i]
				}
			}
		}
	}

	i = 0
	var finalPriceSum int32

	for index, val := range discount {
		if val == 100 {
			indexOfItemsWithFullPrice[i] = index
			i = i + 1
		}
	}
	for index, val := range discount {
		if val == 100 {
			val = 0
		}
		finalPriceSum = finalPriceSum + prices[index] - val
	}

	fmt.Println(finalPriceSum)
	indexOfItemsWithFullPriceStr := ""
	for _, val := range indexOfItemsWithFullPrice {
		if val != 0 {
			indexOfItemsWithFullPriceStr = indexOfItemsWithFullPriceStr + strconv.Itoa(val) + " "
		}
	}
	if discount[0] == 100 {
		indexOfItemsWithFullPriceStr = strconv.Itoa(0) + " " + indexOfItemsWithFullPriceStr
	}
	fmt.Println(indexOfItemsWithFullPriceStr)
}

func finalPrice(prices []int) {

	indexOfItemsWithFullPrice := make([]int, len(prices), len(prices))
	discount := make([]int, len(prices), len(prices))
	for i := 1; i < len(prices); i++ {
		for j := 0; j < i; j++ {
			if prices[j] >= prices[i] {
				if discount[j] == 0 {
					discount[j] = prices[i]
				}
			}
		}
	}

	i := 0
	finalPriceSum := 0

	for index, val := range discount {
		if val == 0 {
			indexOfItemsWithFullPrice[i] = index
			i = i + 1
		}
		finalPriceSum = finalPriceSum + prices[index] - discount[index]
	}
	fmt.Println(finalPriceSum)
	indexOfItemsWithFullPriceStr := ""
	for _, val := range indexOfItemsWithFullPrice {
		if val != 0 {
			indexOfItemsWithFullPriceStr = indexOfItemsWithFullPriceStr + strconv.Itoa(val) + " "
		}
	}
	fmt.Println(indexOfItemsWithFullPriceStr)
}
