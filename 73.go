package main

import "fmt"

func main() {
	var n, w, i, t int
	fmt.Scan(&n)

	t = 0
	for i = 0; i < n; i++ {
		fmt.Scan(&w)
		t = t + w

	}

	rt := t / n
	fmt.Println(t, rt)
}
