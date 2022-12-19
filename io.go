package gostats

import (
	"bufio"
	"strings"
	"strconv"
)

func ReadFloats(s *bufio.Scanner) ([]float64, error) {
	data := make([]float64, 0, 1024)
	for s.Scan() {
		line := s.Text()
		for _, token := range strings.Fields(line) {
			if val, err := strconv.ParseFloat(token, 64); err != nil {
				return nil, err
			} else {
				data = append(data, val)
			}
		}

	}
	return data, nil
}
