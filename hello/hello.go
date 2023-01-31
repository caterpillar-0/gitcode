//-----------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Print("hello world!")
// }

//-----------------------------------------------------------------------------

package main

import "fmt"

type point struct {
	x, y int
}

func main() {

	s := "hello"
	n := 123
	p := point{1, 2}
	fmt.Println(s, n) //hello 123
	fmt.Println(p)    // {1 2}

	fmt.Printf("s=%v\n", s)  // s=hello
	fmt.Printf("n=%v\n", n)  // n=123
	fmt.Printf("p=%v\n", p)  // p={1 2}
	fmt.Printf("p=%+v\n", p) // p={x:1 y:2}
	fmt.Printf("p=%#v\n", p) // p=main.point{x:1, y:2}

}

//print原样输出
//printf格式化输出
//println值+空格+换行输出

//--------------------------------------------
