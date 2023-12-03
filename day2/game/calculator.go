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

func SumGameRecordsPowers(gameResults []Results) int {
	powerResult := 0
	for _, result := range gameResults {
		powerResult += result.GetMinPow()
	}

	return powerResult
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
