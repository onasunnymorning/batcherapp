package main

import (
	"fmt"

	"example.com/m/batch"
)

func main() {
	newBatch := batch.NewBatch("DarkChocolateBatch007")
	newBatch.AddIngredient(batch.BatchIngredient{
		Name:    "Sugar",
		Qty:     10,
		IsCacao: false,
	})
	newBatch.AddIngredient(batch.BatchIngredient{
		Name:    "Cacao Nibs",
		Qty:     90,
		IsCacao: true,
	})

	fmt.Println(newBatch.Name)

	fmt.Println(newBatch.CacaoPercentage)
	fmt.Println(newBatch.Output)

	newBatch.AddNotes("This is the best chocolate in the world")
	newBatch.MarkAptForProduction()

	fmt.Println(newBatch.History)

}
