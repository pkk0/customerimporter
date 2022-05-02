package customerimporter

import (
	"reflect"
	"testing"
)

func TestDomainsAscendSort(t *testing.T) {
	wanted := Domains{{"github.io", 1}, {"cyberchimps.com", 4}, {"gmail", 10}}

	domains := Domains{{"github.io", 1}, {"gmail", 10}, {"cyberchimps.com", 4}}
	domains.sortAscend()

	if !reflect.DeepEqual(domains, wanted) {
		t.Fatalf("domains.sortAscend() = %v, wanted = %v\n", domains, wanted)
	}
}

func TestDomainsDescendSort(t *testing.T) {
	wanted := Domains{{"gmail", 10}, {"cyberchimps.com", 4}, {"github.io", 1}}

	domains := Domains{{"gmail", 10}, {"github.io", 1}, {"cyberchimps.com", 4}}
	domains.sortDescend()

	if !reflect.DeepEqual(domains, wanted) {
		t.Fatalf("domains.sortDescend() = %v, wanted = %v\n", domains, wanted)
	}
}

func TestCsvFileColumnRead1(t *testing.T) {
	column, readError := readCsvFileColumn("test1.csv", 1)

	if readError != nil {
		t.Fatalf("Read CSV file error: %v", readError)
	}

	wanted := []string{"Ortiz", "Hernandez", "Ortiz"}
	if !reflect.DeepEqual(column, wanted) {
		t.Fatalf("readCsvFileColumn(\"test1.csva\", 1) = %v, wanted = %v\n", column, wanted)
	}
}

func TestCsvFileColumnRead2(t *testing.T) {
	column, readError := readCsvFileColumn("test2.csv", 1)

	if readError != nil {
		t.Fatalf("Read CSV file error: %v", readError)
	}

	wanted := []string{"Hernandez", "Ortiz", "last_name", "Ortiz", "Ortiz", "Ortiz"}
	if !reflect.DeepEqual(column, wanted) {
		t.Fatalf("readCsvFileColumn(\"test1.csva\", 1) = %v, wanted = %v\n", column, wanted)
	}
}

func TestCsvFileColumnRead3(t *testing.T) {
	column, readError := readCsvFileColumn("test1.csv", 2)

	if readError != nil {
		t.Fatalf("Read CSV file error: %v", readError)
	}

	wanted := []string{"a@cyberchimps.com", "mhernandez0@github.io", "bortiz1@cyberchimps.com"}
	if !reflect.DeepEqual(column, wanted) {
		t.Fatalf("readCsvFileColumn(\"test1.csva\", 2) = %v, wanted = %v\n", column, wanted)
	}
}

func TestCsvFileColumnRead4(t *testing.T) {
	column, readError := readCsvFileColumn("test2.csv", 2)

	if readError != nil {
		t.Fatalf("Read CSV file error: %v", readError)
	}

	wanted := []string{
		"mhernandez0@github.io", "a@cyberchimps.com", "email", "a@cyberchimps.com",
		"bortiz1@cyberchimps.com", "b@cyberchimps.com",
	}
	if !reflect.DeepEqual(column, wanted) {
		t.Fatalf("readCsvFileColumn(\"test1.csva\", 2) = %v, wanted = %v\n", column, wanted)
	}
}

func TestEmailsProcessing1(t *testing.T) {
	domains := processEmails([]string{"a@example.com", "b@cyberchimps.com", "b@example.com"}, SORT_ASCEND)

	wanted := Domains{{"cyberchimps.com", 1}, {"example.com", 2}}
	if !reflect.DeepEqual(domains, wanted) {
		t.Fatalf(`processEmails([]string{\"a@example.com", "b@cyberchimps.com", "b@example.com"}, SORT_ASCEND) = %v, wanted = %v\n"`,
			domains, wanted)
	}
}

func TestEmailsProcessing2(t *testing.T) {
	domains := processEmails([]string{"a@example.com", "email", "b@cyberchimps.com", "b@example.com"}, SORT_DESCEND)

	wanted := Domains{{"example.com", 2}, {"cyberchimps.com", 1}}
	if !reflect.DeepEqual(domains, wanted) {
		t.Fatalf(`processEmails([]string{\"a@example.com", "email", "b@cyberchimps.com", "b@example.com"}, SORT_DESCEND) = %v, wanted = %v\n"`,
			domains, wanted)
	}
}

func TestCountEmailsByDomainFromFile1(t *testing.T) {
	domains := CountEmailsByDomain("test1.csv", SORT_DESCEND)

	wanted := Domains{{"cyberchimps.com", 2}, {"github.io", 1}}
	if !reflect.DeepEqual(domains, wanted) {
		t.Fatalf(`CountEmailsByDomain("test1.csv", SORT_DESCEND) = %v, wanted = %v\n"`,
			domains, wanted)
	}
}

func TestCountEmailsByDomainFromFile2(t *testing.T) {
	domains := CountEmailsByDomain("test2.csv", SORT_ASCEND)

	wanted := Domains{{"github.io", 1}, {"cyberchimps.com", 4}}
	if !reflect.DeepEqual(domains, wanted) {
		t.Fatalf(`CountEmailsByDomain("test2.csv", SORT_ASCEND) = %v, wanted = %v\n"`,
			domains, wanted)
	}
}
