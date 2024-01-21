package kernel

type TArray struct {
	Data  uintptr
	Count int32
	Max   int32
}

func (arr *TArray) ReplaceData(newData uintptr) {
	arr.Data = newData
}

func (arr *TArray) ReadAtIndex(index int, d *Driver) uintptr {
	return d.Read(arr.Data + (uintptr(index) * 8))
}

func (arr *TArray) ReadAtIndex2(index int, d *Driver) uintptr {
	return d.Read(arr.Data + (uintptr(index) * 8))
}
