package main

import (
	"fmt"
	"github.com/evanzeller/buffer/fifobuffer"
)

func main() {
	fmt.Println("Hello Evan's World")
	buffer := fifobuffer.NewBuffer(25)
	
	fmt.Printf("Capacity: %d\t", cap(buffer.Data))
	fmt.Printf("Length: %d", len(buffer.Data))

	for i:=0; i<=30; i++ {
		buffer.Add(i)
		fmt.Println(buffer.Data)
	}

	fmt.Println(cap(buffer.Data))

	for i:=0; i<=30; i++ {
		fmt.Println(buffer.Pop())
		fmt.Println(buffer.Data)
	}

	fmt.Println(cap(buffer.Data))
	
}