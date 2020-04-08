package tests

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cucumber/godog"
)

var response *http.Response

func iSendRequestTo(httpMethod, url string) (err error) {
	switch httpMethod {
	case http.MethodGet:
		response, err = http.Get(url)
	default:
		err = fmt.Errorf("unknown method: %s", httpMethod)
	}
	if err != nil {
		return err
	}
	return err
}

func theResponseCodeShouldBe(code int) (err error) {
	if response.StatusCode != code {
		err = fmt.Errorf("http server return wrong code: %d", code)
	}
	return err
}

func theResponseShouldMatchText(text string) (err error) {
	responseBodyText, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if string(responseBodyText) != text {
		err = fmt.Errorf("http server return wrong responce: %s", responseBodyText)
	}
	return err
}

func FeatureWebContext(s *godog.Suite) {
	s.Step(`^I send "([^"]*)" request to "([^"]*)"$`, iSendRequestTo)
	s.Step(`^The response code should be (\d+)$`, theResponseCodeShouldBe)
	s.Step(`^The response should match text "([^"]*)"$`, theResponseShouldMatchText)
}
