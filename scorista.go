package scorista

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"time"

	"fmt"

	"github.com/pkg/errors"
)

type Status string

const (
	ST_OK    Status = "OK"
	ST_ERROR Status = "ERROR"
	ST_WAIT  Status = "WAIT"
	ST_DONE  Status = "DONE"
)

const SCORISTA_END_POINT = "https://api.scorista.ru/mixed/json"

type S map[string]interface{}

type Scorista struct {
	username   string
	secretKey  string
	httpClient *http.Client
}

type Error struct {
	Code    uint            `json:"code"`
	Message string          `json:"message"`
	Details json.RawMessage `json:"details"`
}

type CreditExamResponse struct {
	Status    Status `json:"status"`
	RequestID string `json:"requestid"`
	Error     Error  `json:"error"`
}

type CreditDecisionResponse struct {
	Status Status          `json:"status"`
	Data   json.RawMessage `json:"data"`
	Error  Error           `json:"error"`
}

func New(username, secretKey string) *Scorista {
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 10 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
	}

	httpClient := &http.Client{Transport: transport}

	return &Scorista{
		username: username, secretKey: secretKey, httpClient: httpClient,
	}
}

func (s *Scorista) CreditDecision(requestID string) (CreditDecisionResponse, error) {
	result := CreditDecisionResponse{}

	dataByte, _ := json.Marshal(map[string]interface{}{
		"requestID": requestID,
	})

	req, err := s.makeRequest(bytes.NewBuffer(dataByte))
	if err != nil {
		return result, err
	}

	responseDataByte, err := s.sendRequest(req)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(responseDataByte, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *Scorista) CreditExam(data S) (CreditExamResponse, error) {
	result := CreditExamResponse{}

	dataByte, err := json.Marshal(data)
	if err != nil {
		return result, err
	}

	req, err := s.makeRequest(bytes.NewBuffer(dataByte))
	if err != nil {
		return result, err
	}

	responseDataByte, err := s.sendRequest(req)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(responseDataByte, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *Scorista) sendRequest(req *http.Request) ([]byte, error) {
	response, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// create new request
func (s *Scorista) makeRequest(body *bytes.Buffer) (*http.Request, error) {
	req, err := http.NewRequest(
		"POST", SCORISTA_END_POINT, body,
	)
    req.Close = true

	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("SOAPAction", "urn:InnControllerwsdl#request")

	s.setAuthHeaders(req)

	return req, nil
}

// set auth headers
func (s *Scorista) setAuthHeaders(req *http.Request) {
	nonce := s.genNonce()
	req.Header.Set("username", s.username)
	req.Header.Set("nonce", nonce)
	req.Header.Set("password", s.makePassword(nonce))
}

// make password
func (s *Scorista) makePassword(nonce string) string {
	return s.encodeStrToSha(nonce + s.secretKey)
}

// generate nonce string
func (s *Scorista) genNonce() string {
	return s.encodeStrToSha(strconv.Itoa(time.Now().Nanosecond()))
}

// encode string to SHA1
func (s *Scorista) encodeStrToSha(str string) string {
	hash := sha1.New()
	hash.Write([]byte(str))

	return hex.EncodeToString(hash.Sum(nil))
}
