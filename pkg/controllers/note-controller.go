package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/slashramin/crud-with-gplang/pkg/models"
	"github.com/slashramin/crud-with-gplang/pkg/utils"
	"net/http"
	"strconv"
)

var NewNote models.Note

func CreateNote(w http.ResponseWriter, r *http.Request) {
	CreateNote := &models.Note{}
	utils.ParseBody(r, CreateNote)
	b:= CreateNote.CreateNote()
	res,_ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetNote(w http.ResponseWriter, r *http.Request) {
	newNotes:= models.GetAllNotes()
	res, _ := json.Marshal(newNotes)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetNoteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	noteId := vars["noteId"]
	ID, err:= strconv.ParseInt(noteId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	noteDetails, _:= models.GetNoteById(ID)
	res, _ := json.Marshal(noteDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	var updateNote = &models.Note{}
	utils.ParseBody(r, updateNote)
	vars := mux.Vars(r)
	noteId := vars["noteId"]
	ID, err:= strconv.ParseInt(noteId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	noteDetails, db:= models.GetNoteById(ID)
	if updateNote.Name != "" {
		noteDetails.Name = updateNote.Name
	}
	if updateNote.Author != "" {
		noteDetails.Author = updateNote.Author
	}
	if updateNote.Publication != "" {
		noteDetails.Publication = updateNote.Publication
	}
	db.Save(&noteDetails)
	res, _ := json.Marshal(noteDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	noteId := vars["noteId"]
	ID, err := strconv.ParseInt(noteId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	note := models.DeleteNote(ID)
	res, _ := json.Marshal(note)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}