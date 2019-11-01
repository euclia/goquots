package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
	// ClientVersion is used in User-Agent request header to provide server with API level.
	ClientVersion = "0.0.1"

	// httpClientTimeout is used to limit http.Client waiting time.
	httpClientTimeout = 15 * time.Second
)

func main() {}

type QuotsClientProperties struct {
	quotsBase string
	appId     string
	appSecret string
}

var quotsProperties QuotsClientProperties

type QuotsClient struct {
	BaseURL    string
	HttpClient *http.Client
	appId      string
	appSecret  string
}

type Quots struct {
	c QuotsClient
}

func InitQuots(quotsBaseUrl string, appId string, appSecret string) *Quots {
	quotsProperties.quotsBase = quotsBaseUrl
	quotsProperties.appId = appId
	quotsProperties.appSecret = appSecret

	return &Quots{
		c: QuotsClient{
			BaseURL: quotsBaseUrl,
			appId:   appId,
			HttpClient: &http.Client{
				Timeout: 100000000000,
			},
			appSecret: appSecret,
		},
	}
	// return &QuotsClient{
	// 	BaseURL:   quotsBaseUrl,
	// 	appId:     appId,
	// 	appSecret: appSecret,
	// }
}

type IQuots interface {
	// Creates a User on quots with his id, username and email
	CreateUser(userid string, username string, email string) (quotsUser QuotsUser, err error)

	// Finds a User on quots by it's id
	GetUser(userid string) (quotsUser QuotsUser, err error)

	// Checks if a user can proceed on a specific task. User id, usage and the size of usage is needed
	CanUserProceed(id string, usageType string, usageSize string) (canProceed CanProceed, err error)

	// Updates the users credits. The PASSED CREDITS WILL BE THE NEW AVAILABLE CREDITS OF THE USER
	UpdateUserCredits(qu QuotsUser) (quotsUser QuotsUser, err error)

	// Delets a users by it's id
	DeleteUser(id string) (deleted int, err error)
}

func (quo *Quots) CreateUser(userid string, username string, email string) (quotsUser QuotsUser, err error) {
	var createUserUrl = quo.c.BaseURL + "/users"
	var user QuotsUser
	user.Id = userid
	user.Email = email
	user.Username = username
	body, err := json.Marshal(user)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	req, err := http.NewRequest("POST", createUserUrl, bytes.NewBuffer(body))
	req.Header.Set("Authorization", "QUOTSAPP")
	req.Header.Set("app-id", quo.c.appId)
	req.Header.Set("app-secret", quo.c.appSecret)
	resp, err := quo.c.HttpClient.Do(req)
	if err != nil {
		fmt.Printf(err.Error())
		return quotsUser, err
	}
	if resp.StatusCode != 200 {
		var errorReport ErrorReport
		_ = json.NewDecoder(resp.Body).Decode(&errorReport)
		defer resp.Body.Close()
		err = errors.New(errorReport.Message)
		return quotsUser, err
	}
	// if resp.StatusCode > 200{

	// }
	defer resp.Body.Close()
	var userCreated QuotsUser
	err = json.NewDecoder(resp.Body).Decode(&userCreated)
	return userCreated, err
}

func (quo *Quots) GetUser(userid string) (quotsUser QuotsUser, err error) {
	var getUserUrl = quo.c.BaseURL + "/users/" + userid
	req, err := http.NewRequest("GET", getUserUrl, nil)
	req.Header.Set("Authorization", "QUOTSAPP")
	req.Header.Set("app-id", quo.c.appId)
	req.Header.Set("app-secret", quo.c.appSecret)
	resp, err := quo.c.HttpClient.Do(req)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	defer resp.Body.Close()
	var userCreated QuotsUser
	err = json.NewDecoder(resp.Body).Decode(&userCreated)
	return userCreated, err
}

func (quo *Quots) CanUserProceed(id string, usageType string, usageSize string) (canProceed CanProceed, err error) {
	var canProceedUrl = quo.c.BaseURL + "/users/" + id + "/quots"
	req, err := http.NewRequest("GET", canProceedUrl, nil)
	req.Header.Set("Authorization", "QUOTSAPP")
	req.Header.Set("app-id", quo.c.appId)
	req.Header.Set("app-secret", quo.c.appSecret)
	q := req.URL.Query()
	q.Add("appid", quo.c.appId)
	q.Add("usage", usageType)
	q.Add("size", usageSize)
	req.URL.RawQuery = q.Encode()
	resp, err := quo.c.HttpClient.Do(req)
	var canProceedGot CanProceed
	if err != nil {
		fmt.Printf(err.Error())
		return canProceedGot, err
	}
	if resp.StatusCode != 200 {
		var errorReport ErrorReport
		_ = json.NewDecoder(resp.Body).Decode(&errorReport)
		defer resp.Body.Close()
		err = errors.New(errorReport.Message)
		return canProceedGot, err
	}
	defer resp.Body.Close()
	// var canProceedGot CanProceed
	err = json.NewDecoder(resp.Body).Decode(&canProceedGot)
	return canProceedGot, err
}

func (quo *Quots) UpdateUserCredits(qu QuotsUser) (quotsUser QuotsUser, err error) {
	var updateUserCreditsUrl = quo.c.BaseURL + "/users/credits"
	body, err := json.Marshal(qu)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	req, err := http.NewRequest("PUT", updateUserCreditsUrl, bytes.NewBuffer(body))
	req.Header.Set("Authorization", "QUOTSAPP")
	req.Header.Set("app-id", quo.c.appId)
	req.Header.Set("app-secret", quo.c.appSecret)
	resp, err := quo.c.HttpClient.Do(req)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	defer resp.Body.Close()
	var userUpdated QuotsUser
	err = json.NewDecoder(resp.Body).Decode(&userUpdated)
	return userUpdated, err
}

func (quo *Quots) DeleteUser(id string) (deleted int, err error) {
	var deleteUserUrl = quo.c.BaseURL + "/users/" + id
	req, err := http.NewRequest("DELETE", deleteUserUrl, nil)
	req.Header.Set("Authorization", "QUOTSAPP")
	req.Header.Set("app-id", quo.c.appId)
	req.Header.Set("app-secret", quo.c.appSecret)
	resp, err := quo.c.HttpClient.Do(req)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	defer resp.Body.Close()
	var userDeleted int
	err = json.NewDecoder(resp.Body).Decode(&userDeleted)
	return userDeleted, err
}
