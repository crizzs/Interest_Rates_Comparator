package main

import(
	"time"
	"net/http"
	"crypto/tls"
	"io/ioutil"
	"bytes"
	"encoding/json"
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
//Http transport settings
var tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true} ,
			TLSHandshakeTimeout: 2 * time.Second,
		}
//This function energizes the financial period object
func InitiatizeFinancialPeriod(fromStr string,toStr string) FinancialPeriod{

	if fromStr == "" && toStr == "" {
		return FinancialPeriod{make([]InterestRate, 0)}
	}

	//Don't have to set the offset because the 100 rows limit is not observed by MAS
	concatURL := URL + "&between[end_of_month]=" +fromStr+ "," + toStr

	client := &http.Client{Transport: tr,}
	var query = []byte(``)

	//Retrieve Result from MAS
	req, err := http.NewRequest("GET", concatURL ,bytes.NewBuffer(query))
	if err != nil {
		//If error is encountered, returns empty object
		return FinancialPeriod{make([]InterestRate, 0)}
	}
	//Request for data from MAS API
	resp, respErr := client.Do(req)
	if respErr != nil {
		//If error is encountered, returns empty object
		return FinancialPeriod{make([]InterestRate, 0)}
	}

	masResult, _ := ioutil.ReadAll(resp.Body)

	//Unmarshal the results into struct
	var interestRatesResult InterestRatesResult;

	json.Unmarshal([]byte(masResult), &interestRatesResult)

	defer resp.Body.Close()

	total := StrToInt(interestRatesResult.Result.Total)
	
	//Initialize the Financial Period
	finResultPeriod := FinancialPeriod{make([]InterestRate,total)}

	for i := 0; i < total; i++{
		var eachRecord = interestRatesResult.Result.Records[i]
		finResultPeriod.interestRateArr[i] = CreateInterestRateObj(eachRecord.EndOfMonth,StrToFloat(eachRecord.BanksFixedDeposits3M),StrToFloat(eachRecord.BanksFixedDeposits6M),StrToFloat(eachRecord.BanksFixedDeposits12M),StrToFloat(eachRecord.BanksSavingsDeposits),StrToFloat(eachRecord.FcFixedDeposits3M),StrToFloat(eachRecord.FcFixedDeposits6M),StrToFloat(eachRecord.FcFixedDeposits12M),StrToFloat(eachRecord.FcSavingsDeposits))
		//Return the financial period retrieved
		if i == (total-1){
			return finResultPeriod
		}
	}
	
	return FinancialPeriod{make([]InterestRate, 0)}
}
//This function is to visualise every month on a plain sheet
func (fp *FinancialPeriod) VisualiseData() string{
	var str = "Date | Banks Interest Rate (Normalised) | FCs Interest Rate (Normalised) | Overall Rate (Normalised)|"

	if len(fp.interestRateArr) == 0 {
		return str +"\n\n"+ GetRepliesText(6)
	}

	for i:=0;i<len(fp.interestRateArr);i++{
		var eachIRObj = fp.interestRateArr[i]
		str += "\n" + eachIRObj.GetDisplay()
	}

	return str
}
//This function will collated the results of all months with higher FC IR
func (fp *FinancialPeriod) MonthsWithFCHigherThanBanksIR() []InterestRate{
	var collatedInterestRate []InterestRate;

	for i := 0; i<len(fp.interestRateArr);i++{
		var eachIRObj = fp.interestRateArr[i]

		objValidity := eachIRObj.CheckInterestRateValidity()

		if objValidity {
			result,_:= eachIRObj.GetHigherInterestRate()
			
			//If result is fc, we will place it inside the slice
			if result == "fc" {
				collatedInterestRate = append(collatedInterestRate,eachIRObj)
			}
		}
	}
	return collatedInterestRate
}
//This function returns the average for banks and FCs interest rate
func (fp *FinancialPeriod) RetrieveAvgOfBankAndFCRatesForPeriod() (float64,float64){
	var bankTotal = 0.0
	var fcTotal = 0.0
	//Counts the amount of valid records
	count := 0

	for i := 0; i<len(fp.interestRateArr);i++{
		var eachIRObj =  fp.interestRateArr[i]

		objValidity := eachIRObj.CheckInterestRateValidity()
		//Ensures a fair comparison (Will not compare in months which either Banks or FCs rates are having null values)

		if objValidity {
			count++
			bankTotal = bankTotal + eachIRObj.AvgBankInterestRate()
			fcTotal = fcTotal + eachIRObj.AvgFCInterestRate()
		}
	}

	if count ==0 {
		return -9999.99,-9999.99
	}

	return (bankTotal/float64(count)),(fcTotal/float64(count))
}

//This function will tell whether the interest rate in this period is in the downtrend or uptrend
//Data provided by MAS is already sorted by date
func (fp *FinancialPeriod) RetrieveIRTrendForPeriod() string{
	var startAndEndObj []InterestRate;

	if len(fp.interestRateArr) == 1 || len(fp.interestRateArr) == 0{
		return GetRepliesText(8)
	}

	findFirstValidObj := 0
	findLastValidObj := len(fp.interestRateArr)-1
	//Search from the front (Will never take last obj)
	for findFirstValidObj < len(fp.interestRateArr)-1 {

		var eachIRObj =  fp.interestRateArr[findFirstValidObj]
		objValidity := eachIRObj.CheckInterestRateValidity()

		if objValidity {
			startAndEndObj = append(startAndEndObj,eachIRObj)
			break
		}
		findFirstValidObj++
	}
	//Search from the back(Will nevr take first Obj)
	for findLastValidObj > 0 {
		var eachIRObj =  fp.interestRateArr[findLastValidObj]
		objValidity := eachIRObj.CheckInterestRateValidity()

		if objValidity {
			startAndEndObj = append(startAndEndObj,eachIRObj)
			break
		}
		findLastValidObj--
	}
	//This is when result returned is valid for trending of this period
	if len(startAndEndObj) == 2 {
		startOfPeriod := startAndEndObj[0]
		endOfPeriod := startAndEndObj[1]

		positiveOrNegativeInterestRate := endOfPeriod.AvgInterestRate() - startOfPeriod.AvgInterestRate()

		if positiveOrNegativeInterestRate > 0 {
			return "UpTrend"
		}else if positiveOrNegativeInterestRate < 0 {
			return "DownTrend"
		}else if positiveOrNegativeInterestRate == 0{
			return "Steady"
		}
	}

	return GetRepliesText(8)

}

