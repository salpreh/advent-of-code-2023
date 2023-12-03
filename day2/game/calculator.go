package game

import (
	"regexp"
	"strconv"
	"strings"
)

const matchSep = ";"

var (
	gameRegex, _  = regexp.Compile(`Game (\d+):`)
	blueRegex, _  = regexp.Compile(`(\d+) blue`)
	redRegex, _   = regexp.Compile(`(\d+) red`)
	greenRegex, _ = regexp.Compile(`(\d+) green`)
)

func SumPossibleGames(gamesResults []Results, gamePieces Record) int {
	result := 0
	for _, gameResults := range gamesResults {
		if gameResults.IsFeasible(gamePieces) {
			result += gameResults.Num
		}
	}

	return result
}

func SumGameRecordsPowers(gameResults []Results) int {
	powerResult := 0
	for _, result := range gameResults {
		powerResult += result.GetMinPow()
	}

	return powerResult
}

func ParseInputLines(lines []string) []Results {
	gamesResults := make([]Results, 0)
	for _, line := range lines {
		gameNum, _ := strconv.Atoi(gameRegex.FindStringSubmatch(line)[1])
		gameRecords := make([]Record, 0)
		for _, match := range strings.Split(line, matchSep) {
			gameRecord := Record{}

			blueNum := blueRegex.FindStringSubmatch(match)
			if blueNum != nil && len(blueNum) > 1 {
				gameRecord.Blue, _ = strconv.Atoi(blueNum[1])
			}

			redNum := redRegex.FindStringSubmatch(match)
			if redNum != nil && len(redNum) > 1 {
				gameRecord.Red, _ = strconv.Atoi(redNum[1])
			}

			greenNum := greenRegex.FindStringSubmatch(match)
			if greenNum != nil && len(greenNum) > 1 {
				gameRecord.Green, _ = strconv.Atoi(greenNum[1])
			}

			gameRecords = append(gameRecords, gameRecord)
		}

		gamesResults = append(gamesResults, Results{gameNum, gameRecords})
	}

	return gamesResults
}

type Results struct {
	Num     int
	Records []Record
}

func (r *Results) IsFeasible(gamePieces Record) bool {
	result := true
	for _, record := range r.Records {
		if !record.IsFeasible(gamePieces) {
			result = false
			break
		}
	}

	return result
}

func (r *Results) GetMinPow() int {
	return r.GetMinPieces().Pow()
}

func (r *Results) GetMinPieces() *Record {
	minPieces := Record{}
	for _, record := range r.Records {
		if record.Red > minPieces.Red {
			minPieces.Red = record.Red
		}
		if record.Green > minPieces.Green {
			minPieces.Green = record.Green
		}
		if record.Blue > minPieces.Blue {
			minPieces.Blue = record.Blue
		}
	}

	return &minPieces
}

type Record struct {
	Red   int
	Green int
	Blue  int
}

func (r *Record) IsFeasible(gamePieces Record) bool {
	return r.Red <= gamePieces.Red && r.Green <= gamePieces.Green && r.Blue <= gamePieces.Blue
}

func (r *Record) Pow() int {
	return r.Red * r.Green * r.Blue
}
