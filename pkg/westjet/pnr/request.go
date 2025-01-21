package pnr

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	reqEndpoint = `https://apiw.westjet.com/triplookup/itinerary?pnr={conf}&lastname={lname}&eligibility=true&includeStandby=true`

	userAgent = `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36`
)

var (
	requestHeaders = map[string]string{
		"Accept":          "application/json",
		"Accept-Language": "en-US",
		"Content-Type":    "application/json",
		"DNT":             "1",
		"Origin":          "https://www.westjet.com",
		"Referer":         "https://www.westjet.com/",
		"User-Agent":      userAgent,
	}

	client = http.Client{
		Timeout: time.Second * 15,

		Transport: &http.Transport{
			Proxy: func(r *http.Request) (*url.URL, error) {
				if p := os.Getenv("HTTP_PROXY"); p != "" {
					return url.Parse(p)
				}

				return nil, nil
			},
		},
	}
)

func generateReqUrl(lastName, confirmationCode string) string {
	url := reqEndpoint
	url = strings.Replace(url, "{conf}", confirmationCode, -1)
	url = strings.Replace(url, "{lname}", lastName, -1)
	return url
}

func setRequestHeaders(r *http.Request) {
	for k, v := range requestHeaders {
		r.Header.Set(k, v)
	}
}

func sendRequest(lastName, confirmationCode string) ([]byte, error) {
	req, err := http.NewRequest("GET", generateReqUrl(lastName, confirmationCode), nil)

	if err != nil {
		return []byte{}, err
	}

	setRequestHeaders(req)

	res, err := client.Do(req)

	if err != nil || res == nil || res.StatusCode != 200 {
		if res != nil {
			log.Println("status code from westjet:", res.StatusCode)
		}

		if err != nil {
			log.Println("error from westjet:", err)
		}

		return []byte{}, errors.New("status code was not 200")
	}

	defer res.Body.Close()
	return io.ReadAll(res.Body)
}

func performRequest(lastName string, confirmationCode string) (res GetPNRResponse, err error) {
	data, err := sendRequest(lastName, confirmationCode)

	if err != nil {
		return res, err
	}

	if err := json.Unmarshal(data, &res); err != nil {
		return res, err
	}

	return res, nil
}

func convertResponse(res GetPNRResponse) (pnr PNR) {
	convertPnrResponse(res, &pnr)
	convertFlights(res, &pnr)

	return pnr
}
