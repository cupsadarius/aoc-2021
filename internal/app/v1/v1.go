// Package v1 run the v1 of the app
package v1

import (
	day1 "aoc/internal/app/aoc/day1"
	day2 "aoc/internal/app/aoc/day2"
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

	log.Info("Day 1 Part " + part + ": " + strconv.Itoa(result))
}

func logVersionDetails(_ context.Context) {
	log := logger.New("aoc2021", "logVersionDetails")
	log.Info(fmt.Sprintf("AppVersion=%s, GitCommit=%s", version.AppVersion, version.GitCommit))
}
