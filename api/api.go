package api

import (
	"fmt"
	"github.com/startzerokong/basego/define"
	"github.com/startzerokong/basego/util"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func do(parameter map[string]string, header map[string]string, apiResponse define.ApiResponse) error {
	var err error
	config := util.GetApiConfig()
	fullUrl := config.BaseUri + config.Uri

	urlParameter := &url.Values{}
	for key, value := range parameter {
		urlParameter.Set(key, value)
	}

	for count := 0; count < config.RequestRetryCount; count++ {
		var (
			response  *http.Response
			timeStart = time.Now()
		)

		if header == nil {
			header = make(map[string]string, 0)
		}

		request, _ := http.NewRequest(config.Method, fullUrl, strings.NewReader(urlParameter.Encode()))

		client := &http.Client{Timeout: time.Millisecond * time.Duration(config.RequestTimeout)}

		response, err = client.Do(request)

		apiResponse.HttpCost = (time.Since(timeStart).Nanoseconds()) / 1000 / 1000

		if err == nil {
			for i := 0; i == 0; i++ {
				var byteBody []byte

				byteBody, err = ioutil.ReadAll(response.Body)

				// Build response
				apiResponse.HttpCode = response.StatusCode

				apiResponse.RealResponse = string(byteBody)

				apiResponse.Code = gjson.GetBytes(byteBody, config.CodeKey).String()
				apiResponse.Message = gjson.GetBytes(byteBody, config.MessageKey).String()

				apiResponse.Data = gjson.GetBytes(byteBody, config.DataKey).String()

				apiResponse.ErrorType = ""

				if response.StatusCode != http.StatusOK {
					apiResponse.ErrorType = HttpError
					err = fmt.Errorf(HttpError)
					break
				}

				// when response code is not empty, execute judge, otherwise pass it
				businessError := false

				if config.SuccessCode != nil {
					if apiResponse.Code != strconv.Itoa(*config.SuccessCode) {
						businessError = true
					}
				}

				if businessError {
					apiResponse.ErrorType = BusinessError
					err = fmt.Errorf("%s code [%s] message [%s]", BusinessError, apiResponse.Code, apiResponse.Message)
					break
				}
			}
		}

		if err == nil {
			break
		}

	}

	return err
}