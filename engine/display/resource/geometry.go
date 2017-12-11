package resource

import (
	"encoding/json"
)

type GeometryResource struct {
	ID        string              `json:"id"`
	Shader    string              `json:"shader"`
	Material  string              `json:"material"`
	Max       []float32           `json:"max"`
	Min       []float32           `json:"min"`
	Center    []float32           `json:"center"`
	Radius    float32             `json:"radius"`
	Skined    bool                `json:"skined"`
	Geometrie *GeometrieComponent `json:"geometrie"`
}
type GeometrieComponent struct {
	Index  []uint16     `json:"index"`
	Vertex [8][]float32 `json:"vertex"`
}

func (this *GeometryResource) Parser(value []byte) error {
	var err error
	err = json.Unmarshal(value, this)
	return err
}
