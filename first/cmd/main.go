package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	jsonFile, err := os.ReadFile("files/hard.json")
	if err != nil {
		fmt.Println(err)
	}

	var numbers [][]int
	json.Unmarshal([]byte(jsonFile), &numbers)

	fmt.Println(getMaxPath(numbers))
}

func getMaxPath(data [][]int) int {
	if len(data) == 0 {
		return 0
	}

	dp := make([]int, len(data[len(data)-1]))
	copy(dp, data[len(data)-1])

	for i := len(data) - 2; i >= 0; i-- {
		for j := 0; j < len(data[i]); j++ {
			dp[j] = data[i][j] + max(dp[j], dp[j+1])
		}
		fmt.Println(dp)
	}

	return dp[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
