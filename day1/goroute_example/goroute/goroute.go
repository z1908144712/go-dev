package goroute

func Add(a int, b int, c chan int) {
	c <- (a + b)
	c <- (a - b)
}