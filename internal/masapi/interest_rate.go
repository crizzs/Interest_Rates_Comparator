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

//A Struct Function to find Average Interest Rate across both banks and financial companies
func (ir *InterestRate) AvgInterestRate() float64{
	return (ir.banks_fixed_deposits_3m+ir.banks_fixed_deposits_6m+ir.banks_fixed_deposits_12m+ir.banks_savings_deposits+ir.fc_fixed_deposits_3m+ir.fc_fixed_deposits_6m+ir.fc_fixed_deposits_12m+ir.fc_savings_deposits)/8;
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
	bankInterestRate := ir.AvgBankInterestRate()
	fcInterestRate := ir.AvgFCInterestRate()

	if bankInterestRate > fcInterestRate {
		return "bank",(bankInterestRate-fcInterestRate)
	}else if fcInterestRate > bankInterestRate {
		return "fc",(fcInterestRate-bankInterestRate)	
	}
	//This will only return if the interest is the same
	return "same",0;
}
//Check if this result is a valid piece of Interest Rate obj
func (ir *InterestRate) CheckInterestRateValidity() bool{
	if ir.banks_fixed_deposits_3m== -9999.99 || ir.banks_fixed_deposits_6m== -9999.99 || ir.banks_fixed_deposits_12m== -9999.99 || ir.banks_savings_deposits== -9999.99 || ir.fc_savings_deposits== -9999.99 || ir.fc_fixed_deposits_3m== -9999.99 || ir.fc_fixed_deposits_6m == -9999.99|| ir.fc_fixed_deposits_12m== -9999.99 {
		return false;
	}
	return true;
}

//Generates a string for display (Avg Bank/FCs/Overall)
func (ir *InterestRate) GetDisplay() string{
	if ir.CheckInterestRateValidity() == false{
		return ConvertResultDateStrForDisplay(ir.end_of_month)+"| Not Available | Not Available | Not Available|"
	}

	avgIR := ir.AvgInterestRate()
	avgBankIR := ir.AvgBankInterestRate()
	avgFCIR := ir.AvgFCInterestRate()
	return ConvertResultDateStrForDisplay(ir.end_of_month)+"|" + FloatToStr(avgBankIR) +" percent|" + FloatToStr(avgFCIR) + " percent|" + FloatToStr(avgIR) +" percent|"
}

//Generates a string for all display (General)
func (ir *InterestRate) GetAllDisplay() string{
	if ir.CheckInterestRateValidity() == false{
		return ConvertResultDateStrForDisplay(ir.end_of_month)+"| Not Available | Not Available | Not Available| Not Available | Not Available | Not Available | Not Available | Not Available |"
	}

	return ConvertResultDateStrForDisplay(ir.end_of_month)+"|" + FloatToStr(ir.banks_fixed_deposits_3m) +" percent|" + FloatToStr(ir.banks_fixed_deposits_6m) + " percent|" + FloatToStr(ir.banks_fixed_deposits_12m) +" percent|" +FloatToStr(ir.fc_fixed_deposits_3m)+" percent|"+FloatToStr(ir.fc_fixed_deposits_6m)+" percent|"+FloatToStr(ir.fc_fixed_deposits_12m)+" percent|"+FloatToStr(ir.fc_savings_deposits)+" percent|"
}


