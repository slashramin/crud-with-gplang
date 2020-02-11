package models

import (
	"github.com/slashramin/crud-with-gplang/pkg/config"
	"github.com/jinzhu/gorm"
)

type Note struct {
	gorm.Model
	//Id          string `json:"id"`
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Note{})
}

func (b *Note) CreateNote() *Note {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func  GetAllNotes() []Note {
	var Notes []Note
	db.Find(&Notes)
	return Notes
}

func GetNoteById(Id int64) (*Note , *gorm.DB){
	var getNote Note
	db:=db.Where("ID = ?", Id).Find(&getNote)
	return &getNote, db
}

func DeleteNote(ID int64) Note {
	var note Note
	db.Where("ID = ?", ID).Delete(note)
	return note
}