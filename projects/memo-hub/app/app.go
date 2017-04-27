// 역할
// 1. 앱 초기화
// 2. 라우터 설정
// 3. 앱으로의 모든 핸들러 등록
// 4. 앱 실행

package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/letsget23/go-playground/projects/memo-hub/app/handler"
	"github.com/letsget23/go-playground/projects/memo-hub/app/model"
	"github.com/letsget23/go-playground/projects/memo-hub/config"
)

// App 은 라우터와 DB 인스턴스를 가집니다.
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Initialize 는 미리 정의된 설정으로 앱을 초기화한다.
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.Charset)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// setRouters 라우팅을 설정한다.
func (a *App) setRouters() {
	a.Get("/memos", a.GetAllMemos)
	a.Post("/memos", a.CreateMemo)
	a.Get("/memos/{id}", a.GetMemo)
	a.Put("/memos/{id}", a.UpdateMemo)
	a.Delete("/memos/{id}", a.DeleteMemo)
}

// Get 은 GET 메소드를 위해 라우터를 래핑한다.
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post 은 GET 메소드를 위해 라우터를 래핑한다.
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put 은 GET 메소드를 위해 라우터를 래핑한다.
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete 은 GET 메소드를 위해 라우터를 래핑한다.
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

/*
** Memos Handlers
 */

// GetAllMemos .
func (a *App) GetAllMemos(w http.ResponseWriter, r *http.Request) {
	handler.GetAllMemos(a.DB, w, r)
}

// CreateMemo .
func (a *App) CreateMemo(w http.ResponseWriter, r *http.Request) {
	handler.CreateMemo(a.DB, w, r)
}

// GetMemo .
func (a *App) GetMemo(w http.ResponseWriter, r *http.Request) {
	handler.GetMemo(a.DB, w, r)
}

// UpdateMemo .
func (a *App) UpdateMemo(w http.ResponseWriter, r *http.Request) {
	handler.UpdateMemo(a.DB, w, r)
}

// DeleteMemo .
func (a *App) DeleteMemo(w http.ResponseWriter, r *http.Request) {
	handler.DeleteMemo(a.DB, w, r)
}

// Run 앱을 라우트에 실행한다.
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
