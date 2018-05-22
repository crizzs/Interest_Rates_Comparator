package masapi

import(
	"strconv"
	"time"
)
/*
This contains utilites of
masapi package
*/

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
	parsedDate ,_ := time.Parse("2006-02",dateStr)
	formattedDate := parsedDate.Format("Jan-2006")

	return formattedDate
}

//Test Date Inputs from users are correct
func TestFromAndToDateValidity(fromDateStr string, toDateStr string) bool{
	fromDate ,_ := time.Parse("Jan-2006",fromDateStr)
	toDate ,_ := time.Parse("Jan-2006",toDateStr)

	return fromDate.Before(toDate)
}