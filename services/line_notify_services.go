package services

import (
	"errors"
	"io/ioutil"
	"strings"

	"net/http"
	"net/url"
)

type LineNotifyConfig struct {
	Bearer   string
	Messages string
}

func (LNC *LineNotifyConfig) LineNotify(LineNotifyConfig) error {
	const token = "s5ZFwpC3BPdrTMF9VtM7qCYDvmerUCGznXOrBiEaScA"
	if LNC.Bearer == "" {
		LNC.Bearer = token
	}
	if LNC.Messages == "" {
		return	errors.New("messages empty.")
	}
	URL := "https://notify-api.line.me/api/notify"

	client := &http.Client{}

data:=url.Values{"message":{LNC.Messages}}
	req, err := http.NewRequest(http.MethodPost, URL, strings.NewReader(data.Encode()))
	if err != nil {
		return errors.New(err.Error())
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+LNC.Bearer)

	resp, err := client.Do(req)
	if err != nil {
		return errors.New(err.Error())
	}

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {

		return errors.New("errror from line notify:" + string(body))
	}
	return nil
}
