package main

import (
	"fmt"
)

func main() {
	//fmt.Println(svg_circle_generator())
	//fmt.Println(svg_polyLine_generator())
	fmt.Println(svg_line_generator())
}

func svg_circle_generator() [][]float64 {

	//input := 1,20 2,22 3,15 3.5,25
	inputs := [][]float64{{1, 20}, {2, 22}, {3, 15}, {3.5, 25}}
	cirPoints := [][]float64{}
	x := 0.0
	y := 0.0

	for i := 0; i < len(inputs); i++ {

		x = (inputs[i][0] * 100)
		y = 300 - ((inputs[i][1]*50)/500)*50
		fmt.Println(x, y)
		cirPoints = append(cirPoints, []float64{x, y})

	}
	//fmt.Println(input[0][1], cirPoint)

	return cirPoints
}

func svg_polyLine_generator() []string {

	//input := 1,20 2,22 3,15 3.5,25
	//input := [][]float64{{1, 20}, {2, 22}, {3, 15}, {3.5, 25}}
	cirPoints := svg_circle_generator()
	var poly_line []string
	x := 0.0
	y := 0.0
	//var s string

	for i := 0; i < len(cirPoints); i++ {
		x = cirPoints[i][0]
		poly_line = append(poly_line, fmt.Sprintf("%.2f", x))

		y = cirPoints[i][1]
		poly_line = append(poly_line, fmt.Sprintf("%.2f", y))

	}
	//fmt.Println(input[0][1], cirPoint)

	return poly_line
}

func svg_line_generator() []string {

	//input := 1,20 2,22 3,15 3.5,25
	//input := [][]float64{{1, 20}, {2, 22}, {3, 15}, {3.5, 25}}
	cirPoints := svg_circle_generator()
	var lines []string

	//var s string

	for i := 0; i < len(cirPoints)-1; i++ {

		x1 := cirPoints[i][0]
		y1 := cirPoints[i][1]
		x2 := cirPoints[i+1][0]
		y2 := cirPoints[i+1][1]

		lines = append(lines, fmt.Sprintf(`x1="%v" y1="%v" x2="%v" y2="%v"`, x1, y1, x2, y2))

	}
	//fmt.Println(input[0][1], cirPoint)

	return lines
}
