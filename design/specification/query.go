package main

/*
import (
	"fmt"
	"github.com/google/uuid"
	"strings"
)

type MaterialType string

type Product struct {
	ID            uuid.UUID
	Material      MaterialType
	IsDeliverable bool
	Quantity      int
}

type ProductSpecification interface {
	Query() string
	Value() []interface{}
}

type AndSpecification struct {
	specifications []ProductSpecification
}

func NewAndSpecification(specifications ...ProductSpecification) ProductSpecification {
	return AndSpecification{
		specifications: specifications,
	}
}

func (s AndSpecification) Query() string {
	var queries []string
	for _, specification := range s.specifications {
		queries = append(queries, specification.Query())
	}

	query := strings.Join(queries, " AND ")

	return fmt.Sprintf("(%s)", query)
}

func (s AndSpecification) Value() []interface{} {
	var values []interface{}
	for _, specification := range s.specifications {
		values = append(values, specification.Value()...)
	}
	return values
}

type OrSpecification struct {
	specifications []ProductSpecification
}

func NewOrSpecification(specifications ...ProductSpecification) ProductSpecification {
	return OrSpecification{
		specifications: specifications,
	}
}

func (s OrSpecification) Query() string {
	var queries []string
	for _, specification := range s.specifications {
		queries = append(queries, specification.Query())
	}

	query := strings.Join(queries, " OR ")
	return fmt.Sprintf("(%s)", query)
}

func (s OrSpecification) Value() []interface{} {
	var values []interface{}
	for _, specification := range s.specifications {
		values = append(values, specification.Value()...)
	}
	return values
}

type HasAtLeast struct {
	pieces int
}

func NewHasAtLeast(pieces int) ProductSpecification {
	return HasAtLeast{
		pieces: pieces,
	}
}

func (h HasAtLeast) Query() string {
	return "quantity >= ?"
}

func (h HasAtLeast) Value() []interface{} {
	return []interface{}{h.pieces}
}

func IsPlastic() string {
	return "material = 'plastic'"
}

func IsDeliverable() string {
	return "deliverable = 1"
}

type FunctionSpecification func() string

func (fs FunctionSpecification) Query() string {
	return fs()
}

func (fs FunctionSpecification) Value() []interface{} {
	return nil
}

func main() {

	spec := NewOrSpecification(
		NewAndSpecification(
			NewHasAtLeast(10),
			FunctionSpecification(IsPlastic),
			FunctionSpecification(IsDeliverable),
		),
		NewAndSpecification(
			NewHasAtLeast(100),
			FunctionSpecification(IsPlastic),
		),
	)

	fmt.Println(spec.Query())
	// output: ((quantity >= ? AND material = 'plastic' AND deliverable = 1) OR (quantity >= ? AND material = 'plastic'))

	fmt.Println(spec.Value())
	// output: [10 100]
}
*/
