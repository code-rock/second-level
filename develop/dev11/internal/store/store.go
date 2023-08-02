package store

import (
	"strings"
	"sync"
	"time"

	"github.com/go-eden/etime"
)

type SEvent struct {
	ID    int64  `json:"id"`
	Event string `json:"event"`
	Date  string `json:"date"`
}

type EventStore struct {
	mu     sync.RWMutex
	events map[int64]SEvent
}

func New() (store *EventStore) {
	var once sync.Once
	once.Do(func() {
		store = &EventStore{
			events: map[int64]SEvent{},
		}
	})
	return store
}

func (es *EventStore) CreateEvent(e string, d string) SEvent {
	es.mu.Lock()
	defer es.mu.Unlock()

	id := etime.NowMicrosecond()

	es.events[id] = SEvent{
		ID:    id,
		Event: e,
		Date:  d,
	}

	return es.events[id]
}

func (es *EventStore) UpdateEvent(e SEvent) {
	if _, ok := es.events[e.ID]; ok {
		es.mu.Lock()
		defer es.mu.Unlock()

		es.events[e.ID] = e
	}
}

func (es *EventStore) DeleteEvent(e SEvent) {
	if _, ok := es.events[e.ID]; ok {
		es.mu.Lock()
		defer es.mu.Unlock()

		delete(es.events, e.ID)
	}
}

func (es *EventStore) GetEventsForDay(date string) []SEvent {
	var samplePerDay []SEvent

	es.mu.RLock()
	defer es.mu.RUnlock()

	for _, v := range es.events {
		if v.Date == date {
			samplePerDay = append(samplePerDay, v)
		}
	}

	return samplePerDay
}

func (es *EventStore) GetEventsForWeek(date string) []SEvent {
	t, _ := time.Parse("2003-12-25", date)
	t2 := t.AddDate(0, 0, 6)
	var samplePerDay []SEvent

	es.mu.RLock()
	defer es.mu.RUnlock()

	for _, v := range es.events {
		if ct, err := time.Parse("2003-12-25", v.Date); err == nil && ct.Before(t2) && ct.After(t) {
			samplePerDay = append(samplePerDay, v)
		}
	}

	return samplePerDay
}

func (es *EventStore) GetEventsForMonth(date string) []SEvent {
	dateArr := strings.Split(date, "-")
	dateMonth := strings.Join(dateArr[0:1], "-")
	var samplePerDay []SEvent

	es.mu.RLock()
	defer es.mu.RUnlock()

	for _, v := range es.events {
		if strings.HasPrefix(v.Date, dateMonth) {
			samplePerDay = append(samplePerDay, v)
		}
	}

	return samplePerDay
}
