package queue

import (
	"time"
)

type Queue interface {
	Push([]byte, time.Time)
	Pop() ([]byte)
	Len() int64
}