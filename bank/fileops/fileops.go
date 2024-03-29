package fileops

import (
	"os"
	"fmt"
	"errors"
	"strconv"
)


func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 0, errors.New("Failed to find file.")
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 0, errors.New("Failed to pass value")
	}

	return value, nil
}