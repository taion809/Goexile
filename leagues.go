package goexile

import (
    "net/url"
)

func (g *Goexile) GetAllLeagues() ([]League, error) {
    return g.GetLeagues("all")
}

func (g *Goexile) GetEventLeagues() ([]League, error) {
    return g.GetLeagues("event")
}

func (g *Goexile) GetLeagues(event_type string) ([]League, error) {
    v := url.Values{"type": {event_type}}

    var leagues []League
    q := &query{
        endpoint: "/leagues",
        method:   "GET",
        form:     v,
        data:     &leagues,
    }

    err := g.execute(q)
    return leagues, err
}

func (g *Goexile) GetLeague(id string, v url.Values) (League, error) {
    if v == nil {
        v = url.Values{}
    }

    var league League
    q := &query{
        endpoint: "/leagues/" + id,
        method:   "GET",
        form:     v,
        data:     &league,
    }

    err := g.execute(q)
    return league, err
}
