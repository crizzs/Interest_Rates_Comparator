package masapi
/*
This test the process
command functions
*/
import (
	assert "github.com/stretchr/testify/assert"
	"testing"
)

func TestCommandProcessor(t *testing.T) {
	assert.Equal(t, ProcessCommands("1,Jan-2017,Dec-2017") ,"Your Financial Period for Interest Rate Analysis is created." ,"Will return appropriate results")
	assert.Equal(t, ProcessCommands("2") ,"Date | Bank FD 3m | Bank FD 6m | Bank FD 12m | Bank Savings Deposits | FC FD 3m | FC FD 6m | FC FD 12m | FC Savings Deposits |\nJan-2017|0.15 percent|0.2 percent|0.33 percent|0.18 percent|0.3 percent|0.38 percent|0.5 percent|0.17 percent|\nFeb-2017|0.15 percent|0.2 percent|0.33 percent|0.18 percent|0.3 percent|0.38 percent|0.5 percent|0.17 percent|\nMar-2017|0.14 percent|0.2 percent|0.33 percent|0.16 percent|0.3 percent|0.38 percent|0.5 percent|0.17 percent|\nApr-2017|0.14 percent|0.2 percent|0.33 percent|0.16 percent|0.3 percent|0.38 percent|0.5 percent|0.17 percent|\nMay-2017|0.14 percent|0.2 percent|0.33 percent|0.16 percent|0.3 percent|0.38 percent|0.5 percent|0.17 percent|\nJun-2017|0.14 percent|0.2 percent|0.33 percent|0.16 percent|0.3 percent|0.38 percent|0.5 percent|0.17 percent|\nJul-2017|0.14 percent|0.2 percent|0.33 percent|0.16 percent|0.3 percent|0.38 percent|0.5 percent|0.17 percent|\nAug-2017|0.14 percent|0.2 percent|0.33 percent|0.16 percent|0.3 percent|0.38 percent|0.5 percent|0.17 percent|\nSep-2017|0.14 percent|0.2 percent|0.33 percent|0.16 percent|0.3 percent|0.38 percent|0.5 percent|0.17 percent|\nOct-2017|0.14 percent|0.2 percent|0.33 percent|0.16 percent|0.3 percent|0.38 percent|0.5 percent|0.17 percent|\nNov-2017|0.14 percent|0.2 percent|0.33 percent|0.16 percent|0.3 percent|0.38 percent|0.5 percent|0.17 percent|\nDec-2017|0.14 percent|0.2 percent|0.33 percent|0.16 percent|0.3 percent|0.38 percent|0.5 percent|0.17 percent|" ,"Will return appropriate results")
	assert.Equal(t, ProcessCommands("3") ,"Date | Banks Interest Rate (Normalised) | FCs Interest Rate (Normalised) | Overall Rate (Normalised)|\nJan-2017|0.21499999999999997 percent|0.33749999999999997 percent|0.27625 percent|\nFeb-2017|0.21499999999999997 percent|0.33749999999999997 percent|0.27625 percent|\nMar-2017|0.20750000000000002 percent|0.33749999999999997 percent|0.2725 percent|\nApr-2017|0.20750000000000002 percent|0.33749999999999997 percent|0.2725 percent|\nMay-2017|0.20750000000000002 percent|0.33749999999999997 percent|0.2725 percent|\nJun-2017|0.20750000000000002 percent|0.33749999999999997 percent|0.2725 percent|\nJul-2017|0.20750000000000002 percent|0.33749999999999997 percent|0.2725 percent|\nAug-2017|0.20750000000000002 percent|0.33749999999999997 percent|0.2725 percent|\nSep-2017|0.20750000000000002 percent|0.33749999999999997 percent|0.2725 percent|\nOct-2017|0.20750000000000002 percent|0.33749999999999997 percent|0.2725 percent|\nNov-2017|0.20750000000000002 percent|0.33749999999999997 percent|0.2725 percent|\nDec-2017|0.20750000000000002 percent|0.33749999999999997 percent|0.2725 percent|","Will return appropriate results")
	assert.Equal(t, ProcessCommands("4") ,"The overall interest rate average (Financial Period) for Banks Versus Financial Companies is as follow,\nBanks: 0.20875 percent\nFinancial Companies: 0.33749999999999997 percent\n" ,"Will return appropriate results")
	assert.Equal(t, ProcessCommands("5") ,"The interest rates are trending DOWN during the defined financial period." ,"Will return appropriate results")
	assert.Equal(t, ProcessCommands("@#$"), "Your Input is not part of the selection. Please try again!" ,"Will show invalid message!")
}