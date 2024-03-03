package solution2

// Greedy Algorithm to Find Minimum Number of Coins
func Solution(coins []int, amount int) int {

	// Create the ways array to 1 plus the amount
	// to stop overflow
	ways := make([]int, amount+1)

	// Set the first way to 1 because its 0 and
	// there is 1 way to make 0 with 0 coins
	ways[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := 1; j < len(ways); j++ {
			if coins[i] <= j {
				ways[j] += ways[j-coins[i]]
			}
		}
	}
	return ways[amount]
}
