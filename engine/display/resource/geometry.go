package resource

import ()

type GeometryResource struct {
	ID        string             `json:"id"`
	Shader    string             `json:"shader"`
	Material  string             `json:"material"`
	Max       []float32          `json:"max"`
	Min       []float32          `json:"min"`
	Center    []float32          `json:"center"`
	Radius    float32            `json:"radius"`
	Skeleton  SkeletonComponent  `json:"skeleton"`
	Geometrie GeometrieComponent `json:"geometrie"`
}
type GeometrieComponent struct {
	Index  []uint16        `json:"index"`
	Vertex VertexComponent `json:"vertex"`
}
type SkeletonComponent struct {
	Struct        []interface{} `json:"struct"`
	Root          []float32     `json:"root"`
	Foot          []float32     `json:"foot"`
	BindShape     []float32     `json:"bindShape"`
	Invbind       [][]float32   `json:"invbind"`
	NodeNameList  []string      `json:"nodeNameList"`
	JointNameList []string      `json:"jointNameList"`
}
type VertexComponent struct {
	Slot0 []float32 `json:"0"`
	Slot1 []float32 `json:"1"`
	Slot2 []float32 `json:"2"`
	Slot3 []float32 `json:"3"`
	Slot4 []float32 `json:"4"`
	Slot5 []float32 `json:"5"`
	Slot6 []float32 `json:"6"`
	Slot7 []float32 `json:"7"`
}
