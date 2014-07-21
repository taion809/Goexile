package goexile

import (
    "encoding/json"
    "net/url"
    "strconv"
)

func (g *Goexile) GetLeagueRules() ([]Rule, error) {
    var rules []Rule
    q := &query{
        endpoint: "/league-rules/",
        method:   "GET",
        form:     url.Values{},
        data:     &rules,
    }

    err := g.execute(q)
    return rules, err
}

func (g *Goexile) GetLeagueRule(id int64) (Rule, error) {
    var rule Rule
    q := &query{
        endpoint: "/league-rules/" + strconv.FormatInt(id, 10),
        method:   "GET",
        form:     url.Values{},
        data:     &rule,
    }

    err := g.execute(q)
    return rule, err
}

func (r *Rule) UnmarshalJSON(b []byte) error {
    var i map[string]interface{}

    err := json.Unmarshal(b, &i)
    if err != nil {
        return err
    }

    r.Description = i["description"].(string)
    r.Name = i["name"].(string)

    switch id := i["id"].(type) {
    case int64:
        r.Id = id
    case string:
        r.Id, err = strconv.ParseInt(id, 10, 64)
        if err != nil {
            return err
        }
    default:
        r.Id = 0
    }

    return nil
}
