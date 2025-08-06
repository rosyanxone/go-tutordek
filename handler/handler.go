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
		{ID: 1, Name: "Mobilio", Price: 220000000, Stock: 3},
		{ID: 2, Name: "Xpander", Price: 500000000, Stock: 1},
		{ID: 3, Name: "Pajero", Price: 340000000, Stock: 2},
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
