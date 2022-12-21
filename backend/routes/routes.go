package routes

import (
	"encoding/json"
	"fmt"
	"github.com/Dispanel/config"
	"github.com/Dispanel/controllers"
	"github.com/Dispanel/middleware"
	"github.com/Dispanel/utils"
	"io/ioutil"
	"net/http"
)

func Run() {
	config := config.LoadConfig()

	http.Handle("/", middleware.Authorization(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", r.RemoteAddr)
	})))

	http.Handle("/containers", middleware.Authorization(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			json.NewEncoder(w).Encode(controllers.GetALLContainers())
		} else {
			utils.WarningHandling("ROUTES", "Unauthorized access to "+r.Host)
			w.WriteHeader(http.StatusUnauthorized)
		}

	})))

	http.Handle("/stop", middleware.Authorization(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				utils.ErrorHandling(err)
			}
			var id string
			err = json.Unmarshal(body, &id)
			if err != nil {
				utils.ErrorHandling(err)
			}

			if id == "" {
				utils.WarningHandling("ROUTES", "No id provided")
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			status, err := controllers.StopContainer(id)
			if err != nil {
				utils.WarningHandling("ROUTES", "Error stopping container: "+err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(status)
		} else {
			utils.WarningHandling("ROUTES", "Unauthorized access to "+r.Host)
			w.WriteHeader(http.StatusUnauthorized)
		}
	})))

	utils.PrintHandling("SERVER", "Server started on port :"+config.Port)
	err := http.ListenAndServe(":"+config.Port, nil)
	if err != nil {
		utils.ErrorHandling(err)
	}

}