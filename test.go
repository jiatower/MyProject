package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    a, b, c  := 1, 2, 3
    fmt.Printf("a=%d, b=%d, c=%d\n", a, b, c)
    m, _ := strconv.Atoi(os.Args[1])
    n, _ := strconv.Atoi(os.Args[2])
    for i := 0 ; i < m ; i++ {
        for j := 0 ; j < n ; j++ {
            if a == 0 {
                a = 1
            }
            a = a + a
            if b == 0 {
                b = 1
            }
            b = b * a
            if a != 0 {
                c = b / a
            }
        }
    }
    fmt.Printf("a=%d, b=%d, c=%d\n", a, b, c)
	fmt.Printf("a")
}
