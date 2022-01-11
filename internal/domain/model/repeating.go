package model

import "time"

type Repeating struct {
	Count    int
	WaitTime time.Duration
}
