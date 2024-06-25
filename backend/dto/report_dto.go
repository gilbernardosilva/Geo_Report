package dto

type ReportCreatedDTO struct {
	Name           string     `json:"name"`
	UserID         uint64     `json:"user_id"`
	Description    string     `json:"description"`
	ReportTypeID   uint64     `json:"issue_type_id"`
	Latitude       float64    `json:"latitude"`
	Longitude      float64    `json:"longitude"`
	ReportStatusID uint       `json:"report_status_id"`
	Photos         []PhotoDTO `json:"photos"`
}

type ReportUpdateDTO struct {
	ID             uint64     `json:"report_id"`
	Name           string     `json:"name"`
	UserID         uint64     `json:"user_id"`
	Description    string     `json:"description"`
	ReportTypeID   uint64     `json:"issue_type_id"`
	Latitude       float64    `json:"latitude"`
	Longitude      float64    `json:"longitude"`
	ReportStatusID uint       `json:"report_status_id"`
	Photos         []PhotoDTO `json:"photos"`
}

type PhotoDTO struct {
	URL      string `json:"url"`
	ReportID uint64 `json:"report_id"`
}
