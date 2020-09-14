package items

import (
	"api/app/mock"
	"net/http"
	"net/http/httptest"
	"testing"

	"api/app/models"

	"github.com/gin-gonic/gin"
)

func TestHandler(t *testing.T) {
	router := gin.Default()
	Configure(router, nil)

	// Inject our mock into our handler.
	var is mock.ItemService
	Is = &is

	// Mock our User() call.
	is.ItemFn = func(id int) (*models.Item, error) {
		if id != 100 {
			t.Fatalf("unexpected id: %d", id)
		}
		return &models.Item{ID: "100", Name: "DaItam", Description: "Elnesto"}, nil
	}
	
	is.ItemsFn = func() ([]*models.Item, error) {
		var items []*models.Item
		items = append(items, &models.Item{ID: "100", Name: "DaItam", Description: "Elnesto"})
		return items, nil
	}

	is.CreateItemFn = func(i *models.Item) error {
		return nil
	}

	is.DeleteItemFn = func(id int) error {
		t.Fatalf("unexpected id: %d", id)
		return nil
	}

	// Invoke the handler.
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/item/100", nil)
	r2, _ := http.NewRequest("DELETE", "/item/100", nil)
	router.ServeHTTP(w, r)
	router.ServeHTTP(w, r2)

	// Validate mock.
	if !is.ItemInvoked {
		t.Fatal("expected User() to be invoked")
	}
}
