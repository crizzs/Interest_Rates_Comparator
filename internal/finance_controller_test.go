package masapi
/***************
Test Business
Logic and Display
Layer for Interest
Rate Module
****************/
import (
	assert "github.com/stretchr/testify/assert"
	"testing"
	//"fmt"
)

func TestCreateFinancialPeriod(t *testing.T) {
	//Wrong Date Format
	status := CreateFinancialPeriod("01-2017","Dec-2017")
	assert.Equal(t, status,"Your Start Date is incorrect.\nThis is an example of a correct input (Jan-2017)" ,"Wrong Date")
	//Swapped Date and not valid
	status = CreateFinancialPeriod("Dec-2017","Jan-2017")
	assert.Equal(t, status,"Your End Date is before Start Date." ,"Swapped Dates and Invalid")
	//Correct
	status = CreateFinancialPeriod("Jan-2017","Dec-2017")
	assert.Equal(t, status,"Your Financial Period for Interest Rate Analysis is created." ,"It should acknowledge this program have created financial period.")
}

func TestDataVisualisation(t *testing.T){
	CreateFinancialPeriod("Jan-2099","Dec-2099")
	assert.Equal(t, VisualizeIRComparisonByMonth(),"Date | Banks Interest Rate (Normalised) | FCs Interest Rate (Normalised) | Overall Rate (Normalised)|\n\nThere is no interest rates data found!" ,"Should show appropriate wordings")

	CreateFinancialPeriod("Jan-2017","Dec-2017")
	str := VisualizeIRComparisonByMonth()

	assert.Equal(t,str ,"Date | Banks Interest Rate (Normalised) | FCs Interest Rate (Normalised) | Overall Rate (Normalised)|\nJan-2017|0.21499999999999997 percent|0.33749999999999997 percent|0.27625 percent|\nFeb-2017|0.21499999999999997 percent|0.33749999999999997 percent|0.27625 percent|\nMar-2017|0.20750000000000002 percent|0.33749999999999997 percent|0.2725 percent|\nApr-2017|0.20750000000000002 percent|0.33749999999999997 percent|0.2725 percent|\nMay-2017|0.20750000000000002 percent|0.33749999999999997 percent|0.2725 percent|\nJun-2017|0.20750000000000002 percent|0.33749999999999997 percent|0.2725 percent|\nJul-2017|0.20750000000000002 percent|0.33749999999999997 percent|0.2725 percent|\nAug-2017|0.20750000000000002 percent|0.33749999999999997 percent|0.2725 percent|\nSep-2017|0.20750000000000002 percent|0.33749999999999997 percent|0.2725 percent|\nOct-2017|0.20750000000000002 percent|0.33749999999999997 percent|0.2725 percent|\nNov-2017|0.20750000000000002 percent|0.33749999999999997 percent|0.2725 percent|\nDec-2017|0.20750000000000002 percent|0.33749999999999997 percent|0.2725 percent|" ,"It should give a generalised view of interest rates.")
}

func TestVizForFCBeatsBank(t *testing.T){
	CreateFinancialPeriod("Jan-2017","Dec-2017")
	str := VisualizeMonthsThatFCsWin()

	assert.Equal(t, str,"Date | Banks Interest Rate (Normalised) | FCs Interest Rate (Normalised) | Overall Rate (Normalised)|\nJan-2017|0.21499999999999997 percent|0.33749999999999997 percent|0.27625 percent|\nFeb-2017|0.21499999999999997 percent|0.33749999999999997 percent|0.27625 percent|\nMar-2017|0.20750000000000002 percent|0.33749999999999997 percent|0.2725 percent|\nApr-2017|0.20750000000000002 percent|0.33749999999999997 percent|0.2725 percent|\nMay-2017|0.20750000000000002 percent|0.33749999999999997 percent|0.2725 percent|\nJun-2017|0.20750000000000002 percent|0.33749999999999997 percent|0.2725 percent|\nJul-2017|0.20750000000000002 percent|0.33749999999999997 percent|0.2725 percent|\nAug-2017|0.20750000000000002 percent|0.33749999999999997 percent|0.2725 percent|\nSep-2017|0.20750000000000002 percent|0.33749999999999997 percent|0.2725 percent|\nOct-2017|0.20750000000000002 percent|0.33749999999999997 percent|0.2725 percent|\nNov-2017|0.20750000000000002 percent|0.33749999999999997 percent|0.2725 percent|\nDec-2017|0.20750000000000002 percent|0.33749999999999997 percent|0.2725 percent|" ,"Should show appropriate wordings")
}


func TestViewOverallAvgBanksVersusFCs(t *testing.T){
	//Check for different scenarios
	CreateFinancialPeriod("Jan-2099","Dec-2099")
	assert.Equal(t, ShowOverallBanksVersusFCsAvg(),"There is no data on Banks and Financial Companies Interest Rate (Overall Average)." ,"Shows overall average of banks and fcs")

	CreateFinancialPeriod("Jan-2017","Dec-2017")
	assert.Equal(t, ShowOverallBanksVersusFCsAvg(),"The overall interest rate average (Financial Period) for Banks Versus Financial Companies is as follow,\nBanks: 0.20875 percent\nFinancial Companies: 0.33749999999999997 percent\n" ,"Shows overall average of banks and fcs")
}

func TestViewTrendForSelectedFinancialPeriod(t *testing.T){
	CreateFinancialPeriod("Jan-2099","Dec-2099")
	assert.Equal(t, ShowTrend(),"Not able to tell the trending of interest rates during this period." ,"Shows trend related messages")

	CreateFinancialPeriod("Jan-2017","Dec-2017")
	assert.Equal(t, ShowTrend(),"The interest rates are trending DOWN during the defined financial period." ,"Shows trend related messages")
}