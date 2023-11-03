package controller

import (
	"encoding/json"
	"net/http"
	"refactoring/internal/app/middleware"
	"refactoring/internal/app/payload"
	"refactoring/internal/app/serverr"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type User struct {
	s middleware.User
}

func New(s middleware.User) User {
	return User{s: s}
}

func (u *User) Search(w http.ResponseWriter, r *http.Request) {
	list, err := u.s.Search()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, list)
}

func (u *User) Create(w http.ResponseWriter, r *http.Request) {
	var newUser payload.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(newUser.DisplayName) == 0 || len(newUser.Email) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := u.s.Create(newUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]interface{}{
		"user_id": id, 
	})
}

func (u *User) Get(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return
	}

	user, err := u.s.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(user.DisplayName) == 0 && len(user.Email) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	render.JSON(w, r, user)
}

func (u *User) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	var updateUser payload.UpdateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&updateUser); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(updateUser.DisplayName) == 0 || len(updateUser.Email) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := u.s.Update(updateUser, id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusNoContent)
}

func (u *User) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")	

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := u.s.Delete(id); err != nil {
		switch err.(type) {
		case *serverr.UserNotFound:
			w.WriteHeader(http.StatusNotFound)
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	render.Status(r, http.StatusNoContent)
}
