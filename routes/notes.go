package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"notes.com/app/models"
)

func GetNotes(context *gin.Context) {
	notes, err := models.GetAllNotes()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnt fetch notes"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": notes})
}

func GetSingleNote(context *gin.Context) {
	noteId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnt fetch note"})
		return
	}

	note, err := models.GetSingleNote(noteId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnt fetch note by given id"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Note with given id found", "note": note})
}

func CreateNote(context *gin.Context) {
	var note models.Note
	err := context.ShouldBindJSON(&note)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid note input!", "error": err})
		return
	}

	note.UserId = context.GetInt64("userId")
	err = note.CreateNote()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnt create note!", "error": err})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Note created", "note": note})
}

func UpdateNote(context *gin.Context) {
	noteId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnt fetch note"})
		return
	}

	note, err := models.GetSingleNote(noteId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnt fetch note"})
		return
	}

	if note.Id != noteId {
		context.JSON(http.StatusMethodNotAllowed, gin.H{"message": "Wrong NOTE"})
		return
	}

	err = context.ShouldBindJSON(&note)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid note input!", "error": err})
		return
	}

	note.UserId = 1

	err = note.UpdateNote()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnt update note!", "error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Note updated", "note": note})
}

func DeleteNote(context *gin.Context) {
	noteId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnt delete note"})
		return
	}

	note, err := models.GetSingleNote(noteId)
	userId := context.GetInt64("userId")

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnt delete note"})
		return
	}

	if note.UserId != userId {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Not the user who created the note"})
		return
	}

	err = models.DeleteNote(noteId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnt delete note"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Note deleted"})
}
