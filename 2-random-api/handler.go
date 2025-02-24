package main

import (
	"go/adv-random-api/random"
	"net/http"
)

type HandlerRandomNumber struct{}

func NewRandomHandler(router *http.ServeMux) {
	handler := &HandlerRandomNumber{}
	router.HandleFunc("/random", handler.Random())
}

func (handler *HandlerRandomNumber) Random() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(random.Random()))
	}
}
