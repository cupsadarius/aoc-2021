package day2

import (
	types "aoc/internal/app/aoc/types"
)

func Part1(input []types.Heading) int {
	position := types.Position{X: 0, Y: 0, Aim: 0}
	for _, heading := range input {
		switch heading.Direction {
		case "forward":
			position.X += heading.Change
		case "up":
			position.Y -= heading.Change
		case "down":
			position.Y += heading.Change
		}
	}
	return position.X * position.Y
}

func Part2(input []types.Heading) int {
	position := types.Position{X: 0, Y: 0, Aim: 0}
	for _, heading := range input {
		switch heading.Direction {
		case "forward":
			position.X += heading.Change
			position.Y += heading.Change * position.Aim
		case "up":
			position.Aim -= heading.Change
		case "down":
			position.Aim += heading.Change
		}
	}
	return position.X * position.Y
}
