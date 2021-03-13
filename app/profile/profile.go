package profile

import (
	"github.com/pristupaanastasia/matcha42/app/model"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request){

	if r.Method == "POST"{
		_ = 1
	}
	if r.Method == "GET" {
		http.ServeFile(w, r, model.ServerVue +"profile")
	}
}
