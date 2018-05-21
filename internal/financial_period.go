package masapi

import(
	"time"
	"net/http"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"bytes"
)
/***************
Created Model 
class for a specific
Financial Period
****************/

type FinancialPeriod struct {  
	interestRateArr []InterestRate
}

type InterestRatesResult struct {
	Success bool `json:"success"`
	Result  struct {
		ResourceID []string `json:"resource_id"`
		Limit      int      `json:"limit"`
		Total      string   `json:"total"`
		Records    []struct {
			EndOfMonth            string `json:"end_of_month"`
			PrimeLendingRate      string `json:"prime_lending_rate"`
			BanksFixedDeposits3M  string `json:"banks_fixed_deposits_3m"`
			BanksFixedDeposits6M  string `json:"banks_fixed_deposits_6m"`
			BanksFixedDeposits12M string `json:"banks_fixed_deposits_12m"`
			BanksSavingsDeposits  string `json:"banks_savings_deposits"`
			FcHirePurchaseMotor3Y string `json:"fc_hire_purchase_motor_3y"`
			FcHousingLoans15Y     string `json:"fc_housing_loans_15y"`
			FcFixedDeposits3M     string `json:"fc_fixed_deposits_3m"`
			FcFixedDeposits6M     string `json:"fc_fixed_deposits_6m"`
			FcFixedDeposits12M    string `json:"fc_fixed_deposits_12m"`
			FcSavingsDeposits     string `json:"fc_savings_deposits"`
			Timestamp             string `json:"timestamp"`
		} `json:"records"`
	} `json:"result"`
}

//This is the URL to call MAS's interest rate API (Set Max Limit to 1000)
var URL = "https://eservices.mas.gov.sg/api/action/datastore/search.json?resource_id=5f2b18a8-0883-4769-a635-879c63d3caac&limit=1000"

var tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true} ,
			TLSHandshakeTimeout: 2 * time.Second,
		}

func InitiatizeFinancialPeriod(fromStr string,toStr string) FinancialPeriod{
	//Don't have to set the offset because the 100 rows limit is not observed by MAS
	concatURL := URL + "&between[end_of_month]=" +fromStr+ "," + toStr

	client := &http.Client{Transport: tr,}
	var query = []byte(``)

	//Retrieve Result from MAS (Initial Result)
	req, err := http.NewRequest("GET", concatURL ,bytes.NewBuffer(query))
	if err != nil {
		fmt.Printf("A connection error to MAS API has occured!");
	}
	resp, err := client.Do(req)
	masResult, _ := ioutil.ReadAll(resp.Body)


	return FinancialPeriod{}
}
/* This check should go up one level

func ConvertStrToTime(dateStr string) time.Time{
	date ,err := time.Parse("Jan-2006",dateStr)
	if err != nil {
		return _
	}
	return date
}
*/