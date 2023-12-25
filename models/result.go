package models

// CalculationResult represents the result of a calculation.
type CalculationResult struct {
	Operation string  `json:"operation"`
	Result    float64 `json:"result"`
}
