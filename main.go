package main

import "sync"

var g sync.WaitGroup

func main() {
	go try1()
	go try2()
	g.Wait()
	print("All coroutines finished")
}

func try2() {
	g.Add(1)
	defer print("Go2")
	print("Hello World2!!")
	g.Done()
}

func try1() {
	g.Add(1)
	defer print("Go1")
	print("Hello World1!!")
	g.Done()
}
func try3() {

}
