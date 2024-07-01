package dto

type ChartData struct {
	Labels   []string         `json:"labels"`
	Datasets []ReportTypeData `json:"datasets"`
}

type ReportTypeData struct {
	Label           string   `json:"label"`
	Data            []int    `json:"data"`
	BackgroundColor []string `json:"backgroundColor"`
	HoverOffset     int      `json:"hoverOffset"`
}
