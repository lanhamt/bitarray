package bitarray

// well-known type implementation for byte
const bitsPerByte = 8

/*
 * Struct: BArray 
 * ------------------------------
 * Struct for bit array. numBits keeps track of 
 * size of array. a is array of bytes used for 
 * internal storage. 
 */
type BArray struct {
    numBits uint32
    a[]     byte
}

/*
 * Constructor: BitArray
 * ------------------------------
 * Constructs bit array and returns pointer
 * to allocated struct, Uses length given to 
 * determine allocation size (on 0 length returns 
 * null pointer). 
 */
func BitArray(length uint32) *BArray {
    if length == 0 {
        return nil
    }
    arr := new(BArray)
    arr.numBits = length
    size := 1 + ((length -1) / bitsPerByte) // ceil for number of bytes in array
    arr.a = make([] byte, size)
    return arr
}

/*
 * Method: SetBit
 * ------------------------------
 * Sets bit in array at given position to given 
 * value. Panics on invalid position or invalid 
 * value. Retrieves appropriate byte from array and 
 * sets appropriate bit to given value. 
 */
func (arr *BArray) SetBit(position uint32, value byte) {
    if position >= arr.numBits {
        panic("set: out of bounds")
    }
    if value != 0 && value != 1 {
        panic("set: invalid value to set")
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

/*
 * Method: GetBit
 * ------------------------------
 * Gets value of bit at given position, returned 
 * as bool. Panics on invalid position. Retrieves 
 * appropriate byte in array and returns true if 
 * bit is 1, false if 0. 
 */
func (arr *BArray) GetBit(position uint32) bool {
    if position >= arr.numBits {
        panic("get: out of bounds")
    }

    nByte := position / bitsPerByte
    nBit := position % bitsPerByte

    return (arr.a[nByte] & (0x1 << nBit)) != 0
}

/*
 * Method: GetLength
 * ------------------------------
 * Returns number of bits in array. 
 */
func (arr *BArray) GetLength() uint32 {
    return arr.numBits
}