package service

import (
	"avitoStart/internal/model"
	"fmt"
	"log"
)

type Database interface {
	//users
	AddUser(user model.User) (bool, error)
	DeleteUser(id string) (bool, error)
	ExtractUsers() ([]model.User, error)

	//slug
	DeleteSlug(name string) (bool, error)
	CreateSlug(name string) (bool, error)
	ExecSlugNamesUser(iduser string) ([]model.Slug, error)
	DeleteRelation(iduser string, name string) (bool, error)
	CreateRelation(iduser string, name string) (bool, error)
}
type Service struct {
	db Database
}

func New(db Database) *Service {
	return &Service{
		db: db,
	}
}

// work with users
func (s *Service) AddUser(user model.User) (bool, error) {
	fmt.Println("User:", user)
	res, err := s.db.AddUser(user)
	if err != nil {
		return false, err
	}
	return res, err
}

func (s *Service) DeleteUser(id string) (bool, error) {
	res, err := s.db.DeleteUser(id)
	if err != nil {
		return res, err
	}
	return res, err
}

func (s *Service) ExtractUsers() ([]model.User, error) {
	res, err := s.db.ExtractUsers()
	if err != nil {
		return res, err
	}
	return res, err
}

//work with slug

func (s *Service) DeleteSlug(name string) (bool, error) {
	res, err := s.db.DeleteSlug(name)
	if err != nil {
		return res, err
	}
	return res, err
}

func (s *Service) CreateSlug(name string) (bool, error) {
	res, err := s.db.CreateSlug(name)
	if err != nil {
		return res, err
	}
	return res, err
}

func (s *Service) ExecSlugNamesUser(iduser string) ([]model.Slug, error) {
	res, err := s.db.ExecSlugNamesUser(iduser)
	if err != nil {
		return res, err
	}
	return res, err
}

func (s *Service) MasterFunc(data model.MasterData) (bool, error) {
	if data.MasAdd != nil && len(data.MasAdd) > 0 {
		for _, d := range data.MasAdd {
			log.Println(d)
			_, err := s.db.CreateRelation(data.Id, d)
			if err != nil {
				return false, err
			}
		}
	}
	if data.MasDel != nil && len(data.MasDel) > 0 {
		for _, del := range data.MasDel {
			log.Println(del)
			_, err := s.db.DeleteRelation(data.Id, del)
			if err != nil {
				return false, err
			}
		}
	}
	return true, nil
}
