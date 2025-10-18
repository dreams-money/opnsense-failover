package health

import (
	"fmt"
	"net/http"
	"time"
)

var httpClient = http.Client{Timeout: 5 * time.Second}

func CheckPeer(name, address string) error {
	req, err := http.NewRequest("GET", "http://"+address, nil)
	if err != nil {
		return err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%v is down", name)
	}

	return nil
}
