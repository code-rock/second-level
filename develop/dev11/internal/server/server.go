package server

import (
	"fmt"
	"log"
	"main/internal/middleware"
	"main/internal/store"
	"net/http"
	"sync"
)

type SServer struct {
	store *store.EventStore
}

func New(store *store.EventStore) (server *SServer) {
	var once sync.Once
	once.Do(func() {
		server = &SServer{
			store: store,
		}
	})
	return server
}

func (s *SServer) Api(port interface{}) {
	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", s.CreateEventHandler)
	mux.HandleFunc("/update_event", s.UpdateEventHandler)
	mux.HandleFunc("/delete_event", s.DeleteEventHandler)
	mux.HandleFunc("/events_for_day", s.GetEventsForDayHandler)
	mux.HandleFunc("/events_for_week", s.GetEventsForWeekHandler)
	mux.HandleFunc("/events_for_month", s.GetEventsForMonthHandler)

	handler := middleware.Logging(mux)
	handler = middleware.PanicRecovery(handler)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), handler)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
