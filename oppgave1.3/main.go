//Echo2 print its command-line arguments
package main

import (
	"fmt"
	"os"
	"time"
	"strings"
)

func benchmark(f func(), name string, n int) {
	start := time.Now()
	for i := 0; i < n; i++ {
		f()
	}
	elapsed := time.Since(start)
	avg := elapsed.Milliseconds() / int64(n)
	fmt.Printf("%s average: %d ms over %d runs\n", name, avg, n)
}

func main() {
	benchmark(fastFunction, "Fast function", 1000)
	benchmark(slowFunction, "Slow function", 1000)
}

func fastFunction() {
	start := time.Now()
	fmt.Println("Fast function: ")
	fmt.Println(strings.Join(os.Args[1:], " "))
	elapsed := time.Since(start).Milliseconds()
	fmt.Printf("%d ms elapsed\n", elapsed)	
}

func slowFunction() {
	// Simulate a slow function
	fmt.Println("Slow function: ")
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
    elapsed := time.Since(start).Milliseconds()
	fmt.Printf("%d ms elapsed\n", elapsed)
	fmt.Println(s)
}
