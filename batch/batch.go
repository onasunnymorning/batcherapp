package batch

import (
	"time"

	"github.com/google/uuid"
)

type Batch struct {
	Id               uuid.UUID         `json:"id"`
	Name             string            `json:"name"`
	CacaoPercentage  float32           `json:"cacao_percentage"`
	Output           int               `json:"output"`
	Ingredients      []BatchIngredient `json:"ingredients"`
	History          []BatchEvent      `json:"history"`
	Notes            string            `json:"notes"`
	AptForProduction bool              `json:"apt_for_production"`
}

func NewBatch(name string) Batch {
	return Batch{
		Id:               uuid.New(),
		Name:             name,
		CacaoPercentage:  float32(0),
		Output:           0,
		Ingredients:      []BatchIngredient{},
		History:          []BatchEvent{},
		AptForProduction: false,
	}
}

func (b *Batch) AddIngredient(i BatchIngredient) {
	b.Ingredients = append(b.Ingredients, i)
	b.CacaoPercentage = b.calculateCacaoPct()
	b.Output = b.calculateTotalQty()
	b.History = append(b.History, BatchEvent{
		Ts:      time.Now(),
		Action:  "add_ingredient",
		Message: i.Name,
	})
}

func (b *Batch) MarkAptForProduction() {
	b.AptForProduction = true
	b.History = append(b.History, BatchEvent{
		Ts:      time.Now(),
		Action:  "marked_apt_for_production",
		Message: "",
	})
}

func (b *Batch) AddNotes(notes string) {
	b.Notes = notes
	b.History = append(b.History, BatchEvent{
		Ts:      time.Now(),
		Action:  "add_notes",
		Message: notes,
	})
}

func (b *Batch) calculateCacaoPct() float32 {
	var cacao int
	for _, i := range b.Ingredients {
		if i.IsCacao {
			cacao += i.Qty
		}
	}
	return float32(cacao) / float32(b.calculateTotalQty()) * 100
}

func (b *Batch) calculateTotalQty() int {
	var qty int
	for _, i := range b.Ingredients {
		qty += i.Qty
	}
	return qty
}

type BatchIngredient struct {
	Name    string `json:"name"`
	Qty     int    `json:"qty"`
	IsCacao bool   `json:"is_cacao"`
}

type BatchEvent struct {
	Ts      time.Time `json:"ts"`
	Action  string    `json:"action"`
	Message string    `json:"message"`
}
