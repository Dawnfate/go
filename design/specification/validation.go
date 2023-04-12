package main

import (
	"fmt"
	"github.com/google/uuid"
)

type MaterialType string

type Product struct {
	ID            uuid.UUID
	Material      MaterialType
	IsDeliverable bool
	Quantity      int
}

type ProductSpecification interface {
	IsValid(product Product) bool
}

type AndSpecification struct {
	specifications []ProductSpecification
}

func NewAndSpecification(specifications ...ProductSpecification) ProductSpecification {
	return AndSpecification{
		specifications: specifications,
	}
}

func (s AndSpecification) IsValid(product Product) bool {
	for _, specification := range s.specifications {
		if !specification.IsValid(product) {
			return false
		}
	}
	return true
}

type HasAtLeast struct {
	pieces int
}

func NewHasAtLeast(pieces int) ProductSpecification {
	return HasAtLeast{
		pieces: pieces,
	}
}

func (h HasAtLeast) IsValid(product Product) bool {
	return product.Quantity >= h.pieces
}

func IsPlastic(product Product) bool {
	return product.Material == "Plastic"
}

func IsDeliverable(product Product) bool {
	return product.IsDeliverable
}

type FunctionSpecification func(product Product) bool

func (fs FunctionSpecification) IsValid(product Product) bool {
	return fs(product)
}

func main() {
	spec := NewAndSpecification(
		NewHasAtLeast(10),
		FunctionSpecification(IsPlastic),
		FunctionSpecification(IsDeliverable),
	)

	fmt.Println(spec.IsValid(Product{}))
	// output: false

	fmt.Println(spec.IsValid(Product{
		Material:      "Plastic",
		IsDeliverable: true,
		Quantity:      50,
	}))
	// output: true
}
