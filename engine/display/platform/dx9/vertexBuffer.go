package dx9

import ()

//	if len(this.geometryResource.Geometrie.Vertex.Slot0) > 0 {
//		this.VertexBuffer[0], err = context.CreateVertexBuffer(uint(len(this.geometryResource.Geometrie.Vertex.Slot0)*4), d3d9.USAGE_WRITEONLY, 0, d3d9.POOL_MANAGED, nil)
//		if err == nil {
//			data, err := this.VertexBuffer[0].Lock(0, 0, d3d9.LOCK_DISCARD)
//			if err == nil {
//				data.SetFloat32s(0, this.geometryResource.Geometrie.Vertex.Slot0)
//			} else {
//				fmt.Println(err)
//			}
//			this.VertexBuffer[0].Unlock()
//		} else {
//			fmt.Println(err)
//		}
//	}
