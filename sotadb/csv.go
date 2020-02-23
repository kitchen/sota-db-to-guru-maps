package sotadb

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"

	"github.com/gocarina/gocsv"
)

func ParseCSVFile(filename string) ([]*Summit, error) {
	summitsFile, err := os.Open(filename)
	if err != nil {
		return []*Summit{}, err
	}
	defer summitsFile.Close()

	// chomp first line from file
	br := bufio.NewReader(summitsFile)
	_, err = br.ReadBytes('\n')
	if err != nil {
		return []*Summit{}, err
	}

	return ParseCSV(br)
}

func ParseCSV(r io.Reader) ([]*Summit, error) {
	cr := csv.NewReader(r)
	cr.FieldsPerRecord = -1 // don't error with inconsistent field numbers

	summits := []*Summit{}
	err := gocsv.UnmarshalCSV(cr, &summits)
	if err != nil {
		return []*Summit{}, err
	}

	return summits, nil
}
