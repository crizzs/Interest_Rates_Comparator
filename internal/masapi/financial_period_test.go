package masapi
/***************
Test Financial
Period Class
****************/
import (
	assert "github.com/stretchr/testify/assert"
	"testing"
)

func TestFinancialPeriod(t *testing.T) {
	testEmptyPeriod := InitiatizeFinancialPeriod("","")

	assert.Equal(t, len(testEmptyPeriod.interestRateArr), 0, "Should initialise an empty arr")

	emptyFinancialPeriod := InitiatizeFinancialPeriod("2099-05","2099-10")

	assert.Equal(t, len(emptyFinancialPeriod.interestRateArr),0 , "It should contain a total of 0 records.")

	sampleFinancialPeriod := InitiatizeFinancialPeriod("2017-01","2018-04")

	//Test if the amount of results is the same as requested
	assert.Equal(t, len(sampleFinancialPeriod.interestRateArr), 16, "There should be 16 results coming from MAS.")
}

func TestVisualiseEntireFinPeriod(t *testing.T) {
	emptyFinancialPeriod := InitiatizeFinancialPeriod("2099-05","2099-10")

	assert.Equal(t, emptyFinancialPeriod.VisualiseData(), "Date | Bank FD 3m | Bank FD 6m | Bank FD 12m | Bank Savings Deposits | FC FD 3m | FC FD 6m | FC FD 12m | FC Savings Deposits |\n\nThere is no interest rates data found!", "It should contain a total of 0 records.")

	sampleFinancialPeriod := InitiatizeFinancialPeriod("2017-01","2018-04")

	assert.Equal(t, sampleFinancialPeriod.VisualiseData(),"Date | Bank FD 3m | Bank FD 6m | Bank FD 12m | Bank Savings Deposits | FC FD 3m | FC FD 6m | FC FD 12m | FC Savings Deposits |\nJan-2017|0.15 percent|0.2 percent|0.33 percent|0.18 percent|0.3 percent|0.38 percent|0.5 percent|0.17 percent|\nFeb-2017|0.15 percent|0.2 percent|0.33 percent|0.18 percent|0.3 percent|0.38 percent|0.5 percent|0.17 percent|\nMar-2017|0.14 percent|0.2 percent|0.33 percent|0.16 percent|0.3 percent|0.38 percent|0.5 percent|0.17 percent|\nApr-2017|0.14 percent|0.2 percent|0.33 percent|0.16 percent|0.3 percent|0.38 percent|0.5 percent|0.17 percent|\nMay-2017|0.14 percent|0.2 percent|0.33 percent|0.16 percent|0.3 percent|0.38 percent|0.5 percent|0.17 percent|\nJun-2017|0.14 percent|0.2 percent|0.33 percent|0.16 percent|0.3 percent|0.38 percent|0.5 percent|0.17 percent|\nJul-2017|0.14 percent|0.2 percent|0.33 percent|0.16 percent|0.3 percent|0.38 percent|0.5 percent|0.17 percent|\nAug-2017|0.14 percent|0.2 percent|0.33 percent|0.16 percent|0.3 percent|0.38 percent|0.5 percent|0.17 percent|\nSep-2017|0.14 percent|0.2 percent|0.33 percent|0.16 percent|0.3 percent|0.38 percent|0.5 percent|0.17 percent|\nOct-2017|0.14 percent|0.2 percent|0.33 percent|0.16 percent|0.3 percent|0.38 percent|0.5 percent|0.17 percent|\nNov-2017|0.14 percent|0.2 percent|0.33 percent|0.16 percent|0.3 percent|0.38 percent|0.5 percent|0.17 percent|\nDec-2017|0.14 percent|0.2 percent|0.33 percent|0.16 percent|0.3 percent|0.38 percent|0.5 percent|0.17 percent|\nJan-2018|0.15 percent|0.22 percent|0.34 percent|0.16 percent|0.3 percent|0.38 percent|0.5 percent|0.17 percent|\nFeb-2018|0.15 percent|0.22 percent|0.34 percent|0.16 percent|0.3 percent|0.38 percent|0.5 percent|0.17 percent|\nMar-2018|0.15 percent|0.22 percent|0.34 percent|0.16 percent|0.3 percent|0.38 percent|0.5 percent|0.17 percent|\nApr-2018|0.15 percent|0.22 percent|0.34 percent|0.16 percent|0.3 percent|0.38 percent|0.5 percent|0.17 percent|", "It should contain a total of a few records.")
}

func TestPeriodForFCHigherThanBanks(t *testing.T) {
	sampleFinancialPeriod := InitiatizeFinancialPeriod("2017-01","2018-03")

	assert.Equal(t, len(sampleFinancialPeriod.MonthsWithFCHigherThanBanksIR()),15 , "It should contain a total of 15 records.")
}

func TestRetrieveAvgOfBankAndFCRatesForPeriod(t *testing.T){

	sampleFinancialPeriod := InitiatizeFinancialPeriod("2017-01","2018-03")
	bankAvgRate,fcAvgRate := sampleFinancialPeriod.RetrieveAvgOfBankAndFCRatesForPeriod()

	assert.Equal(t, bankAvgRate,0.21049999999999996 , "It should be a float number.")
	assert.Equal(t, fcAvgRate,0.3375000000000001 , "It should be a float number.")
}

func TestPredictTrendFunction(t *testing.T){
	sampleWeirdFinancialPeriod := InitiatizeFinancialPeriod("2017-01","2017-01")

	badResult := sampleWeirdFinancialPeriod.RetrieveIRTrendForPeriod()
	assert.Equal(t, badResult,"Not able to tell the trending of interest rates during this period.", "This should return a message telling it is not able to see trend.")
	//Test all scenario across different period
	sampleFinancialPeriod := InitiatizeFinancialPeriod("2017-01","2018-03")
	goodResult := sampleFinancialPeriod.RetrieveIRTrendForPeriod()
	assert.Equal(t, goodResult,"UpTrend" , "This should return the UpTrend for this period.")

	sampleFinancialPeriod = InitiatizeFinancialPeriod("2013-01","2014-03")
	goodResult = sampleFinancialPeriod.RetrieveIRTrendForPeriod()
	assert.Equal(t, goodResult,"Steady" , "This should return the Steady for this period.")

	sampleFinancialPeriod = InitiatizeFinancialPeriod("2013-01","2013-05")
	goodResult = sampleFinancialPeriod.RetrieveIRTrendForPeriod()
	assert.Equal(t, goodResult,"DownTrend" , "This should return the DownTrend for this period.")
}