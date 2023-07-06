package backendconnection

import (
	"MessengerClient/internal/utils"
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func connectChannelRequest(args []string, serverIp string) utils.ConnectJson {
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
	response, postErr := client.Post(("https://" + serverIp + "/connectChannel"), "text/plain", dataReader)
	if postErr != nil {
		log.Printf("Post connection data about channel error: %v", postErr)
	}
	respBody, respReadErr := ioutil.ReadAll(response.Body)
	if respReadErr != nil {
		log.Printf("Read channel connection status error: %v", respReadErr)
	}
	return utils.ReadConnectJSON(respBody)
}
