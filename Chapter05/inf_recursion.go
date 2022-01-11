package main

func countdownToZero(num int) {
	//fmt.Println(num)
	countdownToZero(num - 1)
}

func main() {
	countdownToZero(10)
}
