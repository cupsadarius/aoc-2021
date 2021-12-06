package types

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y  int
	Value int
}

type Line struct {
	A, B Point
}

type Plane struct {
	Plane         [][]Point
	Width, Height int
}

func (line *Line) GetPoints() []Point {
	points := []Point{}
	if !(line.A.X == line.B.X || line.A.Y == line.B.Y) {
		return points
	}
	if line.A.X == line.B.X {
		if line.A.Y < line.B.Y {
			for y := line.A.Y; y <= line.B.Y; y++ {
				points = append(points, Point{line.A.X, y, 0})
			}
		} else {
			for y := line.A.Y; y >= line.B.Y; y-- {
				points = append(points, Point{line.A.X, y, 0})
			}
		}
	} else {
		if line.A.X < line.B.X {

			for x := line.A.X; x <= line.B.X; x++ {
				points = append(points, Point{x, line.A.Y, 0})
			}
		} else {
			for x := line.A.X; x >= line.B.X; x-- {
				points = append(points, Point{x, line.A.Y, 0})
			}
		}
	}
	return points
}

func (line *Line) GetDiagonalPoints() []Point {
	points := []Point{}
	isAligned := math.Abs(float64(line.A.X-line.B.X)) == math.Abs(float64(line.A.Y-line.B.Y))

	if line.A.X == line.B.X || line.A.Y == line.B.Y || !isAligned {
		return points
	}
	xSign := -1
	ySign := -1
	diff := math.Abs(float64(line.A.X - line.B.X))
	if line.A.X < line.B.X {
		xSign = 1
	}
	if line.A.Y < line.B.Y {
		ySign = 1
	}
	points = append(points, line.A)

	for i := 1; i < int(diff); i++ {
		points = append(points, Point{X: points[i-1].X + xSign, Y: points[i-1].Y + ySign, Value: 0})
	}
	points = append(points, line.B)
	return points
}

func (plane *Plane) AddStraightLine(line Line) {
	points := line.GetPoints()
	for _, point := range points {
		plane.AddPoint(point)
	}
}

func (plane *Plane) AddDiagonalLine(line Line) {
	points := line.GetDiagonalPoints()
	for _, point := range points {
		plane.AddPoint(point)
	}
}

func (plane *Plane) AddAllLines(line Line) {
	points := line.GetPoints()
	points = append(points, line.GetDiagonalPoints()...)

	for _, point := range points {
		plane.AddPoint(point)
	}
}

func (plane *Plane) AddPoint(point Point) {
	plane.Plane[point.X][point.Y].Value += 1
}

func (plane *Plane) CountPointsHigherThan(value int) int {
	count := 0
	for _, row := range plane.Plane {
		for _, point := range row {
			if point.Value > value {
				count++
			}
		}
	}
	return count
}

func (plane *Plane) Print() {
	for _, row := range plane.Plane {
		for _, point := range row {
			fmt.Print(" ", point.Value)
		}
		fmt.Println("")
	}
}
