package main

import (
	"github.com/nats-io/stan.go"
	"strconv"
)

type MsgProcessor struct {
	SubjName      string
	cache         Saver
	SubsAmount    int
	StanClusterID string
	ClientID      string
	QGroup        string
	DurableName   string
}

func (p *MsgProcessor) Connect(cb stan.MsgHandler) error {
	for i := 0; i < p.SubsAmount; i++ {
		sc, err := stan.Connect(p.StanClusterID, p.ClientID+strconv.Itoa(i))
		if err != nil {
			return err
		}
		_, err = sc.QueueSubscribe(p.SubjName, p.QGroup, cb, stan.DurableName(p.DurableName))
		if err != nil {
			return err
		}

	}

	return nil
}

func (p *MsgProcessor) MessageHandler() stan.MsgHandler {
	return p.cache.SaveInformation
}
