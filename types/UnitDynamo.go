package types

type UnitDynamodb struct {
	UnitId      string `json:"unitId"`
	Name        string `json:"name"`
	Unit        string `json:"unit"`
	Description string `json:"description"`
	State       string `json:"state"`
}
