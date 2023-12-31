package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/rogudator/tender/pkg/fabrikant"
)

func main() {
	// Открываем файл со списком ключевых слов закупок
	list, err := os.Open("list.txt")
	if err != nil {
		log.Fatal("failed to read list of tenders")
	}
	defer list.Close()
	scanner := bufio.NewScanner(list)
	scanner.Split(bufio.ScanLines)
	// Так как каждое новое ключевое слово начинается с новой
	// строки, мы построчно читаем из файла и выводим найденные тендеры
	for scanner.Scan() {
		item := scanner.Text()
		purchasesID, purchasesName, err := fabrikant.GetPurchases(item)
		if err != nil {
			log.Fatal(err)
		}
		printTendersOfItem(item, purchasesID, purchasesName)
	}

}

func printTendersOfItem(item string, purchasesID, purchasesName []string) {
	fmt.Printf("\nItem: %s\n\n", item)
	for i := range purchasesID {
		fmt.Printf("ID: %s\n", purchasesID[i])
		fmt.Printf("Name: %s\n\n", purchasesName[i])
	}
}
