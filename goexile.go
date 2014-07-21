package goexile

import (
    "encoding/json"
    "log"
    "net/http"
    "net/url"
)

type Goexile struct {
    client *http.Client
}

type query struct {
    endpoint string
    method   string
    form     url.Values
    data     interface{}
}

const (
    BaseUrl = "http://api.pathofexile.com"
)

func NewApi() *Goexile {
    return &Goexile{
        client: &http.Client{},
    }
}

func (g *Goexile) execute(q *query) error {
    req := buildRequest(q)
    resp, err := g.client.Do(req)

    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return newApiError(resp)
    }

    return json.NewDecoder(resp.Body).Decode(q.data)
}

func buildRequest(q *query) *http.Request {
    qs := q.form.Encode()

    endpoint := BaseUrl + q.endpoint

    if qs != "" {
        endpoint = endpoint + "?" + qs
    }

    req, err := http.NewRequest(q.method, endpoint, nil)

    if err != nil {
        log.Fatal(err)
    }

    return req
}
