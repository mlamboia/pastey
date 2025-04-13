package entity

import "time"

type ClipboardItem struct {
    ID        int
    Content   string
	Pinned 	  bool
    Timestamp time.Time
}