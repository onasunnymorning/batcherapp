package batch

import (
	"context"
	json "encoding/json"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type gormBatch struct {
	Id               uuid.UUID      `gorm:"type:uuid;primary_key;"`
	Name             string         `gorm:"type:string;not null;unique;"`
	CacaoPercentage  float32        `gorm:"type:float;"`
	Output           int            `gorm:"type:int;"`
	Ingredients      datatypes.JSON `gorm:"type:json;"`
	History          datatypes.JSON `gorm:"type:json;"`
	Notes            string         `gorm:"type:text;"`
	AptForProduction bool           `gorm:"type:boolean;"`
}

func (gormBatch) TableName() string {
	return "batches"
}

type PostgreSQLGORMRepository struct {
	db *gorm.DB
}

func NewPostgreSQLGORMRepository(db *gorm.DB) *PostgreSQLGORMRepository {
	return &PostgreSQLGORMRepository{
		db: db,
	}
}

func (r *PostgreSQLGORMRepository) Migrate(ctx context.Context) error {
	return r.db.WithContext(ctx).AutoMigrate(&gormBatch{})
}

func (r *PostgreSQLGORMRepository) Create(ctx context.Context, batch Batch) (*Batch, error) {
	ingredients, err := json.Marshal(&batch.Ingredients)
	if err != nil {
		return nil, err
	}

	history, err := json.Marshal(batch.History)
	if err != nil {
		return nil, err
	}

	gormBatch := gormBatch{
		Id:               batch.Id,
		Name:             batch.Name,
		CacaoPercentage:  batch.CacaoPercentage,
		Output:           batch.Output,
		Ingredients:      ingredients,
		History:          history,
		Notes:            batch.Notes,
		AptForProduction: false,
	}

	result := r.db.Create(&gormBatch)
	if result.Error != nil {
		return nil, result.Error
	}

	return &batch, nil
}

func (r *PostgreSQLGORMRepository) GetByID(ctx context.Context, id uuid.UUID) (*Batch, error) {
	var gormBatch gormBatch
	result := r.db.WithContext(ctx).First(&gormBatch, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	fmt.Println(gormBatch.Ingredients)
	var ingredients []BatchIngredient
	err := json.Unmarshal(gormBatch.Ingredients, &ingredients)
	if err != nil {
		return nil, err
	}

	fmt.Println(gormBatch.History)
	var history []BatchEvent
	err = json.Unmarshal(gormBatch.History, &history)
	if err != nil {
		return nil, err
	}

	batch := Batch{
		Id:               gormBatch.Id,
		Name:             gormBatch.Name,
		CacaoPercentage:  gormBatch.CacaoPercentage,
		Output:           gormBatch.Output,
		Notes:            gormBatch.Notes,
		AptForProduction: gormBatch.AptForProduction,
		Ingredients:      ingredients,
		History:          history,
	}

	return &batch, nil
}

func (r *PostgreSQLGORMRepository) GetByName(ctx context.Context, name string) (*Batch, error) {
	var gormBatch gormBatch
	result := r.db.WithContext(ctx).First(&gormBatch, "name = ?", name)
	if result.Error != nil {
		return nil, result.Error
	}

	var ingredients []BatchIngredient
	err := json.Unmarshal(gormBatch.Ingredients, &ingredients)
	if err != nil {
		return nil, err
	}

	var history []BatchEvent
	err = json.Unmarshal(gormBatch.History, &history)
	if err != nil {
		return nil, err
	}

	batch := Batch{
		Id:               gormBatch.Id,
		Name:             gormBatch.Name,
		CacaoPercentage:  gormBatch.CacaoPercentage,
		Output:           gormBatch.Output,
		Notes:            gormBatch.Notes,
		AptForProduction: gormBatch.AptForProduction,
		Ingredients:      ingredients,
		History:          history,
	}

	return &batch, nil
}

func (r *PostgreSQLGORMRepository) Update(ctx context.Context, batch Batch) (*Batch, error) {
	ingredients, err := json.Marshal(&batch.Ingredients)
	if err != nil {
		return nil, err
	}

	history, err := json.Marshal(batch.History)
	if err != nil {
		return nil, err
	}

	gormBatch := gormBatch{
		Id:               batch.Id,
		Name:             batch.Name,
		CacaoPercentage:  batch.CacaoPercentage,
		Output:           batch.Output,
		Ingredients:      ingredients,
		History:          history,
		Notes:            batch.Notes,
		AptForProduction: batch.AptForProduction,
	}

	result := r.db.WithContext(ctx).Save(&gormBatch)
	if result.Error != nil {
		return nil, result.Error
	}

	return &batch, nil
}

func (r *PostgreSQLGORMRepository) DeleteByID(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&gormBatch{}, "id = ?", id)
	return result.Error
}

func (r *PostgreSQLGORMRepository) List(ctx context.Context) ([]Batch, error) {
	var gormBatches []gormBatch
	result := r.db.WithContext(ctx).Find(&gormBatches)
	if result.Error != nil {
		return nil, result.Error
	}

	var batches []Batch
	for _, gormBatch := range gormBatches {
		var ingredients []BatchIngredient
		err := json.Unmarshal(gormBatch.Ingredients, &ingredients)
		if err != nil {
			return nil, err
		}

		var history []BatchEvent
		err = json.Unmarshal(gormBatch.History, &history)
		if err != nil {
			return nil, err
		}

		batch := Batch{
			Id:               gormBatch.Id,
			Name:             gormBatch.Name,
			CacaoPercentage:  gormBatch.CacaoPercentage,
			Output:           gormBatch.Output,
			Notes:            gormBatch.Notes,
			AptForProduction: gormBatch.AptForProduction,
			Ingredients:      ingredients,
			History:          history,
		}

		batches = append(batches, batch)
	}

	return batches, nil
}
