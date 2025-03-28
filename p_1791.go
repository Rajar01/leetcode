package main

import "fmt"

func FindCenter(edges [][]int) int {
	al := map[int][]int{}

	for _, v := range edges {
		al[v[0]] = append(al[v[0]], v[1])
		al[v[1]] = append(al[v[1]], v[0])
	}

	for k, v := range al {
		if len(v) == (len(al) - 1) {
			return k
		}
	}

	return -1
}

func main() {
	result := FindCenter([][]int{{1, 2}, {2, 3}, {4, 2}})

	fmt.Println(result)
}
