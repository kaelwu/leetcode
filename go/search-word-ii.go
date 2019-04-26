package main

func findWords(board [][]byte, words []string) []string {
	res := []string{}
	for _, v := range words {
		r = match(v, board)
		if r {
			res = append(res, v)
		}
	}
	return res
}

func match(s string, board [][]byte) bool {
}
