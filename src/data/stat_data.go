package data

type SavedStat struct {
	Date   string  `json:"date" validate:"required,datetime_fmt"`
	Views  int64   `json:"views" validate:"omitempty,gte=0"`
	Clicks int64   `json:"clicks" validate:"omitempty,gte=0"`
	Cost   float64 `json:"cost" validate:"omitempty,gte=0"`
}

type FilterData struct {
	From    string `json:"from" validate:"required,datetime_fmt"`
	To      string `json:"to" validate:"required,datetime_fmt"`
	OrderBy string `json:"order_by" validate:"required,date|views|clicks|cost|cpc|cpm"`
}

type ResultStat struct {
	Date   string
	Views  uint64
	Clicks uint64
	Cost   float64
	Cpc    float64
	Cpm    float64
}
