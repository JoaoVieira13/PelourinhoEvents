package helpers

import (
	"fmt"
	"log"
	"net/http"
	"pe/services"

	"gorm.io/gorm"
)

type Error interface {
	error
	Status() int
}

type StatusError struct {
	Code int
	Err  error
}

func (se StatusError) Error() string {
	return se.Err.Error()
}

func (se StatusError) Status() int {
	return se.Code
}

type Env struct {
	DB *gorm.DB
}

type Handler struct {
	*Env
	H func(e *Env, w http.ResponseWriter, r *http.Request) error
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.H(h.Env, w, r)
	if err != nil {
		switch e := err.(type) {
		case Error:
			log.Printf("HTTP %d - %s", e.Status(), e)
			http.Error(w, e.Error(), e.Status())
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		}
	}
}

func GetIndex(env *Env, w http.ResponseWriter, r *http.Request) error {
	eventService := services.NewEventService(env.DB)
	events, err := eventService.GetEvents()
	if err != nil {
		return StatusError{500, err}
	}

	fmt.Fprintf(w, "%+v", events)
	return nil
}
