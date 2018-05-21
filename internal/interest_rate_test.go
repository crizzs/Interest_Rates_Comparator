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
	bankOrFCStr,_ := interestRateObj.GetHigherInterestRate()
	//Assert Value
	assert.Equal(t, bankOrFCStr,"","This should be Financial Companies")

}