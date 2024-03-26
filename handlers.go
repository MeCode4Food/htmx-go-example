package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
)

func GetMainTemplateHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("tpl/main.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, MainPageContext{PageTitle: "Main Page"})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func GetCatsTemplateHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("tpl/cats.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "cats", MainPageContext{PageTitle: "Cats Page", Cats: cats})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func GetCatEditTemplateHandler(w http.ResponseWriter, r *http.Request) {
	var cat *Cat
	for _, c := range cats {
		if c.ID == mux.Vars(r)["id"] {
			cat = c
			break
		}
	}
	if cat == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	tmpl, err := template.ParseFiles("tpl/cats-edit.tmpl")
	if err != nil {
		fmt.Println("Error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	v := map[string]string{
		"dialogOpen": "editCatDialog",
	}
	jsonData, err := json.Marshal(v)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	jsonString := string(jsonData)
	w.Header().Add("HX-Trigger-After-Settle", jsonString)
	err = tmpl.Execute(w, EditCatPageContext{
		EditType:  EditTypeUpdate,
		PageTitle: "Edit Cat " + mux.Vars(r)["id"] + " Page",
		Cat:       cat,
	})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func GetCatCreateTemplateHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("tpl/cats-edit.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	v := map[string]string{
		"dialogOpen": "editCatDialog",
	}
	jsonData, err := json.Marshal(v)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	jsonString := string(jsonData)
	w.Header().Add("HX-Trigger-After-Settle", jsonString)
	err = tmpl.Execute(w, EditCatPageContext{EditType: EditTypeCreate, PageTitle: "New Cat"})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func PostCreateNewTemplateHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("tpl/cats.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	name := r.FormValue("name")
	breed := r.FormValue("breed")
	id := rand.Intn(1000000)
	c := &Cat{
		Name:  name,
		Breed: breed,
		ID:    fmt.Sprintf("%d", id),
	}
	cats = append(cats, c)
	v := map[string]string{
		"dialogClose": "editCatDialog",
	}
	jsonData, err := json.Marshal(v)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	jsonString := string(jsonData)
	w.Header().Add("HX-Trigger-After-Settle", jsonString)
	err = tmpl.ExecuteTemplate(w, "cat", c)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func DeleteCatByIDTemplateHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var removed bool
	for _, c := range cats {
		// remove cat from cats
		if c.ID == id {
			cats = append(cats[:0], cats[1:]...)
			removed = true
			break
		}
	}
	if !removed {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func PutCatByIDTemplateHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var (
		cat *Cat
		err error
	)
	for _, c := range cats {
		// remove cat from cats
		if c.ID == id {
			cat = c
			break
		}
	}
	if cat == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	name := r.FormValue("name")
	breed := r.FormValue("breed")
	cat.Name = name
	cat.Breed = breed
	tmpl, err := template.ParseFiles("tpl/cats.tmpl")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	v := map[string]string{
		"dialogClose": "editCatDialog",
	}
	jsonData, err := json.Marshal(v)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	jsonString := string(jsonData)
	w.Header().Add("HX-Trigger-After-Settle", jsonString)
	err = tmpl.ExecuteTemplate(w, "cat", cat)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

}
