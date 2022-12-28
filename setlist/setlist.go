package setlist

import (
	"crypto/md5"
	"fmt"
)

type SetList[T any] struct {
	data map[string]T
}

func New[T any]() *SetList[T] {
	return &SetList[T]{
		data: map[string]T{},
	}
}

func (s *SetList[T]) Insert(value T) {
	obectIdentifier := s.generateId(value)
	_, ok := s.data[obectIdentifier]
	if !ok {
		s.data[obectIdentifier] = value
	}
}

func (s *SetList[T]) InsertSlice(slice []T) {
	for _, sl := range slice {
		s.Insert(sl)
	}
}

func (s SetList[T]) Has(value T) bool {
	objectId := s.generateId(value)
	_, ok := s.data[objectId]
	return ok
}

func (s SetList[T]) generateId(value T) string {
	sValue := fmt.Sprint(value)
	return fmt.Sprint(md5.Sum([]byte(sValue)))
}

func (s SetList[T]) Size() int {
	return len(s.data)
}

func (s *SetList[T]) Remove(value T) {
	objectId := s.generateId(value)
	_, ok := s.data[objectId]
	if ok {
		delete(s.data, objectId)
	}
}

func (s SetList[T]) AsSlice() []T {
	r := make([]T, 0, len(s.data))
	for _, v := range s.data {
		r = append(r, v)
	}
	return r
}
