package routes

import (
	"log"
	"net/http"
)

func Main() {
	log.Fatal(http.ListenAndServe(":8000", nil))
}
