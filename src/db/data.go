package db

type SavedStat struct {
	Date   string
	Views  int
	Clicks int
	Cost   float64
}

type ResultStat struct {
	Date   string
	Views  int
	Clicks int
	Cost   float64
	Cpc    float64
	Cpm    float64
}
