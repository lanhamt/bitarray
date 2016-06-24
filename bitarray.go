package bitarray

const bitsPerByte = 8

type BArray struct {
    numBits uint32
    a[]     byte
}

func BitArray(length uint32) *BArray {
    arr := new(BArray)
    arr.numBits = length
    size := 1 + ((length -1) / bitsPerByte)
    arr.a = make([] byte, size)
    return arr
}

func (arr *BArray) SetBit(position uint32, value byte) {
    if position >= arr.numBits {
        panic("set: out of bounds")
    }

    nByte := position / bitsPerByte
    nBit := position % bitsPerByte

    if value == 1 {
        mask := byte(0x1 << nBit)
        arr.a[nByte] |= mask
    } else {
        mask := ^byte(0x1 << nBit)
        arr.a[nByte] &= mask
    }
}

func (arr *BArray) GetBit(position uint32) bool {
    if position >= arr.numBits {
        panic("get: out of bounds")
    }

    nByte := position / bitsPerByte
    nBit := position % bitsPerByte

    return (arr.a[nByte] & (0x1 << nBit)) != 0
}

func (arr *BArray) GetLength() uint32 {
    return arr.numBits
}