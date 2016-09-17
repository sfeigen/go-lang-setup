package main

import(
    "fmt"
)

type Vert struct {
    X, Y float64
}

func (v Vert) Abs() float64 {
    return v.X * v.Y
}

func main() {
    x := Vert{3, 4}
    fmt.Println(x.Abs())
}