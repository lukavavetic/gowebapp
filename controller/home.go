package controller

import (
	"creatingWebApp/webapp/viewmodel"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type home struct {
	homeTemplate         *template.Template
	standLocatorTemplate *template.Template
	loginTemplate        *template.Template
}

func (h home) registerRoutes() {
	http.HandleFunc("/", h.handleHome)
	http.HandleFunc("/home", h.handleHome)
	http.HandleFunc("/stand-locator", h.handleStandLocator)
	http.HandleFunc("/login", h.handleLogin)
}

func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewHome()

	//time.Sleep(4 * time.Second) //test timeout middleware

	w.Header().Add("Content-Type", "text/html")

	h.homeTemplate.Execute(w, vm)
}

func (h home) handleStandLocator(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewStandLocator()

	h.standLocatorTemplate.Execute(w, vm)
}

func (h home) handleLogin(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewLogin()

	if r.Method == http.MethodPost {
		err := r.ParseForm()

		if err != nil {
			log.Println(fmt.Errorf("Error logging in: %v", err))
		}

		email := r.Form.Get("email")
		password := r.Form.Get("password")

		if email == "test@gmail.com" && password == "password" {
			http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)

			return
		} else {
			vm.Email = email
			vm.Password = password
		}
	}

	w.Header().Add("Content-Type", "text-html")

	h.loginTemplate.Execute(w, vm)
}