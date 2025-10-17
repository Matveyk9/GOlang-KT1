package main

import (
	"errors"
	"fmt"
)

type Storage interface {
	Set(key string, value string) error   
	Get(key string) (string, error)      
	Delete(key string) error             
}

type InMemoryStorage struct {
	data map[string]string
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		data: make(map[string]string),
	}
}

func (s *InMemoryStorage) Set(key string, value string) error {
	s.data[key] = value
	return nil
}

func (s *InMemoryStorage) Get(key string) (string, error) {
	val, ok := s.data[key]
	if !ok {
		return "", errors.New("ключ не найден")
	}
	return val, nil
}

func (s *InMemoryStorage) Delete(key string) error {
	if _, ok := s.data[key]; !ok {
		return errors.New("ключ не найден")
	}
	delete(s.data, key)
	return nil
}

func main() {

	var store Storage = NewInMemoryStorage()

	store.Set("name", "Alice")
	store.Set("city", "Moscow")

	name, _ := store.Get("name")
	city, _ := store.Get("city")
	fmt.Println("Name:", name)
	fmt.Println("City:", city)

	if _, err := store.Get("age"); err != nil {
		fmt.Println("Ошибка:", err)
	}

	store.Delete("city")
	if _, err := store.Get("city"); err != nil {
		fmt.Println("После удаления:", err)
	}
}
