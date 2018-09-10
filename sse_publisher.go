package lappy

import (
	"encoding/json"

	"github.com/r3labs/sse"
)

// Stream names
const (
	RecordsStream = "records"
	NoiseStream   = "noise"
)

type SsePublisher struct {
	Server *sse.Server
	Log    Logger
}

func NewSsePublisher() *SsePublisher {
	s := sse.New()
	s.CreateStream(RecordsStream)
	s.CreateStream(NoiseStream)
	return &SsePublisher{
		Server: s,
	}
}

func (s *SsePublisher) PublishRecord(rec Record) {
	data, err := json.Marshal(rec)
	if err != nil {
		s.Log.Errorf("failed to marshal record: %s", err)
	}
	s.Server.Publish(RecordsStream, &sse.Event{Data: data})
}

func (s *SsePublisher) PublishNoise(noise uint16) {
	data, err := json.Marshal(noise)
	if err != nil {
		s.Log.Errorf("failed to marshal noise: %s", err)
	}
	s.Server.Publish(NoiseStream, &sse.Event{Data: data})
}
