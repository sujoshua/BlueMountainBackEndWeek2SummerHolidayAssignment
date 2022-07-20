package main

import "fmt"

type node struct {
	x int
	y int
}

func main() {
	var n, m, x, y int
	_, err := fmt.Scan(&n)
	_, err = fmt.Scan(&m)
	_, err = fmt.Scan(&x)
	_, err = fmt.Scan(&y)
	if err != nil {
		panic("Input Error")
		return
	}
	HorseTraversal(n, m, x, y)
}

func HorseTraversal(n, m, x, y int) {
	resultMap := make([][]int, n)

	visitedMap := make([][]bool, n)
	for i := 0; i < n; i++ {
		resultMap[i] = make([]int, m)
		visitedMap[i] = make([]bool, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			resultMap[i][j] = -1
		}

	}
	horseBfs(y-1, x-1, n, m, resultMap, visitedMap)

	for i := 0; i < n; i++ {

		for j := 0; j < m-1; j++ {
			fmt.Printf("%-5v", resultMap[i][j])
		}
		fmt.Printf("%-5v\n", resultMap[i][m-1])
	}

}

func horseBfs(x, y, n, m int, resultMap [][]int, visitedMap [][]bool) {
	var queueX []int
	var queueY []int

	queueX = append(queueX, x)
	queueY = append(queueY, y)
	resultMap[y][x] = 0
	visitedMap[y][x] = true

	for {
		if len(queueX) == 0 {
			return
		}

		for i := 0; i < 8; i++ {
			tempX := queueX[0]
			tempY := queueY[0]
			if i == 0 || i == 5 {
				tempX -= 1
			}
			if i == 1 || i == 4 {
				tempX += 1
			}
			if i == 2 || i == 3 {
				tempX += 2
			}
			if i == 6 || i == 7 {
				tempX -= 2
			}
			if i == 0 || i == 1 {
				tempY -= 2
			}
			if i == 2 || i == 7 {
				tempY -= 1
			}
			if i == 4 || i == 5 {
				tempY += 2
			}
			if i == 3 || i == 6 {
				tempY += 1
			}

			if tempY < 0 || tempY > n-1 || tempX < 0 || tempX > m-1 {
				continue
			}
			if visitedMap[tempY][tempX] == false {
				visitedMap[tempY][tempX] = true
				resultMap[tempY][tempX] = resultMap[queueY[0]][queueX[0]] + 1
				queueX = append(queueX, tempX)
				queueY = append(queueY, tempY)

			}

		}
		queueX = queueX[1:]
		queueY = queueY[1:]
	}

}
