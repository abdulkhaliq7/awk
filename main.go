package awk

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Awk struct {
	Data string
}

func (a *Awk) DataSplit(splitField, printingField string, chosenColumns ...string) (*Awk, error) {

	if a.Data == "" {
		return &Awk{Data: ""}, nil
	}

	// Precompute the chosen column indices to avoid repeated string-to-int conversion
	columns := make([]int, len(chosenColumns))
	for i, col := range chosenColumns {
		index, err := strconv.Atoi(col)
		if err != nil {
			return nil, fmt.Errorf("column conversion error: %v", err)
		}
		columns[i] = index
	}

	var fieldsChosen strings.Builder

	scanner := bufio.NewScanner(strings.NewReader(a.Data))

	for scanner.Scan() {

		line := scanner.Text()
		splittedData := strings.Split(line, splitField)

		if len(columns) == 0 {
			for _, all := range splittedData {
				fieldsChosen.WriteString(all)
				fieldsChosen.WriteString(printingField)
			}
			newLine := "\n"
			fieldsChosen.WriteString(newLine)

		} else {
			for _, column := range columns {
				if column >= 0 && column < len(splittedData) {
					fieldsChosen.WriteString(splittedData[column])
					fieldsChosen.WriteString(printingField)
				} else {
					return nil, fmt.Errorf("chosen column %v is out of bounds (line length: %v)", column, len(splittedData))
				}
			}
			newLine := "\n"
			fieldsChosen.WriteString(newLine)
		}
	}
	return &Awk{Data: fieldsChosen.String()}, nil
}

func (a *Awk) Replace(old, new string) *Awk {

	updatedData := strings.ReplaceAll(a.Data, old, new)
	return &Awk{Data: updatedData}
}

func NewAwk(data string) *Awk { return &Awk{Data: data} }
