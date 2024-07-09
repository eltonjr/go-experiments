package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
)

var kept []byte

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Run with 'discard' or 'keep' argument")
		os.Exit(1)
		return
	}

	mode := os.Args[1]

	mem, err := os.Create(fmt.Sprintf("mem-%s.pprof", mode))
	if err != nil {
		fmt.Println("Error creating mem profile:", err)
		os.Exit(1)
		return
	}
	defer mem.Close()

	for i := 0; i < 3; i++ {
		switch mode {
		case "discard":
			fmt.Println("Running in discard mode")
			arr := bigalloc()
			fmt.Printf("Size of arr: %d\n", len(arr))

		case "keep":
			fmt.Println("Running in keep mode")
			arr := bigalloc()
			kept = arr[:1]
			fmt.Printf("Size of arr: %d\n", len(kept))

		default:
			fmt.Println("Invalid mode")
		}

		runtime.GC()
	}

	if err := pprof.WriteHeapProfile(mem); err != nil {
		fmt.Println("Error writing mem profile:", err)
		os.Exit(1)
		return
	}
}

func bigalloc() []byte {
	return make([]byte, 10<<30)
}
