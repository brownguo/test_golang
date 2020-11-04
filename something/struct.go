package main

import "fmt"

type Coordinate struct {
	X,Y float32
}

func (c *Coordinate) GetCoordinate()  {
	fmt.Printf("%.2f,%.2f,%T\n",c.X,c.Y,c)
}
func main() {
	c := Coordinate{0.1,0.2}
	c.GetCoordinate()
}
