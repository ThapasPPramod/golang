package main

import (
"golang.org/x/tour/tree";
"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int){
	if t.Left != nil{
		Walk(t.Left, ch)
	}
	fmt.Println("walking", t.Value)
	ch <- t.Value
	if t.Right != nil{
		Walk(t.Right, ch)
	}
	return
	
}

	
// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool{
	ch1, ch2 := make(chan int), make(chan int)
	go func(){
	Walk(t1, ch1)
	close(ch1)
	}()
	go func(){
	Walk(t2, ch2)
	close(ch2)
	}()
	for {
	value1, ok1 := <-ch1
	value2, ok2 := <-ch2
	if !ok1 || !ok2{
	return ok1 == ok2
	}
	if value1 != value2{
		return false
	}
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
	return
}
