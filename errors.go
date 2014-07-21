package goexile

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "net/url"
)

type ExileErrorResponse struct {
    SingleError ExileError `json:"error"`
}

type ExileError struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}

type ApiError struct {
    StatusCode int
    Header     http.Header
    Body       string
    Data       ExileErrorResponse
    URL        *url.URL
}

func newApiError(resp *http.Response) *ApiError {
    var apiExileError ExileErrorResponse
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    err = json.Unmarshal(body, &apiExileError)
    if err != nil {
        log.Fatal(err)
    }

    respError := &ApiError{
        StatusCode: resp.StatusCode,
        Header:     resp.Header,
        URL:        resp.Request.URL,
        Data:       apiExileError,
        Body:       string(body),
    }

    return respError
}

func (err ApiError) Error() string {
    return fmt.Sprintf("Api call returned status %d, \nURL:\t\t%s\nBody:\t\t%s", err.StatusCode, err.URL, err.Body)
}

func (err ExileErrorResponse) Error() string {
    return err.SingleError.Message
}

func (err ExileError) Error() string {
    return err.Message
}
