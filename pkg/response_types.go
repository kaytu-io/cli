package pkg

type SpendTableRow struct {
	DimensionID   string             `json:"dimensionId" example:"compute"`
	AccountID     string             `json:"accountID" example:"1239042"`
	Connector     string             `json:"connector" example:"AWS"`
	Category      string             `json:"category" example:"Compute"`
	DimensionName string             `json:"dimensionName" example:"Compute"`
	CostValue     map[string]float64 `json:"costValue"`
}
