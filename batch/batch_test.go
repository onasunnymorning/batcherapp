package batch

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewBatch(t *testing.T) {
	name := "Test Batch"
	b := NewBatch(name)

	if b.Name != name {
		t.Errorf("Expected batch name %q, but got %q", name, b.Name)
	}

	if b.AptForProduction {
		t.Errorf("Expected AptForProduction to be false, but got true")
	}

	if b.Id == uuid.Nil {
		t.Errorf("Expected batch id to not be zero, but got zero")
	}
}

func TestAddIngredient(t *testing.T) {
	b := NewBatch("Test Batch")
	ingredient := BatchIngredient{Name: "Test Ingredient", IsCacao: true, Qty: 50}
	b.AddIngredient(ingredient)

	if len(b.Ingredients) != 1 {
		t.Errorf("Expected 1 ingredient, but got %d", len(b.Ingredients))
	}

	if b.CacaoPercentage != 100 {
		t.Errorf("Expected cacao percentage to be 100, but got %f", b.CacaoPercentage)
	}

	if b.Output != ingredient.Qty {
		t.Errorf("Expected total quantity to be %d, but got %d", ingredient.Qty, b.Output)
	}

	if len(b.History) != 1 {
		t.Errorf("Expected 1 history event, but got %d", len(b.History))
	}

	event := b.History[0]
	if event.Action != "add_ingredient" {
		t.Errorf("Expected history event action to be 'add_ingredient', but got %q", event.Action)
	}

	if event.Message != ingredient.Name {
		t.Errorf("Expected history event message to be %q, but got %q", ingredient.Name, event.Message)
	}
}

func TestMarkAptForProduction(t *testing.T) {
	b := NewBatch("Test Batch")
	b.MarkAptForProduction()

	if !b.AptForProduction {
		t.Errorf("Expected AptForProduction to be true, but got false")
	}

	if len(b.History) != 1 {
		t.Errorf("Expected 1 history event, but got %d", len(b.History))
	}

	event := b.History[0]
	if event.Action != "marked_apt_for_production" {
		t.Errorf("Expected history event action to be 'marked_apt_for_production', but got %q", event.Action)
	}
}

func TestAddNotes(t *testing.T) {
	b := NewBatch("Test Batch")
	notes := "Test notes"
	b.AddNotes(notes)

	if b.Notes != notes {
		t.Errorf("Expected notes to be %q, but got %q", notes, b.Notes)
	}

	if len(b.History) != 1 {
		t.Errorf("Expected 1 history event, but got %d", len(b.History))
	}

	event := b.History[0]
	if event.Action != "add_notes" {
		t.Errorf("Expected history event action to be 'add_notes', but got %q", event.Action)
	}

	if event.Message != notes {
		t.Errorf("Expected history event message to be %q, but got %q", notes, event.Message)
	}
}
