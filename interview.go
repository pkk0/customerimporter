// package customerimporter reads from the given customers.csv file and returns a
// sorted (data structure of your choice) of email domains along with the number
// of customers with e-mail addresses for each domain.  Any errors should be
// logged (or handled). Performance matters (this is only ~3k lines, but *could*
// be 1m lines or run on a small machine).
package customerimporter

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

type Domain struct {
	Name        string
	EmailsCount uint
}
type Domains []Domain

type SortType int

const (
	SORT_ASCEND SortType = iota
	SORT_DESCEND
)

// Don't need pointers down here.
// Slices does not store any data and are like references to arrays.
func (domains Domains) sortAscend() {
	sort.Slice(domains, func(i, j int) bool {
		return domains[i].EmailsCount < domains[j].EmailsCount
	})
}
func (domains Domains) sortDescend() {
	sort.Slice(domains, func(i, j int) bool {
		return domains[i].EmailsCount > domains[j].EmailsCount
	})
}

// Read only selected column from given CSV file and return it.
func readCsvFileColumn(filePath string, columnNumber uint) (columnData []string, err error) {
	file, openFileError := os.Open(filePath)

	if openFileError != nil {
		return nil, openFileError
	}

	// Make sure to close file when open was successful.
	defer file.Close()

	reader := csv.NewReader(file)
	columnData = make([]string, 0)
	header := true

	for {
		// Read the CSV file line by line to not fill up all the RAM.
		line, readError := reader.Read()

		if readError == io.EOF {
			break
		}

		// Skip header
		if header {
			header = false
			continue
		}

		if readError != nil {
			return nil, openFileError
		}

		columnData = append(columnData, line[columnNumber])
	}

	return columnData, nil
}

// Count given emails by domain and sort them ascend/descend.
func processEmails(emails []string, sortType SortType) Domains {
	tmpMap := make(map[string]uint)

	for _, email := range emails {
		splittedEmail := strings.Split(email, "@")

		if len(splittedEmail) != 2 {
			log.Printf("processEmails(emails []string, sortType SortType) Found invalid email '%v', skipping...\n", email)
			continue
		}

		tmpMap[splittedEmail[1]]++
	}

	i, slice := 0, make(Domains, len(tmpMap))
	for domain, count := range tmpMap {
		slice[i] = Domain{domain, count}
		i++
	}

	if sortType == SORT_ASCEND {
		slice.sortAscend()
	} else {
		slice.sortDescend()
	}

	return slice
}

// Returns Domains struct sorted ascend(SORT_ASCEND)/descend(SORT_DESCEND) by emails count.
// CSV file from which the data will be read must be provided as filePath.
func CountEmailsByDomain(filePath string, sortType SortType) Domains {
	emails, readError := readCsvFileColumn(filePath, 2)

	if readError != nil {
		log.Fatalf("Fatal error: %v\n", readError)
	}

	return processEmails(emails, sortType)
}
