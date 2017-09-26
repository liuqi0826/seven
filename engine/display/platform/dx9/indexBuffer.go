package dx9

import (
//"github.com/gonutz/d3d9"
)

type IndexBuffer struct {
	Index uint32
	//IndexBuffer d3d9.IndexBuffer
}

func (this *IndexBuffer) Upload(data []uint16) error {
	var err error
	//	if len(this.geometryResource.Geometrie.Index) > 0 {
	//		this.IndexBuffer, err = context.CreateIndexBuffer(uint(len(this.geometryResource.Geometrie.Index)*2), d3d9.USAGE_WRITEONLY, d3d9.FMT_INDEX16, d3d9.POOL_MANAGED, nil)
	//		if err == nil {
	//			data, err := this.IndexBuffer.Lock(0, 0, d3d9.LOCK_DISCARD)
	//			if err == nil {
	//				data.SetUint16s(0, this.geometryResource.Geometrie.Index)
	//			} else {
	//				fmt.Println(err)
	//			}
	//			this.IndexBuffer.Unlock()
	//		} else {
	//			fmt.Println(err)
	//		}
	//	}
	return err
}
func (this *IndexBuffer) Dispose() {

}
