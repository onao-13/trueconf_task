package middleware

import (
	"log"
	"refactoring/internal/app/payload"
	"refactoring/internal/app/serverr"
	"time"
)

type User struct {
	store Store
	log log.Logger
}

func New(store Store, log log.Logger) User {
	return User{store: store, log: log}
}

func (u *User) Search() (users payload.UserStore, err error) {
	users, err = u.store.Get()
	if err != nil {
		u.log.Println("Ошибка получения хранилища: ", err.Error())
	}
	return
}

func (u *User) Create(req payload.CreateUserRequest) (id int64, err error) {
	user := payload.CreateUser(req.DisplayName, req.Email)

	if err = u.store.Add(user); err != nil {
		u.log.Println("Ошибка добавления пользователя: ", err.Error())
		return
	}

	store, err := u.store.Get()
	if err != nil {
		u.log.Println("Ошибка получения хранилища: ", err.Error())
		return 
	}

	return store.Increment, nil 
}

func (u *User) Get(id int64) (user payload.User, err error) {
	return u.store.GetById(id)
}

func (u *User) Update(newUser payload.UpdateUserRequest, id int64) (err error) {
	user := payload.User{
		CreatedAt: time.Now(),
		DisplayName: newUser.DisplayName,
		Email: newUser.Email,
	}

	if err = u.store.Update(id, user); err != nil {
		u.log.Println("Ошибка обновления пользователя: ", err.Error())
		return
	}

	return nil
}

func (u *User) Delete(id int64) error {
	user, err := u.store.GetById(id)
	if err != nil {
		return err
	}

	if len(user.DisplayName) == 0 && len(user.Email) == 0 {
		return &serverr.UserNotFound{}
	}

	if err := u.store.Delete(id); err != nil {
		u.log.Println("Ошибка удаления пользователя: ", err.Error())
		return err
	}
	return nil
}
