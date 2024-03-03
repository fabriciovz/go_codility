package solution1

// Greedy Algorithm to Find Minimum Number of Coins
func Solution(coins []int, amount int) []int {
	var change []int
	i := 0
	for i < len(coins) {
		if coins[i] <= amount {
			amount -= coins[i]
			change = append(change, coins[i])
			continue
		}
		if amount == 0 {
			break
		}
		if i == len(coins)-1 && amount > 0 {
			return []int{}
		}
		i++
	}
	return change
}
