package masapi
/***************
Test Financial
Period
****************/
import (
	assert "github.com/stretchr/testify/assert"
	"testing"
)

func TestFinancialPeriod(t *testing.T) {

	sampleFinancialPeriod := InitiatizeFinancialPeriod("2017-01","2018-01")

	//Test if the amount of results is the same as requested
	assert.Equal(t, len(sampleFinancialPeriod.interestRateArr), 12)
}
/*
func TestConvertStrToDateFunc(t *testing.T) {
	sampleARubbishInput := ConvertStrToTime("fgdgddfg")
	assert.Equal(t, sampleARubbishInput, 12)

}
*/