package domain

type Job struct {
	ID          int    `json:"id"`
	CompanyName string `json:"company_name"`
	HourlyWage  int    `json:"hourly_wage"`
	Description string `json:"description"`
}
