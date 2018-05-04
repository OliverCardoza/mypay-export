package main

type PayData []struct {
	Accruals []struct {
		AmountCurrent   float64 `json:"AmountCurrent"`
		AmountType      string  `json:"AmountType"`
		Balance         float64 `json:"Balance"`
		PlanCode        string  `json:"PlanCode"`
		PlanDescription string  `json:"PlanDescription"`
	} `json:"Accruals"`
	CompanyID      string        `json:"CompanyId"`
	CompanyName    string        `json:"CompanyName"`
	Country        string        `json:"Country"`
	DeductionTaxes []interface{} `json:"DeductionTaxes"`
	Deductions     []struct {
		BasisAmount          float64 `json:"BasisAmount"`
		DeductionCode        string  `json:"DeductionCode"`
		DeductionDescription string  `json:"DeductionDescription"`
		EmployeeAmount       float64 `json:"EmployeeAmount"`
		EmployeeAmountYtd    float64 `json:"EmployeeAmountYtd"`
		EmployerAmount       float64 `json:"EmployerAmount"`
		EmployerAmountYtd    float64 `json:"EmployerAmountYtd"`
	} `json:"Deductions"`
	Earnings []struct {
		Amount         float64     `json:"Amount"`
		AmountYtd      float64     `json:"AmountYtd"`
		Hours          float64     `json:"Hours"`
		IsShift        bool        `json:"IsShift"`
		PayCode        string      `json:"PayCode"`
		PayDescription string      `json:"PayDescription"`
		PeriodEnd      string      `json:"PeriodEnd"`
		PeriodStart    string      `json:"PeriodStart"`
		RecID          int         `json:"RecID"`
		ShiftCode      interface{} `json:"ShiftCode"`
	} `json:"Earnings"`
	NetPay []struct {
		AccountNumber string  `json:"AccountNumber"`
		AccountType   string  `json:"AccountType"`
		Amount        float64 `json:"Amount"`
	} `json:"NetPay"`
	NetPayCurrent   float64 `json:"NetPayCurrent"`
	PayDate         string  `json:"PayDate"`
	PayIdentifier   string  `json:"PayIdentifier"`
	PeriodEndDate   string  `json:"PeriodEndDate"`
	PeriodStartDate string  `json:"PeriodStartDate"`
	Taxes           []struct {
		Amount         float64 `json:"Amount"`
		AmountYtd      float64 `json:"AmountYtd"`
		BasisAmount    float64 `json:"BasisAmount"`
		TaxCode        string  `json:"TaxCode"`
		TaxDescription string  `json:"TaxDescription"`
	} `json:"Taxes"`
	TotalHours float64 `json:"TotalHours"`
}
