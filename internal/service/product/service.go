package product

import (
	"errors"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) AllProducts() []Product {
	return allProducts
}

func (s *Service) ItemByID(id int) (Product, error) {
	if id <= 0 || id > len(allProducts) {
		return Product{}, errors.New("no such item")
	}

	var item Product

	for _, p := range allProducts {
		if p.Id == id {
			item = p
		}
	}

	return item, nil
}
