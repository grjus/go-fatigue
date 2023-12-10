package stress

import (
	"errors"
)

func AlternatingStress(minStress, maxStress float64) (float64, error) {
	if minStress > maxStress {
		return 0, errors.New("Minimum stress should be higher than maximum one")
	}
	return 0.5 * (maxStress - minStress), nil
}

func MeanStress(minStress, maxStress float64) (float64, error) {
	if minStress > maxStress {
		return 0, errors.New("Minimum stress should be higher than maximum one")
	}
	meanStress := 0.5 * (maxStress + minStress)
	if meanStress < 0 {
		return 0, nil
	}
	return 0.5 * (maxStress + minStress), nil
}

func GoodmanStress(meanStress, alternatingStress, ultimateStrength float64) (float64, error) {
	if ultimateStrength <= 0 {
		return 0, errors.New("Invalid value for the ultimate strength")
	}
	return alternatingStress / (1 - meanStress/ultimateStrength), nil

}
