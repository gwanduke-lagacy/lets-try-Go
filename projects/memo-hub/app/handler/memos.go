// Memo를 위한 handlers를 제공
// Memo 데이터를 우히나 쿼리를 수행

package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/letsget23/go-playground/projects/memo-hub/app/model"
)

// GetAllMemos .
func GetAllMemos(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	memos := []model.Memo{}
	db.Find(&memos)
	respondJSON(w, http.StatusOK, memos)
}

// CreateMemo .
func CreateMemo(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	memo := model.Memo{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&memo); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&memo).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, memo)
}

// GetMemo .
func GetMemo(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}

// UpdateMemo .
func UpdateMemo(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}

// DeleteMemo .
func DeleteMemo(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}

// getMemoOr404 .
func getMemoOr404(db *gorm.DB, id uint, w http.ResponseWriter, r *http.Request) *model.Memo {
	memo := model.Memo{}
	if err := db.First(&memo, model.Memo{ID: id}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &memo
}
