package masapi
/*
This function process the commands
from users
*/
import(
	"os"
	"fmt"
	"bufio"
	"strings"
)
//This is the terminal to read the in-line inputs
func ActivateInteractiveCommandTerminal(){
	
	commandReader := bufio.NewScanner(os.Stdin)
	//trackLineCounter := 1
	collectedBothDates := false
	startDate := ""
	endDate := ""
	//First line will always look for start date
	fmt.Print("\nWelcome to Interest Rates Explorer!\n\nEnter Start Date(Eg. Jan-2017) :")
	for commandReader.Scan() {
		var commandRetrieved = commandReader.Text()
		
		if collectedBothDates != false && commandRetrieved != "1" && commandRetrieved != "6" && commandRetrieved != "2" && commandRetrieved != "3" && commandRetrieved != "4" && commandRetrieved != "5"{
			fmt.Print("\nBelow are the list of selections available (1 to 6)\n1. Change the Financial Period (Both Dates)\n2. Display Interest Rates for Financial Period\n3. Display months which Finance Companies offered best Interest Rates\n4. Show Averaged Interest Rates Across the entire Financial Period\n5. Tell Trend of Interest Rates During this period\n6. Exit the Application")
			fmt.Print("\n\nEnter (1 to 6):")
		}

		if commandRetrieved == "6"{
			break
		}else if commandRetrieved == "1"{
			//Resets the financial period
			collectedBothDates = false
			startDate = ""
			endDate = ""
			
		}else if commandRetrieved == "2" || commandRetrieved == "3" || commandRetrieved == "4" || commandRetrieved == "5"{
			fmt.Print("\n\nResult :\n\n"+ProcessCommands(commandRetrieved)+"\n")
			fmt.Print("\nBelow are the list of selections available (1 to 6)\n1. Change the Financial Period (Both Dates)\n2. Display Interest Rates for Financial Period\n3. Display months which Finance Companies offered best Interest Rates\n4. Show Averaged Interest Rates Across the entire Financial Period\n5. Tell Trend of Interest Rates During this period\n6. Exit the Application")
			fmt.Print("\n\nEnter (1 to 6):")
		}

		if collectedBothDates == false && startDate == "" && commandRetrieved != "1"{
			
			startDate = commandRetrieved
			fmt.Print("\nEnter End Date(Eg. Dec-2017) :")

		}else if collectedBothDates == false && startDate == "" && commandRetrieved == "1"{
			
			fmt.Print("\n\nEnter Start Date(Eg. Jan-2017) :")

		}else if collectedBothDates == false && startDate!="" && endDate == ""{
			
			endDate = commandRetrieved
			
			//Check it is correct 
			result := ProcessCommands("1,"+startDate+","+endDate)
			if result != "Your Financial Period for Interest Rate Analysis is created."{
				fmt.Print("\n"+result+"\n\nEnter Start Date(Eg. Jan-2017) :")
				startDate = ""
				endDate = ""
			}else{
				collectedBothDates = true
				startDate = ""
				endDate = ""
				fmt.Print("\n"+result+"\n")
				fmt.Print("\nBelow are the list of selections available (1 to 6)\n1. Change the Financial Period (Both Dates)\n2. Display Interest Rates for Financial Period\n3. Display months which Finance Companies offered best Interest Rates\n4. Show Averaged Interest Rates Across the entire Financial Period\n5. Tell Trend of Interest Rates During this period\n6. Exit the Application\n\n")
				fmt.Print("\n\nEnter (1 to 6):")
			}
		}
	}
	fmt.Println("\n\nBye Bye!")
}
//Process the command and distribute to correct function inside controller.
func ProcessCommands(str string) string{
	var splitCommand = strings.Split(str,",")

	if len(splitCommand) == 3 && splitCommand[0] == "1"{
		return CreateFinancialPeriod(splitCommand[1],splitCommand[2])
	}else if len(splitCommand)==1 && (splitCommand[0] == "2" || splitCommand[0] == "3" || splitCommand[0] == "4" || splitCommand[0] == "5"){
		if splitCommand[0] == "2"{
			return VisualizeIRComparisonByMonth()
		}else if splitCommand[0] == "3"{
			return VisualizeMonthsThatFCsWin()
		}else if splitCommand[0] == "4"{
			return ShowOverallBanksVersusFCsAvg()
		}else if splitCommand[0] == "5"{
			return ShowTrend()
		}
	}
	return GetRepliesText(12)
}