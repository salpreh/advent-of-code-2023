package game

func SumPossibleGames(gamesResults []Results, gamePieces Record) int {
	result := 0
	for _, gameResults := range gamesResults {
		if gameResults.IsFeasible(gamePieces) {
			result += gameResults.Num
		}
	}

	return result
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

type Record struct {
	Red   int
	Green int
	Blue  int
}

func (r *Record) IsFeasible(gamePieces Record) bool {
	return r.Red <= gamePieces.Red && r.Green <= gamePieces.Green && r.Blue <= gamePieces.Blue
}
