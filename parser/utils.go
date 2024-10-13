package parser

import "strconv"

func IsConstant(token string) (float64, error) {
	floatValue, err := strconv.ParseFloat(token, 64)
	if err != nil {
		return 0, err
	}

	return floatValue, nil
}