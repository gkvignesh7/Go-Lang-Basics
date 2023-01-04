// package main

// import "fmt"

// func main() {
// 	var n [10]int
// 	var i, j int
// 	for i = 0; i < 10; i++ {
// 		n[i] = i + 100
// 	}
// 	for j = 0; j < 10; j++ {
// 		fmt.Printf("Element[%d]=%d\n",j,n[j])

//		}
//	}
// package main

// import "fmt"

// type vignesh int

//	func main() {
//		var a = 201
//		var vignesh = 32
//		fmt.Println(a)
//		fmt.Printf("%T\n", a)
//		fmt.Println(vignesh)
//		fmt.Printf("%T\n", vignesh )

// package main

// import "fmt"

// func main() {
// 	x := []int{1, 2, 3, 4, 5, 6}
// 	x = append(x, 101, 102, 103)
// 	fmt.Println(x)
// 	fmt.Println(x[8])
// 	y := []int{98, 99, 97, 94, 96}
// 	x = append(x, y...)
// 	fmt.Println(y)
// 	fmt.Println(x)
// 	x=append(x[:2], x[4:]... )
// 	fmt.Println(x)

// }
// package main

// import "fmt"

// func main() {
// 	x := make([]int, 10, 100)
// 	fmt.Println(x)
// 	fmt.Println(len(x))
// 	fmt.Println(cap(x))
// 	x[7] = 7
// 	fmt.Println(x)
// 	fmt.Println(len(x))
// 	fmt.Println(cap(x))
// 	x = append(x, 55)
// 	fmt.Println(x)
// 	fmt.Println(len(x))
// 	fmt.Println(cap(x))

// }
// package main

// import "fmt"

// func main() {
// 	var x [58]string
// 	fmt.Println(len(x))
// 	fmt.Println(x)
// 	for i := 65; i <= 122; i++ {
// 		x[i-65] = string(i)

//		}
//		fmt.Println(x)
//		fmt.Println(x[48])
//	}
package main

import "fmt"

func main() {
	var x [256]byte
	for i := 0; i < 256; i++ {
		x[i] = byte(i)

	}
	for i, v := range x {
		fmt.Printf("%v-%T-%b\n", v, v, v)
		if i > 50 {
			break
		}
	}
}
