package main

func main() {
	canWinNim(4)
}

func canWinNim(n int) bool {
	return n%4 != 0
}
