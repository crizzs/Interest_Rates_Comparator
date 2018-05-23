package masapi


/***************
Business
Display and Logic 
Layer for Interest
Rate Module
****************/

//Initialised a Financial for controller
var customisedFinancialPeriod = InitiatizeFinancialPeriod("","")
//This attribute is stop situation where financial period yields 0 result
var customisedFinancialPeriodSize = 0
/*
Function creates the financial period for user
*/
func CreateFinancialPeriod(fromDateStr string,toDateStr string) string{
	//Convert and Check validity of from and to dates
	fromDateFormatted,fromDateValidity := ConvertStrToDate(fromDateStr)
	toDateFormatted,toDateValidity := ConvertStrToDate(toDateStr)

	if fromDateValidity == false && toDateValidity == false{
		return GetRepliesText(1)
	}else if fromDateValidity == false{
		return GetRepliesText(2)
	}else if toDateValidity == false{
		return GetRepliesText(3)
	}
	//Check if the End Date is before Start Date
	isAfter := TestFromAndToDateValidity(fromDateStr,toDateStr);
	//To Date is after From Date yield false -> Detect
	if isAfter == false{
		return GetRepliesText(4)
	}
	//After validate, we shall create the financial period
	customisedFinancialPeriod = InitiatizeFinancialPeriod(fromDateFormatted,toDateFormatted)
	customisedFinancialPeriodSize = len(customisedFinancialPeriod.interestRateArr)

	return GetRepliesText(5)
}
/*
Function to list all interest rate comparison by months
*/
func VisualizeIRComparisonByMonth() string{
	return customisedFinancialPeriod.VisualiseData();
}
/*
Function to sieve out display months that fc beats banks for interest rates
*/
func VisualizeMonthsThatFCsWin() string{
	listOfFCWinningMonth := customisedFinancialPeriod.MonthsWithFCHigherThanBanksIR()

	var str = "Date | Banks Interest Rate (Normalised) | FCs Interest Rate (Normalised) | Overall Rate (Normalised)|"

	if len(listOfFCWinningMonth) == 0 {
		return str +"\n\n"+ GetRepliesText(6)
	}

	for i:=0;i<len(listOfFCWinningMonth);i++{
		var eachIRObj = listOfFCWinningMonth[i]
		str += "\n" + eachIRObj.GetDisplay()
	}

	return str
}
/*
Function to compare the average interest rates
*/
func ShowOverallBanksVersusFCsAvg() string{
	bankAvg,fcAvg := customisedFinancialPeriod.RetrieveAvgOfBankAndFCRatesForPeriod()
	//This shows the data retrieved is invalid and empty
	if bankAvg == -9999.99 && fcAvg == -9999.99{
		return GetRepliesText(7)
	}

	return "The overall interest rate average (Financial Period) for Banks Versus Financial Companies is as follow,\nBanks: "+FloatToStr(bankAvg)+" percent\nFinancial Companies: "+FloatToStr(fcAvg)+" percent\n"
}
/*
Function to display trend for stated financial period.
*/
func ShowTrend() string{
	str := customisedFinancialPeriod.RetrieveIRTrendForPeriod()

	if str == "UpTrend"{
		return GetRepliesText(9)
	}else if str == "DownTrend"{
		return GetRepliesText(10)
	}else if str == "Steady"{
		return GetRepliesText(11)
	}

	return str
}