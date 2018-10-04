package problems

var coins []int

func parse(input int) int {
	breakdown(input)
	money := 0

	for i, n := range coins {
		money += n
		coins[i] = 0
	}

	return money
}

func breakdown(coin int) {
	var values []int

	if coin > 11 {
		values = append(values, coin/2, coin/3, coin/4)

		for _, n := range values {
			breakdown(n)
		}
	} else {
		coins = append(coins, coin)
	}
}
