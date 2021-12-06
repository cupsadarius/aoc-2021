package day5

import "aoc/internal/app/aoc/types"

func Part1(lines []types.Line) int {
	plane := makePlane(lines)

	for _, line := range lines {
		plane.AddStraightLine(line)
	}

	return plane.CountPointsHigherThan(1)
}

func Part2(lines []types.Line) int {
	plane := makePlane(lines)

	for _, line := range lines {
		plane.AddAllLines(line)
	}

	return plane.CountPointsHigherThan(1)
}

func makePlane(lines []types.Line) types.Plane {
	width, height := getDimensions(lines)
	points := [][]types.Point{}
	for i := 0; i <= height; i++ {
		row := make([]types.Point, width+1)
		points = append(points, row)
	}

	plane := types.Plane{Plane: points, Width: width, Height: height}
	return plane
}

func getDimensions(lines []types.Line) (int, int) {
	width, height := 0, 0
	for _, line := range lines {
		if line.A.X > width {
			width = line.A.X
		}
		if line.B.X > width {
			height = line.B.X
		}
		if line.A.Y > height {
			height = line.A.Y
		}
		if line.B.Y > height {
			height = line.B.Y
		}
	}
	return width, height
}
