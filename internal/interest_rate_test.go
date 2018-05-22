package masapi
/***************
Test Interest 
Rate Object
****************/
import (
	assert "github.com/stretchr/testify/assert"
	"testing"
)

func TestInterestRateObject(t *testing.T) {
	//Create Interest Rate Object
	interestRateObj :=  CreateInterestRateObj("1983-01",6.75,6.80,7.13,6.50,7.15,7.30,7.70,7.21);
	//Pick 2 random object attributes for testing
	assert.Equal(t, interestRateObj.banks_fixed_deposits_6m, 6.80, "Bank 6m Deposits should be equal.")

	assert.Equal(t, interestRateObj.banks_savings_deposits, 6.50, "Bank Savings Deposits should be equal.")
}

func TestAverageInterestRate(t *testing.T) {
	//Create Interest Rate Object
	interestRateObj :=  CreateInterestRateObj("1983-01",6.75,6.80,7.13,6.50,7.15,7.30,7.70,7.21);

	assert.Equal(t, interestRateObj.AvgInterestRate(), (6.75+6.80+7.13+6.50+7.15+7.30+7.70+7.21)/8, "This should return the interest rate avg.")
}

func TestAverageBankInterestRate(t *testing.T) {
	//Create Interest Rate Object
	interestRateObj :=  CreateInterestRateObj("1983-01",6.75,6.80,7.13,6.50,7.15,7.30,7.70,7.21);

	assert.Equal(t, interestRateObj.AvgBankInterestRate(), (6.75+6.80+7.13+6.50)/4, "The bank interest average should be the same.")
}

func TestAverageFinCompaniesInterestRate(t *testing.T) {
	//Create Interest Rate Object
	interestRateObj :=  CreateInterestRateObj("1983-01",6.75,6.80,7.13,6.50,7.15,7.30,7.70,7.21);

	assert.Equal(t, interestRateObj.AvgFCInterestRate(), (7.15+7.30+7.70+7.21)/4, "The FC interest average should be the same.")
}

func TestWhichInterestRateIsHigher(t *testing.T) {
	//Create Interest Rate Object
	interestRateObj :=  CreateInterestRateObj("1983-01",6.75,6.80,7.13,6.50,7.15,7.30,7.70,7.21);
	//Return if bank or FC is higher with its interest margin
	bankOrFCStr,margin := interestRateObj.GetHigherInterestRate()
	//Assert Value
	assert.Equal(t, bankOrFCStr,"fc","This should be Financial Companies")
	assert.Equal(t, margin,0.5449999999999999,"This is the difference between FC and Bank")

}

func TestInterestRateObjValidity(t *testing.T){
	interestRateObj :=  CreateInterestRateObj("1983-01",6.75,6.80,7.13,6.50,7.15,7.30,7.70,7.21);
	invalidInterestRateObj :=  CreateInterestRateObj("1983-01",6.75,6.80,7.13,6.50,7.15,7.30,-9999.99,7.21);

	assert.Equal(t, interestRateObj.CheckInterestRateValidity(),true,"Should be valid")
	assert.Equal(t, invalidInterestRateObj.CheckInterestRateValidity(),false,"Should be valid")

}

func TestDisplay(t *testing.T){
	interestRateObj :=  CreateInterestRateObj("1983-01",6.75,6.80,7.13,6.50,7.15,7.30,7.70,7.21);
	assert.Equal(t, interestRateObj.GetDisplay(),"Jan-1983|6.795 percent|7.34 percent|7.0675 percent|","This is the display for the struct.")
}