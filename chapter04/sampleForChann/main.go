package main

import (
	"fmt"
	"math"
)

// 4.2.3
func primeNumber() chan int {
	result := make(chan int)
	go func() {
		result <- 2
		for i := 3; i < 1000000; i += 2 {
			l := int(math.Sqrt(float64(i)))
			found := false
			for j := 3; j < l; j += 2 {
				if i%j == 0 {
					found = true
					break
				}
			}
			if !found {
				result <- i
			}
		}
		close(result)
	}()
	return result
}

func main() {
	pn := primeNumber()
	//返ってきたチャネルは、for ... range 構文の中で配列と同じ場所に置と、「値がくるたびにfor ループが回る、個数が未定の動的配列」のように扱える
	// closeされたらループは止まる
	for n := range pn {
		fmt.Println(n)
	}
}
