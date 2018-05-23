package masapi

import(
	"strconv"
	"time"
)
/*
This contains utilites of
main package
*/
var listOfReplies = map[int]string{
									1: "Both of your date values are incorrect.\nThis is an example of a correct input (Jan-2017)",
									2: "Your Start Date is incorrect.\nThis is an example of a correct input (Jan-2017)",
									3: "Your End Date is incorrect.\nThis is an example of a correct input (Jan-2017)",
									4: "Your End Date is before Start Date.",
									5: "Your Financial Period for Interest Rate Analysis is created.",
									6: "There is no interest rates data found!",
									7: "There is no data on Banks and Financial Companies Interest Rate (Overall Average).",
									8: "Not able to tell the trending of interest rates during this period.",
									9: "The interest rates are trending UP during the defined financial period.",
									10: "The interest rates are trending DOWN during the defined financial period.",
									11: "The interest rates are holding STEADY during the defined financial period.",
									12: "Your Input is not part of the selection. Please try again!"}

func StrToFloat(str string) float64{
	s, err := strconv.ParseFloat(str, 64)
	//Return a numerical impossible number for null attribute
	if err != nil{
		return -9999.99;
	}
	return s
}

func StrToInt(str string) int{
	s, err := strconv.Atoi(str)

	if err != nil || s < 0{
		return 0;
	}
	return s
}

func FloatToStr(f float64) string{
	s := strconv.FormatFloat(f,'f', -1, 64)

	return s
}

func ConvertStrToDate(dateStr string) (string,bool){
	date ,err := time.Parse("Jan-2006",dateStr)
	
	if err != nil {
		return "",false
	}
	formattedDate := date.Format("2006-01")

	return formattedDate,true
}

//This is a method used for date display formatting
func ConvertResultDateStrForDisplay(dateStr string) string{
	parsedDate ,_ := time.Parse("2006-01",dateStr)
	formattedDate := parsedDate.Format("Jan-2006")

	return formattedDate
}

//Test Date Inputs from users are correct
func TestFromAndToDateValidity(fromDateStr string, toDateStr string) bool{
	fromDate ,_ := time.Parse("Jan-2006",fromDateStr)
	toDate ,_ := time.Parse("Jan-2006",toDateStr)

	return fromDate.Before(toDate)
}

//Replies Utility Function
func GetRepliesText(txtNum int) string{
	return listOfReplies[txtNum];
}