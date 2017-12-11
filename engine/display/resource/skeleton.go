package resource

import (
	"encoding/json"
)

type SkeletonResource struct {
	ID            string        `json:"id"`
	Struct        []interface{} `json:"struct"`
	Root          []float32     `json:"root"`
	Foot          []float32     `json:"foot"`
	BindShape     []float32     `json:"bindShape"`
	Invbind       [][]float32   `json:"invbind"`
	NodeNameList  []string      `json:"nodeNameList"`
	JointNameList []string      `json:"jointNameList"`
}

func (this *SkeletonResource) Parser(value []byte) error {
	var err error
	err = json.Unmarshal(value, this)
	return err
}
