package userInterface

import (
	"MessengerClient/internal/utils"
	"bufio"
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var (
	responseBody    []byte
	responseReadErr error
	Work            bool
)

var login, password string

func CLI(createFunc func(args []string, serverIp string, connectFunc func(args []string, serverIp string)),
	connectFunc func(args []string, serverIp string),
	disconnectFunc func(),
	listFunc func(),
	sendMessegeFunc func(msg string),
	deleteMessegeFunc func(id string),
	serverIp string) {
	reCreate := regexp.MustCompile("/create")
	reConnect := regexp.MustCompile("/connect")
	reDisconnect := regexp.MustCompile("/disconnect")
	reList := regexp.MustCompile("/list")
	reDeleteMessege := regexp.MustCompile("/delete")
	reQuit := regexp.MustCompile("/quit")
	reInvalid := regexp.MustCompile("/")
	var input string
	var in *bufio.Reader
	var readErr error
	var createBool, connectBool, disconnectBool, listBool, invalidBool, deleteMessegeBool, quitBool bool
	var createArgs, connectArgs, deleteArgs []string
	for Work {
		in = bufio.NewReader(os.Stdin)
		input, readErr = in.ReadString('\n')
		if readErr != nil {
			fmt.Printf("Read error: %v", readErr)
		}
		createBool = false
		connectBool = false
		listBool = false
		disconnectBool = false
		invalidBool = false
		deleteMessegeBool = false
		quitBool = false
		createArgs = nil
		connectArgs = nil
		deleteArgs = nil
		if reCreate.FindStringIndex(string(input)) != nil {
			createBool = (reCreate.FindStringIndex(input)[0] == 0) && (reCreate.FindStringIndex(input)[1] == 7)
		} else if reConnect.FindStringIndex(input) != nil {
			connectBool = (reConnect.FindStringIndex(input)[0] == 0) && (reConnect.FindStringIndex(input)[1] == 8)
		} else if reDisconnect.FindStringIndex(input) != nil {
			disconnectBool = (reDisconnect.FindStringIndex(input)[0] == 0) && (reDisconnect.FindStringIndex(input)[1] == 11)
		} else if reList.FindStringIndex(input) != nil {
			listBool = (reList.FindStringIndex(input)[0] == 0) && (reList.FindStringIndex(input)[1] == 5)
		} else if reDeleteMessege.FindStringIndex(input) != nil {
			deleteMessegeBool = (reDeleteMessege.FindStringIndex(input)[0] == 0) && (reDeleteMessege.FindStringIndex(input)[1] == 7)
		} else if reQuit.FindStringIndex(input) != nil {
			quitBool = (reQuit.FindStringIndex(input)[0] == 0) && (reQuit.FindStringIndex(input)[1] == 5)
		} else if reInvalid.FindStringIndex(input) != nil {
			invalidBool = (reInvalid.FindStringIndex(input)[0] == 0) && (reInvalid.FindStringIndex(input)[1] == 1)
		}
		if createBool {
			createArgs = utils.FindArguments(reCreate.FindStringIndex(input)[1], input)
			if strings.Trim(strings.Trim(createArgs[0], "\n"), " ") != "" {
				createFunc(createArgs, serverIp, connectFunc)
			} else {
				fmt.Println("Empty arguments")
			}
		} else if connectBool {
			connectArgs = utils.FindArguments(reConnect.FindStringIndex(input)[1], input)
			if strings.Trim(strings.Trim(connectArgs[0], "\n"), " ") != "" {
				connectFunc(connectArgs, serverIp)
			} else {
				fmt.Println("Empty arguments")
			}
		} else if disconnectBool {
			disconnectFunc()
		} else if listBool {
			listFunc()
		} else if deleteMessegeBool {
			deleteArgs = utils.FindArguments(reDeleteMessege.FindStringIndex(input)[1], input)
			if strings.Trim(strings.Trim(deleteArgs[0], "\n"), " ") != "" {
				deleteMessegeFunc(deleteArgs[0])
			} else {
				fmt.Println("Empty arguments")
			}
		} else if quitBool {
			os.Exit(1)
		} else if invalidBool {
			println("Invalid command. Please try again.")
		} else {
			if strings.Trim(strings.Trim(input, "\n"), " ") != "" {
				sendMessegeFunc(input)
			} else {
				fmt.Println("Empty messege")
			}
		}
	}
}
func Auth(serverIp string) {
	fmt.Print("Enter login: ")
	fmt.Fscanln(os.Stdin, &login)
	fmt.Print("Enter password: ")
	fmt.Fscanln(os.Stdin, &password)
	loginData, dataErr := utils.LoginMessage(login, password)
	if dataErr != nil {
		log.Fatalf("Create post login data error: %v", dataErr)
	}
	data := bytes.NewReader(loginData)
	caCertPool := x509.NewCertPool()
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:            caCertPool,
				InsecureSkipVerify: true,
			},
		},
	}
	response, postErr := client.Post(("https://" + serverIp + "/authData"), "application/json", data)
	if postErr != nil {
		log.Fatalf("Post error: %v", postErr)
	}
	responseBody, responseReadErr = ioutil.ReadAll(response.Body)
	if responseReadErr != nil {
		log.Fatalf("Response read error: %v", responseReadErr)
	}
	for (string(responseBody) == "user_not_exist") || (string(responseBody) == "invalid_password") {
		fmt.Printf("Auth error: %v\n", string(responseBody))
		Auth(serverIp)
	}
	utils.CurrentSessionData = utils.SessionData{
		CurrentLogin: login,
		JWT:          string(responseBody),
	}
	fmt.Println("Succesful log in")
	Work = true
}
