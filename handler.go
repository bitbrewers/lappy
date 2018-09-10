package lappy

import (
	"time"

	"github.com/bitbrewers/tranx2"
)

type Record struct {
	ID     int64
	RaceID int64
	tranx2.Passing
}

type Race struct {
	ID      int64
	Started *time.Time
	Ended   *time.Time
}

type Storage interface {
	Save(passing tranx2.Passing) (rec Record, err error)
}

type Publisher interface {
	PublishRecord(rec Record)
	PublishNoise(noise uint16)
}

type Handler struct {
	Log       Logger
	Storage   Storage
	Publisher Publisher
}

func (h *Handler) OnPassing(passing tranx2.Passing) {
	h.Log.Debugf("record: %+v", passing)
	rec, err := h.Storage.Save(passing)

	if err == ErrNoOngoingRace {
		return
	}

	if err != nil {
		h.Log.Fatalf("could not store record %+v: %s", rec, err)
	}

	h.Publisher.PublishRecord(rec)
}

func (h *Handler) OnNoise(noise uint16) {
	h.Log.Debugf("noise: %d", noise)
	h.Publisher.PublishNoise(noise)
}

func (h *Handler) OnError(err error) {
	h.Log.Errorf("%s", err)
}
