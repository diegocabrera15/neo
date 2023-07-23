package types

type UnitDynamodb struct {
	UnitId      string `json:"unitId"`
	Name        string `json:"name"`
	unit        string `json:"unit"`
	description int    `json:"description"`
	state       string `json:"state"`
}
