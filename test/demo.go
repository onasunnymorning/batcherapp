package app

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/onasunnymorning/batcherapp/batch"
)

func RunRepositoryDemo(ctx context.Context, batchRepository batch.Repository) {
	fmt.Println("1. MIGRATE REPOSITORY")

	if err := batchRepository.Migrate(ctx); err != nil {
		log.Fatal(err)
	}

	fmt.Println("2. CREATE RECORDS OF REPOSITORY")
	batch1 := batch.NewBatch("Masha-" + time.Now().String())
	batch2 := batch.NewBatch("Geoff-" + time.Now().String())

	createdBatch1, err := batchRepository.Create(ctx, batch1)
	if err != nil {
		fmt.Println(err)
	}

	createdBatch2, err := batchRepository.Create(ctx, batch2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n%+v\n", createdBatch1, createdBatch2)

	fmt.Println("3. GET BY ID")

	createdBatch1, err = batchRepository.GetByID(ctx, createdBatch1.Id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(createdBatch1)

	fmt.Println("4. GET BY NAME")

	createdBatch2, err = batchRepository.GetByName(ctx, createdBatch2.Name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(createdBatch1)

	fmt.Println("5. UPDATE")

	createdBatch2.Notes = "This batch was updated"
	updatedBatch, err := batchRepository.Update(ctx, *createdBatch2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(updatedBatch)

	fmt.Println("6. DELETE")

	err = batchRepository.DeleteByID(ctx, createdBatch1.Id)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("7. ADD ingredient")
	createdBatch2.AddIngredient(batch.BatchIngredient{
		Name:    "Cacao Butter",
		Qty:     100,
		IsCacao: true,
	})
	updatedBatch2, err := batchRepository.Update(ctx, *createdBatch2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(updatedBatch2)

	fmt.Println("8. ADD ANOTHER ingredient")
	createdBatch2.AddIngredient(batch.BatchIngredient{
		Name:    "Cacao Nibs",
		Qty:     900,
		IsCacao: true,
	})
	updatedBatch2, err = batchRepository.Update(ctx, *createdBatch2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(updatedBatch2)

	fmt.Println("9. LIST")
	list, err := batchRepository.List(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(list)
}
