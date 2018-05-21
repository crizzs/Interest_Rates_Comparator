package masapi

/***************
Created Model 
class for banks and 
Financial Companies
Interest Rate
****************/
type InterestRate struct {  
	end_of_month string
	banks_fixed_deposits_3m float64
	banks_fixed_deposits_6m float64
	banks_fixed_deposits_12m float64
	banks_savings_deposits float64
	fc_fixed_deposits_3m float64
	fc_fixed_deposits_6m float64
	fc_fixed_deposits_12m float64
	fc_savings_deposits float64
}

//Creates an Interest Rate Object
func CreateInterestRateObj(end_of_month string,banks_fixed_deposits_3m float64,banks_fixed_deposits_6m float64,banks_fixed_deposits_12m float64,banks_savings_deposits float64,fc_fixed_deposits_3m float64,fc_fixed_deposits_6m float64,fc_fixed_deposits_12m float64,fc_savings_deposits float64)  InterestRate{	
	return InterestRate{end_of_month,banks_fixed_deposits_3m,banks_fixed_deposits_6m,banks_fixed_deposits_12m,banks_savings_deposits,fc_fixed_deposits_3m,fc_fixed_deposits_6m,fc_fixed_deposits_12m,fc_savings_deposits}
}

//A Struct function to find Bank Average Interest Rate for the month
func (ir *InterestRate) AvgBankInterestRate() float64{
	return (ir.banks_fixed_deposits_3m+ir.banks_fixed_deposits_6m+ir.banks_fixed_deposits_12m+ir.banks_savings_deposits)/4;
}

//A Struct function to find Financial Companies Average Interest Rate for the month
func (ir *InterestRate) AvgFCInterestRate() float64{
	return (ir.fc_fixed_deposits_3m+ir.fc_fixed_deposits_6m+ir.fc_fixed_deposits_12m+ir.fc_savings_deposits)/4;
}

//A Struct function to whether Bank or FC Average Interest Rate is higher for the month (Return Margin higher too)
func (ir *InterestRate) GetHigherInterestRate() (string,float64){
	return "",0;
}