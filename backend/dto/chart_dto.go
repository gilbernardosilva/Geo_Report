package dto

type ChartData struct {
	Labels []string `json:"labels"`
	Data   []int    `json:"data"`
}
