package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestCreateUser(t *testing.T) {
	q := InitQuots("http://localhost:8004", "GOQUOTS", "IlFELGMLf^BmJg2MVV")
	var IQuots = q
	quotsUser, err := IQuots.CreateUser("useridquoquots", "goquotsuser", "goquotsusermail")
	if err != nil {
		// t.Log("Passing test with error:" + err.Error())
		fmt.Println("Passing test with error:" + err.Error())
	}
	if quotsUser.Id != "useridquoquots" {
		t.Error("Not the same user id")
	}
}

func TestGetUser(t *testing.T) {
	q := InitQuots("http://localhost:8004", "GOQUOTS", "IlFELGMLf^BmJg2MVV")
	var IQuots = q
	quotsUser, err := IQuots.GetUser("useridquoquots")
	if err != nil {
		// t.Log("Passing test with error:" + err.Error())
		fmt.Println("Passing test with error:" + err.Error())
	}
	if quotsUser.Id != "useridquoquots" {
		t.Error("Not the same user id")
	}
}

func TestCanUserProceed(t *testing.T) {
	q := InitQuots("http://localhost:8004", "GOQUOTS", "IlFELGMLf^BmJg2MVV")
	var IQuots = q
	canUserProceed, err := IQuots.CanUserProceed("useridquoquots", "TASK", "1")
	if err != nil {
		// t.Log("Passing test with error:" + err.Error())
		fmt.Println("Passing test with error:" + err.Error())
	}
	if canUserProceed.UserId != "useridquoquots" {
		t.Error("Not the same user id")
	}
}

func TestUpdateUserCredits(t *testing.T) {
	q := InitQuots("http://localhost:8000", "GOQUOTS", "IlFELGMLf^BmJg2MVV")
	var IQuots = q
	var qu QuotsUser
	qu.Id = "useridquoquots"
	qu.Email = "goquotsusermail"
	qu.Username = "goquotsusername"
	qu.Credits = 64.4
	quGot, err := IQuots.UpdateUserCredits(qu)
	if err != nil {
		// t.Log("Passing test with error:" + err.Error())
		fmt.Println("Passing test with error:" + err.Error())
	}
	if quGot.Credits != 64.4 {
		f := strconv.FormatFloat(quGot.Credits, 'f', 6, 64)
		t.Error("Not the same user credits" + f)
	}
}

func TestDeleteUser(t *testing.T) {
	q := InitQuots("http://localhost:8004", "GOQUOTS", "IlFELGMLf^BmJg2MVV")
	var IQuots = q
	quGot, err := IQuots.DeleteUser("useridquoquots")
	if err != nil {
		// t.Log("Passing test with error:" + err.Error())
		fmt.Println("Passing test with error:" + err.Error())
	}
	if quGot != 1 {
		t.Error("Not 1")
	}
}
