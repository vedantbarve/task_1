package helpers

import (
	"encoding/csv"
	"log"
	"os"
	"slices"
	"strings"
)

func GetApprovers(params string) [][]string {
	ip := strings.Split(params, " ")

	csvPath := ip[0]
	filters := [][]string{}

	for i := 1; i < len(ip); i++ {
		data := strings.Split(ip[i], "=")
		key, value := data[0], data[1]
		filters = append(filters, []string{key, value})
	}

	file, err := os.Open(csvPath)
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	// map headers to indexes
	headers := records[0]

	// result to be returned
	res := [][]string{}
	for i := 1; i < len(records); i++ {
		validRecord := true
		record := records[i]
		for _, filter := range filters {
			if record[slices.Index(headers, filter[0])] != filter[1] {
				validRecord = false
				break
			}
		}

		if validRecord {
			res = append(res, record)
		}
	}

	return res

}
