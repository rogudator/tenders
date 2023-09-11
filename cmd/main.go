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
	// Так как каждый новое ключевой слово начинается с новой
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
	fmt.Println(item)
	for i := range purchasesID {
		fmt.Println(purchasesID[i])
		fmt.Println(purchasesName[i])
	}
}