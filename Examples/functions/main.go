package main

func add(x, y int) {
	return x + y
}

func sayhello() {
	fmt.println("hello")
}

func main() {
	fmt.println(add(42, 23))
	sayhello()
}
