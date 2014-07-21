package goexile

import (
    "time"
)

type League struct {
    Description string         `json:"description"`
    EndAt       *time.Time     `json:"endAt"`
    Event       bool           `json:"event"`
    Id          string         `json:"id"`
    RegisterAt  *time.Time     `json:"registerAt"`
    Rules       []Rule         `json:"rules"`
    StartAt     *time.Time     `json:"startAt"`
    Url         string         `json:"url"`
    Ladder      []LadderEntity `json:"ladder"`
}

type Rule struct {
    Id          int64  `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
}

type LadderEntity struct {
    Total   int     `json:"total"`
    Entries []Entry `json:"entries"`
}

type Entry struct {
    Online    bool            `json:"online"`
    Rank      int             `json:"rank"`
    Dead      bool            `json:"dead"`
    Character CharacterEntity `json:"character"`
    Account   AccountEntity   `json:"account"`
}

type AccountEntity struct {
    Name       string                  `json:"name"`
    Challenges AccountChallengesEntity `json:"challenges"`
    Twitch     AccountTwitchEntity     `json:"twitch"`
}

type AccountChallengesEntity struct {
    Total int `json:"total"`
}

type AccountTwitchEntity struct {
    Name string `json:"name"`
}

type CharacterEntity struct {
    Name       string `json:"name"`
    Level      int    `json:"level"`
    Class      string `json:"class"`
    Experience uint64 `json:"experience"`
}
