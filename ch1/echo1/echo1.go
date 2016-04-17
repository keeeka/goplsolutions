// Echo1 prints its comman-line arguments.
package main

import (
	"fmt"
	"os"
	"time"
  "strings"
)

func main() {
	start := time.Now()
  fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("%.2fns elapsed\n", time.Since(start).Nanoseconds())
}
