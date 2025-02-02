package memory

import (
	"testing"

	"github.com/google/uuid"
	"github.com/percybolmer/tavern/domain/product"
)

func TestMemoryProductRepository_Add(t *testing.T) {
	repo := New()
	newProduct, err := product.NewProduct("Beer", "Good for you're health", 1.99)
	if err != nil {
		t.Error(err)
	}

	err = repo.Add(newProduct)
	if err != nil {
		return
	}
	if len(repo.products) != 1 {
		t.Errorf("Expected 1 newProduct, got %d", len(repo.products))
	}
}
func TestMemoryProductRepository_Get(t *testing.T) {
	repo := New()
	existingProd, err := product.NewProduct("Beer", "Good for you're health", 1.99)
	if err != nil {
		t.Error(err)
	}

	err = repo.Add(existingProd)
	if err != nil {
		return
	}
	if len(repo.products) != 1 {
		t.Errorf("Expected 1 product, got %d", len(repo.products))
	}

	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Get product by id",
			id:          existingProd.GetID(),
			expectedErr: nil,
		}, {
			name:        "Get non-existing product by id",
			id:          uuid.New(),
			expectedErr: product.ErrProductNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.GetByID(tc.id)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

		})
	}

}
func TestMemoryProductRepository_Delete(t *testing.T) {
	repo := New()
	existingProd, err := product.NewProduct("Beer", "Good for you're health", 1.99)
	if err != nil {
		t.Error(err)
	}

	err = repo.Add(existingProd)
	if err != nil {
		return
	}
	if len(repo.products) != 1 {
		t.Errorf("Expected 1 product, got %d", len(repo.products))
	}

	err = repo.Delete(existingProd.GetID())
	if err != nil {
		t.Error(err)
	}
	if len(repo.products) != 0 {
		t.Errorf("Expected 0 products, got %d", len(repo.products))
	}
}
