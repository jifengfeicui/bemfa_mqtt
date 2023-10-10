package model

import "github.com/go-ini/ini"

type Topic struct {
	TopicName string
	Parameter []*ini.Section
}
