package fabrikant

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"github.com/rogudator/tender/gen_err"
)

func GetPurchases(item string) ([]string, []string, error) {

	url := fmt.Sprintf("https://etp-ets.ru/44/catalog/procedure?q=%s", item)
	res, err := http.Get(url)
	if err != nil {
		return nil, nil, err
	}
	return ParsePurchases(res.Body)
}

func ParsePurchases(r io.Reader) ([]string, []string, error) {
	// Решил передавать два массива вместо структуры, потому что
	// показалось, что если еще создавать массив структур с двумя
	// полями, и заполнять его, это будет еще один O(n) и будет медленнее
	purchaseID := make([]string, 0)
	purchaseName := make([]string, 0)

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	// Пакет bluemonday позволяет удалить все теги, оставив только
	// необходимое значение для задачи
	p := bluemonday.StrictPolicy()

	for scanner.Scan() {
		t := scanner.Text()
		if strings.Contains(t, "row-procedure_number") {
			purchaseID = append(purchaseID, strings.TrimSpace(p.Sanitize(t)))
		}
		if strings.Contains(t, "row-procedure_name") {
			index := strings.Index(t, "</a>")
			purchaseName = append(purchaseName, strings.TrimSpace(p.Sanitize(t[:index])))
		}
	}
	if len(purchaseID) == 0 || len(purchaseName) == 0 {
		return nil, nil, gen_err.FailedToParseTenders
	}

	return purchaseID, purchaseName, nil
}
