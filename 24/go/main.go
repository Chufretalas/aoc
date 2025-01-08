package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().UnixMicro()
	D6P2()
	fmt.Printf("Exec time: %d μs\n", time.Now().UnixMicro()-start)
}
