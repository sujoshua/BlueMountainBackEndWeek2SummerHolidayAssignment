package main

import "fmt"

func main1() {
	n := 0

	_, err := fmt.Scan(&n)
	if err != nil {
		panic("Input Value Error")
		return
	}

	a := make([]int, n)

	dfs(0, 1, 1, n, a)
}

func dfs(sum, depth, qos int, n int, a []int) {
	//如果搜索超出，则返回
	if sum > n {
		return
	}

	if sum == n {
		for j := 0; j < depth-2; j++ {
			fmt.Printf("%v+", a[j])
		}
		fmt.Println(a[depth-2])
		return
	}

	for i := qos; i <= n-1; i++ {
		a[depth-1] = i
		dfs(sum+i, depth+1, i, n, a)
		a[depth-1] = 0
	}
}
