package backendconnection

import (
	"MessengerClient/internal/utils"
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var CreateChannelFunc func([]string, string, func([]string, string)) = func(args []string, serverIp string, connectFunc func(args []string, serverIp string)) {

	dataReader := bytes.NewReader([]byte(strings.Trim(utils.DeleteEscape(args[0]), "\n")))
	caCertPool := x509.NewCertPool()
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:            caCertPool,
				InsecureSkipVerify: true,
			},
		},
	}
	response, postErr := client.Post(("https://" + serverIp + "/createChan"), "text/plain", dataReader)
	if postErr != nil {
		log.Printf("Post creation data about channel error: %v", postErr)
	}
	respBody, respReadErr := ioutil.ReadAll(response.Body)
	if respReadErr != nil {
		log.Printf("Read channel creating status error: %v", respReadErr)
	}
	switch string(respBody) {
	case "channel_exist":
		fmt.Println("Channel already exist")
	case "channel_not_exist":
		fmt.Println("Channel created succesfully")
		connectFunc(args, serverIp)
	}
}
