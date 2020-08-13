package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func checkAuthentication() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			flag := true
			fmt.Println("Checking authentication...")
			if flag {
				f(writer, request)
			} else {
				fmt.Println("Authentication failed", time.Now())
				return
			}
		}
	}
}

func logger() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			start := time.Now()
			// Funciones Anonimas
			defer func() {
				log.Println(request.URL.Path, time.Since(start))
			}()
			f(writer, request)
		}
	}
}
