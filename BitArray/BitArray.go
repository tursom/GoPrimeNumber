package BitArray

type Type struct {
	_defaultState bool
	_array        []uint64
}

func bitMapNeedSize(size uint64) uint64 {
	return ((size-1)>>6 + 1) & 0xffffffffffffff
}

func BitArray(size uint64, defaultState bool) Type {
	return Type{
		defaultState,
		make([]uint64, bitMapNeedSize(size)),
	}
}

func (b *Type) Resize(size uint64) bool {
	newSize := int(bitMapNeedSize(size))
	if newSize <= len(b._array) {
		return false
	}
	newArray := make([]uint64, newSize)
	for i := 0; i < len(b._array); i++ {
		newArray[i] = b._array[i]
	}
	var defaultValue uint64
	if b._defaultState {
		defaultValue = ^uint64(0)
	} else {
		defaultValue = 0
	}
	for i := len(b._array); i < newSize; i++ {
		newArray[i] = defaultValue
		//_ = append(b._array, defaultValue)
	}
	b._array = newArray
	return true
}

func (b *Type) Get(index uint64) bool {
	arrayIndex := index >> 6
	if arrayIndex >= uint64(len(b._array)) {
		return b._defaultState
	} else {
		b.Resize(index)
		return b._array[arrayIndex]&(uint64(1)<<(index&63)) != 0
	}
}

func (b *Type) Up(index uint64) {
	b.Resize(index)
	b._array[index>>6] |= uint64(1) << (index & 63)
}

func (b *Type) Down(index uint64) {
	b.Resize(index)
	b._array[index>>6] &= ^(uint64(1) << (index & 63))
}

func (b *Type) Size() uint64 {
	return uint64(len(b._array)) << 6
}
