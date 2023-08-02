package server

import (
	"fmt"
	"main/internal/utils"
	"net/http"
	"net/url"
)

type RequestError struct {
	error string
}

func (s *SServer) CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		event := r.FormValue("event")
		date := r.FormValue("date")

		fmt.Fprintf(w, "%s is a %s\n", event, date)
		res := s.store.CreateEvent(event, date)
		fmt.Println(res)
		w.Write([]byte("Создание нового события..."))
		w.WriteHeader(http.StatusOK)
	} else {
		// http.Error(w, fmt.Sprintf("Nothing found %v %v", r.Method, r.URL), http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		sj, err := utils.SerializingToJSON(RequestError{
			error: "Ничего не найдено",
		})
		if err == nil {
			w.Write(sj)
		} else {
			fmt.Fprintf(w, "Json serializing err: %v", err)
		}

	}
}

func (s *SServer) UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.Write([]byte("Абновление события..."))
	}
}

func (s *SServer) DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.Write([]byte("Удаление события..."))
	}
}

func (s *SServer) GetEventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Write([]byte("Получение списка событий за день..."))
	}
}

func (s *SServer) GetEventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Write([]byte("Получение списка событий за неделю..."))
	}
}

func (s *SServer) GetEventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		query, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "invalid request")
			return
		}
		date := query.Get("date")
		if len(date) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "missing date")
			return
		}
		userId := query.Get("user_id")
		if len(userId) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "missing user_id")
			return
		}
		w.WriteHeader(http.StatusOK)
		// res := s.store.GetEventsForMonth(date)
		fmt.Println(date)
		// body, _ := utils.SerializingToJSON(res)
		// w = body
	}
}
