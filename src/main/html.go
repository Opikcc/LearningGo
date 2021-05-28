package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Notes struct {
	Title       string
	Description string
	CreatedOn   time.Time
}

//View Model for edit
type EditNote struct {
	Notes
	Id string
}

//Store for the Notes collection
var noteStores = make(map[string]Notes)

//Variable to generate key for the collection
var ids int = 0

var templates map[string]*template.Template

//Compile view templates
func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	templates["index"] = template.Must(template.ParseFiles("../../html/templates/index.html", "../../html/templates/base.html"))

	templates["add"] = template.Must(template.ParseFiles("../../html/templates/add.html", "../../html/templates/base.html"))

	templates["edit"] = template.Must(template.ParseFiles("../../html/templates/edit.html", "../../html/templates/base.html"))
}

//Render templates for the given name, template definition and data object
func renderTemplate(w http.ResponseWriter, name string, template string, viewModel interface{}) {
	// Ensure the template exists in the map.
	tmpl, ok := templates[name]
	if !ok {
		http.Error(w, "The template does not exist.", http.StatusInternalServerError)
	}
	err := tmpl.ExecuteTemplate(w, template, viewModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//Handler for "/notes/save" for save a new item into the data store
func saveNote(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	title := r.PostFormValue("title")
	desc := r.PostFormValue("description")
	note := Notes{title, desc, time.Now()}

	//increment the value of id for generating key for the map
	ids++

	//convert id value to string
	k := strconv.Itoa(ids)
	noteStores[k] = note
	http.Redirect(w, r, "/", 302)
}

//Handler for "/notes/add" for add a new item
func addNote(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "add", "base", nil)
}

//Handler for "/notes/edit/{id}" to edit an existing item
func editNote(w http.ResponseWriter, r *http.Request) {
	var viewModel EditNote
	//Read value from route variable
	vars := mux.Vars(r)
	k := vars["ids"]
	if note, ok := noteStores[k]; ok {
		viewModel = EditNote{note, k}
	} else {
		http.Error(w, "Could not find the resource to edit.", http.StatusBadRequest)
	}
	renderTemplate(w, "edit", "base", viewModel)
}

//Handler for "/notes/update/{id}" which update an item into the data store
func updateNote(w http.ResponseWriter, r *http.Request) {
	//Read value from route variable
	vars := mux.Vars(r)
	k := vars["ids"]
	var noteToUpd Notes
	if note, ok := noteStores[k]; ok {
		r.ParseForm()
		noteToUpd.Title = r.PostFormValue("title")
		noteToUpd.Description = r.PostFormValue("description")
		noteToUpd.CreatedOn = note.CreatedOn

		//delete existing item and add the updated item
		delete(noteStores, k)
		noteStores[k] = noteToUpd
	} else {
		http.Error(w, "Could not find the resource to update.", http.StatusBadRequest)
	}
	http.Redirect(w, r, "/", 302)
}

//Handler for "/notes/delete/{id}" which delete an item form the store
func deleteNote(w http.ResponseWriter, r *http.Request) {
	//Read value from route variable
	vars := mux.Vars(r)
	k := vars["ids"]

	// Remove from Store
	if _, ok := noteStores[k]; ok {
		//delete existing item
		delete(noteStores, k)
	} else {
		http.Error(w, "Could not find the resource to delete.", http.StatusBadRequest)
	}
	http.Redirect(w, r, "/", 302)
}

//Handler for "/" which render the index page
func getNotes(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", "base", noteStores)
}

//Entry point of the program
func main() {
	r := mux.NewRouter().StrictSlash(false)
	fs := http.FileServer(http.Dir("public"))
	r.Handle("/public/", fs)
	r.HandleFunc("/", getNotes)
	r.HandleFunc("/notes/add", addNote)
	r.HandleFunc("/notes/save", saveNote)
	r.HandleFunc("/notes/edit/{ids}", editNote)
	r.HandleFunc("/notes/update/{ids}", updateNote)
	r.HandleFunc("/notes/delete/{ids}", deleteNote)
	server := &http.Server{
		Addr:    ":9000",
		Handler: r,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
