package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
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
	writer.Write([]string{"", "", "Accruals", "", "Deductions", "", "Earnings", "", "Taxes"})

	for _, pay := range pays {

		writer.Write([]string{pay.PayDate, strconv.FormatFloat(pay.NetPayCurrent, 'f', -1, 64)})

		maxAccruals := len(pay.Accruals)
		maxDeducations := len(pay.Deductions)
		maxEarnings := len(pay.Earnings)
		maxTaxes := len(pay.Taxes)
		max := int(
			math.Max(
				math.Max(
					math.Max(float64(maxAccruals), float64(maxDeducations)),
					float64(maxEarnings)),
				float64(maxTaxes)))

		line := []string{"", ""}
		for i := 0; i < max; i++ {
			// Accruals
			if i < maxAccruals {
				line = append(line, pay.Accruals[i].PlanDescription, strconv.FormatFloat(pay.Accruals[i].AmountCurrent, 'f', -1, 64))
			} else {
				line = append(line, "", "")
			}

			// Deductions
			if i < maxDeducations {
				line = append(line, pay.Deductions[i].DeductionDescription, strconv.FormatFloat(pay.Deductions[i].EmployeeAmount, 'f', -1, 64))
			} else {
				line = append(line, "", "")
			}

			// Earnings
			if i < maxEarnings {
				line = append(line, pay.Earnings[i].PayDescription, strconv.FormatFloat(pay.Earnings[i].Amount, 'f', -1, 64))
			} else {
				line = append(line, "", "")
			}

			// Taxes
			if i < maxTaxes {
				line = append(line, pay.Taxes[i].TaxDescription, strconv.FormatFloat(pay.Taxes[i].Amount, 'f', -1, 64))
			} else {
				line = append(line, "", "")
			}

			writer.Write(line)
			line = []string{"", ""}
		}

	}
}
