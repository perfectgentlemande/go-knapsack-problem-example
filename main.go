package main

import "fmt"

type Item struct {
	Weight int
	Price  int
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func knapsackDP(capacity int, items []Item) int {
	values := make([][]int, len(items)+1)
	for i := range values {
		values[i] = make([]int, capacity+1)
	}

	for i := 1; i <= len(items); i++ {
		for w := 1; w <= capacity; w++ {
			if items[i-1].Weight <= w {
				values[i][w] = max(values[i-1][w], values[i-1][w-items[i-1].Weight]+items[i-1].Price)
			} else {
				values[i][w] = values[i-1][w]
			}
		}
	}

	return values[len(items)][capacity]
}

func knapsackBrute(i, capacity int, items []Item) int {
	if i == 0 {
		if items[i].Weight > capacity {
			return 0
		} else {
			return items[i].Price
		}
	}

	if items[i].Weight > capacity {
		return knapsackBrute(i-1, capacity, items)
	}

	chooseCur := items[i].Price + knapsackBrute(i-1, capacity-items[i].Weight, items)
	doNotChooseCur := knapsackBrute(i-1, capacity, items)

	return max(chooseCur, doNotChooseCur)
}

func main() {
	var items []Item

	items = []Item{
		{Weight: 2, Price: 2},
		{Weight: 1, Price: 1},
		{Weight: 4, Price: 10},
		{Weight: 1, Price: 2},
		{Weight: 12, Price: 4},
	}

	fmt.Println("Brute: ", knapsackBrute(len(items)-1, 15, items))
	fmt.Println("DP: ", knapsackDP(15, items))

	items = []Item{
		{Weight: 2, Price: 2},
		{Weight: 1, Price: 1},
		{Weight: 4, Price: 10},
	}

	fmt.Println("Brute: ", knapsackBrute(len(items)-1, 4, items))
	fmt.Println("DP: ", knapsackDP(4, items))
}
