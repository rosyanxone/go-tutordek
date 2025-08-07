package handler

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
	"tutordek/entity"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))

	if err != nil {
		log.Println(err)
		http.Error(w, "An error has occured", http.StatusInternalServerError)
		return
	}

	// data := map[string]any{ // the 'any' is equal to interface
	// 	"title": "Software Engineer",
	// 	"age":   22,
	// }

	// data := entity.Product{ID: 1, Name: "Mobilio", Price: 220000000, Stock: 3}

	data := []entity.Product{
		{ID: 1, Name: "Mobilio", Price: 220000000, Stock: 13},
		{ID: 2, Name: "Xpander", Price: 500000000, Stock: 8},
		{ID: 3, Name: "Pajero", Price: 340000000, Stock: 1},
	}

	err = tmpl.Execute(w, data)

	if err != nil {
		log.Println(err)
		http.Error(w, "An error has occured", http.StatusInternalServerError)
		return
	}
}

func MarioHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Mario from Nintendo"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNum, err := strconv.Atoi(id)

	if err != nil || idNum < 1 {
		http.NotFound(w, r)
		return
	}

	// fmt.Fprintf(w, "Product page: %d", idNum)
	data := map[string]any{
		"content": idNum,
	}

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))

	if err != nil {
		log.Println(err)
		http.Error(w, "An error has occured", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)

	if err != nil {
		log.Println(err)
		http.Error(w, "An error has occured", http.StatusInternalServerError)
		return
	}
}

func PostGet(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	case "GET":
		w.Write([]byte("This is a GET method"))
	case "POST":
		w.Write([]byte("This is a POST method"))
	default:
		http.Error(w, "An error has occured", http.StatusBadRequest)
	}
}

func Form(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		templ, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))

		if err != nil {
			log.Println(err)
			http.Error(w, "An error has occured", http.StatusInternalServerError)
			return
		}

		err = templ.Execute(w, nil)

		if err != nil {
			log.Println(err)
			http.Error(w, "An error has occured", http.StatusInternalServerError)
			return
		}

		return
	}

	http.Error(w, "An error has occured", http.StatusBadRequest)
}

func Process(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()

		if err != nil {
			log.Println(err)
			http.Error(w, "An error has occured", http.StatusInternalServerError)
			return
		}

		name := r.Form.Get("name")
		message := r.Form.Get("message")

		data := map[string]any{
			"name":    name,
			"message": message,
		}

		templ, err := template.ParseFiles(path.Join("views", "result.html"), path.Join("views", "layout.html"))

		if err != nil {
			log.Println(err)
			http.Error(w, "An error has occured", http.StatusInternalServerError)
			return
		}

		err = templ.Execute(w, data)

		if err != nil {
			log.Println(err)
			http.Error(w, "An error has occured", http.StatusInternalServerError)
			return
		}

		return
	}

	http.Error(w, "An error has occured", http.StatusBadRequest)
}
