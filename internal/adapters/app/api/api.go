package api

import (
	"pent/internal/ports"
)

type Adapter struct {
	db    ports.DbPort
	arith ports.ArithmeticPort
}

func NewAdapter(arith ports.ArithmeticPort, db ports.DbPort) *Adapter {
	return &Adapter{arith: arith, db: db}
}

func (apiA Adapter) GetAddition(a, b int32) (int32, error) {
	answer, err := apiA.arith.Add(a, b)
	if err != nil {
		return 0, err
	}

	err = apiA.db.AddToHistory(answer, "Addition")

	if err != nil {
		return 0, err
	}

	return answer, nil
}
func (apiA Adapter) GetSubtraction(a, b int32) (int32, error) {
	answer, err := apiA.arith.Subtract(a, b)
	if err != nil {
		return 0, err
	}

	err = apiA.db.AddToHistory(answer, "Subtraction")

	if err != nil {
		return 0, err
	}

	return answer, nil
}
func (apiA Adapter) GetMultiplication(a, b int32) (int32, error) {
	answer, err := apiA.arith.Multiply(a, b)
	if err != nil {
		return 0, err
	}

	err = apiA.db.AddToHistory(answer, "Multiplication")

	if err != nil {
		return 0, err
	}

	return answer, nil
}
func (apiA Adapter) GetDivision(a, b int32) (int32, error) {
	answer, err := apiA.arith.Divide(a, b)
	if err != nil {
		return 0, err
	}

	err = apiA.db.AddToHistory(answer, "Division")

	if err != nil {
		return 0, err
	}

	return answer, nil
}
