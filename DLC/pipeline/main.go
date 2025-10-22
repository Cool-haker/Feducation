// Вы разрабатываете систему для обработки финансовых транзакций.
// Каждая транзакция проходит несколько этапов обработки:
// 1. Чтение транзакций из исходных данных.
// 2. Фильтрация транзакций: убирает транзакции с отрицательными суммами.
// 3. Конвертация валюты: преобразуем сумму транзакции в доллары.
// 4. Сохранение результатов: записываем обработанные транзакции в итоговый список.

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	transaction := generaredTransaction(10)
	filtered := filteredTransaction(transaction)
	convert := convertTransaction(filtered)
	saveTransaction(convert)
}

type Transaction struct {
	ID     int64
	Amount float64
}

func generaredTransaction(count int) <-chan Transaction {
	out := make(chan Transaction, count)

	go func() {
		for i := 0; i < count; i++ {
			out <- Transaction{
				ID:     int64(i),
				Amount: rand.Float64()*200 - 100,
			}
		}

		close(out)
	}()

	return out
}

func filteredTransaction(in <-chan Transaction) chan Transaction {
	out := make(chan Transaction)

	go func() {
		for tr := range in {
			if tr.Amount >= 0 {
				out <- tr
			}
		}

		close(out)
	}()

	return out
}

func convertTransaction(in <-chan Transaction) chan Transaction {
	out := make(chan Transaction)

	go func() {
		for tr := range in {
			tr.Amount *= 0.8
			out <- tr
		}

		close(out)
	}()

	return out
}

func saveTransaction(in chan Transaction) {
	for tr := range in {
		fmt.Printf("Transaction ID: %d Amount: %.2f\n", tr.ID, tr.Amount)
	}
}
