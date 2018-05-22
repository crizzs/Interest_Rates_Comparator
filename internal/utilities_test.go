package masapi
/***************
Test Utility
Functions
****************/
import (
	assert "github.com/stretchr/testify/assert"
	"testing"
)

func TestStrToFloatUtil(t *testing.T) {
	assert.Equal(t, StrToFloat("null"),-9999.99 ,"If null attribute is being received. Return -9999.99")
	
	assert.Equal(t, StrToFloat("9.63"),9.63 ,"This should return a float number")
}

func TestStrToIntUtil(t *testing.T) {
	assert.Equal(t, StrToInt("0"),0 ,"If 0 is being received. Return 0")
	
	assert.Equal(t, StrToInt("9"),9 ,"This should be 9")

	assert.Equal(t, StrToInt("-1"),0 ,"This should be 0")
}

func TestFloatToStr(t *testing.T) {
	assert.Equal(t, FloatToStr(0.356), "0.356" ,"Converted into a string.")
	//This is an invalid(Null attribute generated from MAS)
	assert.Equal(t, FloatToStr(-9999.99), "-9999.99" ,"Converted into a string.")
}

func TestStrToDateStrUtil(t *testing.T) {
	//Provides Test Values to convert String to Date
	_,validDate := ConvertStrToDate("fgdgddfg")

	assert.Equal(t, validDate, false, "This is an invalid date.")

	dateReturn,validDateTest := ConvertStrToDate("MAR-2018")

	assert.Equal(t, validDateTest, true, "This is an valid date.")

	assert.Equal(t, dateReturn, "2018-03", "This is an valid formatted date string.")
}

func TestResultDateStrToDisplayFormat(t *testing.T){
	otherFormat := ConvertResultDateStrForDisplay("1983-02")
	//This is a valid display format
	assert.Equal(t, otherFormat, "Feb-1983", "This is a correct display format.")

}

func TestToDateIsNotBeforeFromDate(t *testing.T) {
	
	assert.Equal(t, TestFromAndToDateValidity("MAR-2018","Jan-2017"), false, "This is an invalid from and to date.")

	assert.Equal(t, TestFromAndToDateValidity("Dec-2016","Jan-2017"), true, "This is a valid from and to date.")
}

func TestReplyTextMethod(t *testing.T) {
	assert.Equal(t, GetRepliesText(1), "Both of your date values are incorrect.\nThis is an example of a correct input (Jan-2017)", "This is a text reply")
}