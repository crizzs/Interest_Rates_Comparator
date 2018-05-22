package masapi
/***************
Test Business
Logic Layer
for Interest
Rate Module
****************/
import (
	assert "github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateFinancialPeriod(t *testing.T) {
	status := CreateFinancialPeriod("Jan-2017","Dec-2017")
	assert.Equal(t, status,"" ,"It should acknowledge this program have created financial period.")
}
