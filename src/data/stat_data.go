package data

type SavedStat struct {
	Date   string  `json:"date" validate:"required, format=datetime"`
	Views  uint64  `json:"views" validate:"optional"`
	Clicks uint64  `json:"clicks" validate:"optional"`
	Cost   float64 `json:"cost" validate:"optional, gte=0"`
}

type ResultStat struct {
	Date   string
	Views  uint64
	Clicks uint64
	Cost   float64
	Cpc    float64
	Cpm    float64
}
