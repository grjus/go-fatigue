package algebra

import "errors"

func Linspace(start, stop float64, n int) ([]float64, error) {
	if n <= 0 {
		return nil, errors.New("Linspace range should be greater than 0")
	}

	delta := (stop - start) / float64(n-1)
	result := make([]float64, n)

	for i := range result {
		result[i] = start + delta*float64(i)
	}

	return result, nil
}
