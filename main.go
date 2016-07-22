package main

import "fmt"

func main() {
	//goDefer()
	//goPanic()
	f()
}
func goDefer() {
	fmt.Println("Start Function")
	err := fmt.Errorf("I have failed you!")

	defer func() {
		fmt.Println("Start Defer")
		if err != nil {
			fmt.Println("Defer Error:", err)
		}
	}()

	fmt.Println("End Test")
}

func goPanic() {
	defer func() {
		fmt.Println("Start Panic Defer")

		if r := recover(); r != nil {
			fmt.Println("Defer Panic:", r)
		}
	}()

	fmt.Println("Start Test")
	panic("I am the trouble")
}
