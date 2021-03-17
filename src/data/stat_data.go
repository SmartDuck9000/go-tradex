package data

type SavedStat struct {
	Date   string  `json:"date" validate:"required,datetime_fmt"`
	Views  int64   `json:"views" validate:"omitempty,gte=0"`
	Clicks int64   `json:"clicks" validate:"omitempty,gte=0"`
	Cost   float64 `json:"cost" validate:"omitempty,gte=0"`
}

type FilterData struct {
	From    string `form:"from" validate:"required,datetime_fmt"`
	To      string `form:"to" validate:"required,datetime_fmt"`
	OrderBy string `form:"order_by" validate:"required,oneof=date views clicks cost cpc cpm"`
}

type ResultStat struct {
	Date   string  `json:"date"`
	Views  uint64  `json:"views"`
	Clicks uint64  `json:"clicks"`
	Cost   float64 `json:"cost"`
	Cpc    float64 `json:"cpc"`
	Cpm    float64 `json:"cpm"`
}
