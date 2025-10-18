package notifications

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	PUSHOVER_URL = "https://api.pushover.net/1/messages.json"
	USER_TOKEN   = "aorhv9q8tcccqo22a6ukzpd14cn3jj"
	API_TOKEN    = "uto89kzozmnpkvg1hzf4boacyym35f"
	DEVICE       = "iphone_14-cesar"
	TITLE        = "Postgres OPNsense Failover"
)

var httpClient = http.Client{Timeout: 5 * time.Second}

func PushMessage(message string) error {
	postBody := "token=%v&user=%v&device=%v&title=%v&message=%v"
	postBody = fmt.Sprintf(postBody, USER_TOKEN, API_TOKEN, DEVICE, TITLE, message)

	body := bytes.NewBuffer([]byte(postBody))

	req, err := http.NewRequest("POST", PUSHOVER_URL, body)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	responseBody := string(bodyBytes)

	if resp.StatusCode != 200 {
		e := "pushover request failed with status: %v, msg: %v"
		return fmt.Errorf(e, resp.StatusCode, responseBody)
	}

	// {"status":1,"request":"97b8a0ba-4113-41d1-a01a-8c64c0ff3c06"}
	type pushoverResponse struct {
		Status  int    `json:"status"`
		Request string `json:"request"`
	}

	jr := pushoverResponse{}
	err = json.Unmarshal(bodyBytes, &jr)
	if err != nil {
		return err
	}

	if jr.Status != 1 {
		e := "pushover response status: %v, msg: %v"
		return fmt.Errorf(e, jr.Status, responseBody)
	}

	return nil
}
