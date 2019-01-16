package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

const mobileUrl = "/mobile/app/services/MobileService.svc/GetYourPayHistory"

type Config struct {
	LoginToken   string
	Url          string
	NumberOfPays string
}

func main() {
	raw, _ := ioutil.ReadFile("./config.json")
	config := Config{}
	json.Unmarshal(raw, &config)

	pays := GetPay(config)

	GenerateCSV(pays)
}

func GetPay(config Config) PayData {
	var jsonStr = []byte(`{"page": "1", "limit": "` + config.NumberOfPays + `"}`)
	req, err := http.NewRequest("POST", config.Url+mobileUrl, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err.Error())
	}
	cookie := http.Cookie{Name: "loginToken", Value: config.LoginToken}
	req.AddCookie(&cookie)
	req.Header.Set("Content-Type", "application/json")

	var client = &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var data PayData
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error reading post data: " + err.Error() + " ::: " + string(body))
	}

	return data
}

func GenerateCSV(pays PayData) {
	file, _ := os.Create("pay.csv")
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header fields
	writer.Write([]string{"Date", "PayId", "PayType", "Name", "Amount"})

	for _, pay := range pays {
		writer.Write([]string{pay.PayDate, pay.PayIdentifier, "NET", "Net Pay", FormatFloat(pay.NetPayCurrent)})

		for _, earning := range pay.Earnings {
			if earning.Amount != 0 {
				writer.Write([]string{pay.PayDate, pay.PayIdentifier, "EARNING", earning.PayDescription, FormatFloat(earning.Amount)})
			}
		}
		for _, accrual := range pay.Accruals {
			if accrual.AmountCurrent != 0 {
				writer.Write([]string{pay.PayDate, pay.PayIdentifier, "ACCRUAL", accrual.PlanDescription, FormatFloat(accrual.AmountCurrent)})
			}
		}
		for _, deduction := range pay.Deductions {
			if deduction.EmployeeAmount != 0 {
				writer.Write([]string{pay.PayDate, pay.PayIdentifier, "DEDUCTION", deduction.DeductionDescription, FormatFloat(deduction.EmployeeAmount)})
			}
		}
		for _, deduction := range pay.DeductionTaxes {
			if deduction.EmployeeAmount != 0 {
				writer.Write([]string{pay.PayDate, pay.PayIdentifier, "DEDUCTION_TAXES", deduction.Description, FormatFloat(deduction.EmployeeAmount)})
			}
		}
		for _, tax := range pay.Taxes {
			if tax.Amount != 0 {
				writer.Write([]string{pay.PayDate, pay.PayIdentifier, "TAXES", tax.TaxDescription, FormatFloat(tax.Amount)})
			}
		}
	}
}

func FormatFloat(floatAmount float64) string {
	return strconv.FormatFloat(floatAmount, 'f', -1, 64)
}
