package goexile

import (
    "net/url"
)

func (g *Goexile) GetLadder(id string, v url.Values) (LadderEntity, error) {
    if v == nil {
        v = url.Values{}
    }

    var ladder LadderEntity
    q := &query{
        endpoint: "/ladders/" + id,
        method:   "GET",
        form:     v,
        data:     &ladder,
    }

    err := g.execute(q)
    return ladder, err
}
