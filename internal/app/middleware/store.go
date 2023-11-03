package middleware

import (
	"encoding/json"
	"os"
	"refactoring/internal/app/payload"
	"refactoring/internal/app/serverr"
)

type Store struct {
	jsonStore string
}

func NewStore(jsonStore string) Store {
	return Store{jsonStore: jsonStore}
}

func (s *Store) Get() (store payload.UserStore, err error) {
	f, err := os.ReadFile(s.jsonStore)
	if err != nil {
		return
	}

	err = json.Unmarshal(f, &store)
	return
}

func (s *Store) Add(user payload.User) (err error) {
	store, err := s.Get()
	if err != nil {
		return err
	}

	store.Increment++
	store.List[store.Increment] = user

	data, err := json.Marshal(&store)
	if err != nil {
		return err
	}

	return os.WriteFile(s.jsonStore, data, os.ModePerm)
}

func (s *Store) Update(id int64, user payload.User) error {
	store, err := s.Get()
	if err != nil {
		return err
	}

	store.List[id] = user

	data, err := json.Marshal(&store)
	if err != nil {
		return err
	}

	return os.WriteFile(s.jsonStore, data, os.ModePerm)
}

func (s *Store) Delete(id int64) error {
	store, err := s.Get()
	if err != nil {
		return err
	}

	delete(store.List, id)

	data, err := json.Marshal(&store)
	if err != nil {
		return err
	}
	
	return os.WriteFile(s.jsonStore, data, os.ModePerm)
}

func (s *Store) GetById(id int64) (user payload.User, err error) {
	store, err := s.Get()
	if err != nil {
		return
	}

	user, ok := store.List[id]
	if !ok {
		return user, &serverr.UserNotFound{}
	}

	return
}

