package main

import (
	"encoding/base64"
	"github.com/l-vitaly/scorista"
	"io/ioutil"
	"log"
)

func main() {
	s := scorista.New("tatiana@devim.team", "88e8c0ee532f3754fa7dbb96667503a81031fd16")

	form := getForm()
	resp, err := s.CreditExam(form)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(resp.Error.Details), resp.Status == scorista.ST_OK, resp.RequestID)
}

func getForm() scorista.S {
	creditHistoryRaw, err := ioutil.ReadFile("./credit_history.xml")
	if err != nil {
		log.Fatalln(err)
	}

	creditHistory := base64.StdEncoding.EncodeToString([]byte(creditHistoryRaw))

	return scorista.S{
		"form": scorista.S{
			"persona": scorista.S{
				"personalInfo": scorista.S{
					"personaID":      "12345",
					"lastName":       "Иванова",
					"firstName":      "Нина",
					"patronimic":     "Викторовна",
					"gender":         2,
					"birthDate":      "26.02.1900",
					"placeOfBirth":   "пос.Ивановка",
					"passportSN":     "1234 000000",
					"issueDate":      "12.03.1900",
					"subCode":        "000-125",
					"issueAuthority": "отделом УФМС России по гор. Москве",
					"maritalStatus":  1,
				},
				"addressRegistration": scorista.S{
					"postIndex": "000000",
					"region":    "Москва город",
					"city":      "Москва город",
					"street":    "Литовский бульвар",
					"house":     "1",
					"building":  "5",
					"flat":      "1",
				},
				"addressResidential": scorista.S{
					"postIndex": "000000",
					"region":    "Москва город",
					"city":      "Москва город",
					"street":    "Литовский бульвар",
					"house":     "1",
					"building":  "5",
					"flat":      "1",
				},
				"contactInfo": scorista.S{
					"cellular":            "89090000000",
					"cellularState":       1,
					"cellularMethod":      8,
					"phone":               "нет",
					"phoneState":          1,
					"phoneMethod":         4,
					"email":               "нет",
					"emailState":          1,
					"emailMethod":         4,
					"relativePhone":       "8962000000",
					"relativeLastName":    "Иванова",
					"relativeFirstName":   "Татьяна",
					"relativePatronimic":  "Николаевна",
					"relativePhoneState":  1,
					"relativePhoneMethod": 4,
					"spousePhone":         "нет",
					"spouseLastName":      "",
					"spouseFirstName":     "",
					"spousePatronimic":    "",
					"spousePhoneState":    "",
					"spousePhoneMethod":   "",
				},
				"employment": scorista.S{
					"jobCategory":        8,
					"employer":           "ГУП Московский метрополитен",
					"employerSite":       "",
					"employerPhone":      "",
					"workPhone":          "нет",
					"workPhoneState":     1,
					"workPhoneMethod":    4,
					"workEmail":          "нет",
					"workEmailState":     1,
					"workEmailMethod":    4,
					"salaryOfficial":     40000.00,
					"salaryActual":       40000.00,
					"occupation":         "маляр",
					"employmentType":     1,
					"employmentTime":     168,
					"jobExpirience":      "",
					"previousEmployment": "",
				},
			},
			"info": scorista.S{
				"loan": scorista.S{
					"loanID":                    "777",
					"staffMember":               "Иванова Мария Григорьевна",
					"loanPeriod":                28,
					"loanSum":                   10000,
					"dayRate":                   2,
					"loanCurrency":              "RUB",
					"fullRepaymentAmount":       15600,
					"day30DelayRepaymentAmount": 23600,
					"applicationSourceType":     2,
					"applicationSourceMethod":   1,
					"agreementSignatureMethod":  1,
					"loanReceivingMethod":       5,
					"loanRepaymentMethod":       5,
				},
				"repaymentSchedule": scorista.S{
					"repaymentDate":   "10.12.2016",
					"repaymentAmount": 15600,
				},
				"borrowingHistory": scorista.S{
					"numberLoansRepaid":              0,
					"previousLoanDate":               "",
					"previousLoanPlanRepaidDate":     "",
					"previousLoanFactRepaidDate":     "",
					"previousLoanAmount":             0,
					"previousLoanRepaymentAmount":    0,
					"previousLoanReceivingMethod":    5,
					"previousLoanRepaymentMethod":    5,
					"previousLoanProlongationNumber": 4,
					"softCollectionFlag":             0,
				},
			},
			"loanReceivingMethod": scorista.S{
				"cash": scorista.S{
					"cash": 1,
				},
			},
			"NBKI": creditHistory,
		},
	}
}
