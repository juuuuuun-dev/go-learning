package main

import "fmt"

type Vertex struct {
	X, Y int
}

// 一般的な関数
func Area(v Vertex) int {
	return v.X * v.Y
}

// メソッド
func (v *Vertex) Area() int {
	return v.X * v.Y
}

func (v *Vertex) Scale(i int) {
	v.X = v.X * i
	v.Y = v.Y * i
}

func main() {
	v := Vertex{3, 4}
	// fmt.Println(Area(v))
	v.Scale(10)
	fmt.Println(v.Area())
}
