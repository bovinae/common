package collection

type BitMap []byte

const BYTE_SIZE = 8 //定义的bitmap为byte的数组，byte为8bit

func NewBitMap(n int64) BitMap {
	return make(BitMap, n/BYTE_SIZE+1)
}

func NewBitMapByIds(ids []int64) BitMap {
	if len(ids) == 0 {
		return NewBitMap(0)
	}
	var bm BitMap
	for _, id := range ids {
		bm.Set(id)
	}
	return bm
}

func (bm *BitMap) Set(n int64) {
	byteIndex := n / BYTE_SIZE // 第x个字节（0,1,2...）
	if byteIndex >= int64(len(*bm)) {
		tmp := make(BitMap, byteIndex-int64(len(*bm))+1)
		*bm = append(*bm, tmp...)
	}
	offsetIndex := n % BYTE_SIZE // 偏移量(0<偏移量<BYTE_SIZE)
	(*bm)[byteIndex] |= 1 << offsetIndex
}

func (bm BitMap) IsExist(n int64) bool {
	if n < 0 || n/BYTE_SIZE >= int64(len(bm)) {
		return false
	}
	byteIndex := n / BYTE_SIZE
	offsetIndex := n % BYTE_SIZE
	return bm[byteIndex]&(1<<offsetIndex) != 0
}

func (bm BitMap) Clean(n int64) {
	if n < 0 || n/BYTE_SIZE >= int64(len(bm)) {
		return
	}
	byteIndex := n / BYTE_SIZE
	offsetIndex := n % BYTE_SIZE
	bm[byteIndex] &= ^byte(1 << offsetIndex) //清零
}
