package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("Models: no matching record found")

type Snippet struct {
	ID    int
	Title string
	content string
	created time.Time
	expires time.Time
}
