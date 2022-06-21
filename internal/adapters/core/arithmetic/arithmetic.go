package arithmetic

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (arth Adapter) Add(a int32, b int32) (int32, error) {
	return a + b, nil
}
func (arth Adapter) Subtract(a int32, b int32) (int32, error) {
	return a - b, nil
}
func (arth Adapter) Multiply(a int32, b int32) (int32, error) {
	return a * b, nil
}
func (arth Adapter) Divide(a int32, b int32) (int32, error) {
	return a / b, nil
}
