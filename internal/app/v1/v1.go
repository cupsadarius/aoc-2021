// Package v1 run the v1 of the app
package v1

import (
	day1 "aoc/internal/app/aoc/day1"
	day2 "aoc/internal/app/aoc/day2"
	day3 "aoc/internal/app/aoc/day3"
	day4 "aoc/internal/app/aoc/day4"
	"aoc/internal/app/aoc/day5"
	cli "aoc/internal/utils/cli"
	"aoc/internal/utils/convertors"
	"aoc/internal/utils/file"
	logger "aoc/internal/utils/logger"
	"aoc/internal/utils/version"
	"context"
	"fmt"
	"strconv"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

const appID = "aoc2021"

var configFile string
var inputFile string
var part string

// Run the application
func Run(ctx context.Context) error {
	cli.Init(appID, "Advent of Code 2021")
	_ = cli.AddCommand("version", "Get the application version and Git commit SHA", logVersionDetails)
	_ = cli.AddCommand("day-1", "Run Day 1", solveDay1)
	_ = cli.AddCommand("day-2", "Run Day 2", solveDay2)
	_ = cli.AddCommand("day-3", "Run Day 3", solveDay3)
	_ = cli.AddCommand("day-4", "Run Day 4", solveDay4)
	_ = cli.AddCommand("day-5", "Run Day 5", solveDay5)
	cli.AssignStringFlag(&configFile, "config", "", "config file (default is ./.config.yaml)")
	cli.AssignStringFlag(&inputFile, "input", "", "input file (default is ./inputs/input.txt)")
	cli.AssignStringFlag(&part, "part", "", "solution part")
	log := logger.New("aoc2021", "start")
	cfg, err := newConfig()
	if err != nil {
		log.Fatal("Unable to initialize config", err)
	}

	logger.Init(logger.Config{
		Level:  cfg.Logger.Level,
		Source: cfg.Logger.Source,
		Format: cfg.Logger.Format,
	})

	return cli.Run(ctx)
}

func solveDay1(parentCtx context.Context) {
	log := logger.New("aoc2021", "Day 1")
	input := file.ReadFile(inputFile)
	var result int
	switch part {
	case "1":
		result = day1.Part1(convertors.GetStringListAsInt(input))
	case "2":
		result = day1.Part2(convertors.GetStringListAsInt(input))
	}

	log.Info("Day 1 Part " + part + ": " + strconv.Itoa(result))
}

func solveDay2(parentCtx context.Context) {
	log := logger.New("aoc2021", "Day 2")
	input := file.ReadFile(inputFile)
	var result int
	switch part {
	case "1":
		result = day2.Part1(convertors.GetAsHeading(input))
	case "2":
		result = day2.Part2(convertors.GetAsHeading(input))
	}

	log.Info("Day 2 Part " + part + ": " + strconv.Itoa(result))
}

func solveDay3(parentCtx context.Context) {
	log := logger.New("aoc2021", "Day 3")
	input := file.ReadFile(inputFile)
	var result int64
	switch part {
	case "1":
		result = day3.Part1(convertors.GetAsBytes(input))
	case "2":
		result = day3.Part2(convertors.GetAsBytes(input))
	}

	log.Info("Day 3 Part " + part + ": " + strconv.Itoa(int(result)))
}

func solveDay4(parentCtx context.Context) {
	log := logger.New("aoc2021", "Day 4")
	input := file.ReadFile(inputFile)
	var result int
	switch part {
	case "1":
		result = day4.Part1(convertors.GetAsBoards(input))
	case "2":
		result = day4.Part2(convertors.GetAsBoards(input))
	}

	log.Info("Day 4 Part " + part + ": " + strconv.Itoa(int(result)))
}

func solveDay5(parentCtx context.Context) {
	log := logger.New("aoc2021", "Day 5")
	input := file.ReadFile(inputFile)
	var result int
	switch part {
	case "1":
		result = day5.Part1(convertors.GetAsLines(input))
	case "2":
		result = day5.Part2(convertors.GetAsLines(input))
	}

	log.Info("Day 5 Part " + part + ": " + strconv.Itoa(int(result)))
}

func logVersionDetails(_ context.Context) {
	log := logger.New("aoc2021", "logVersionDetails")
	log.Info(fmt.Sprintf("AppVersion=%s, GitCommit=%s", version.AppVersion, version.GitCommit))
}
