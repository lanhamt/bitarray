package bitarray

import (
    "testing"
)

func TestBitArrayInit(t *testing.T) {
    ba := BitArray(28)
    if ba.GetBit(0) {
        t.Error("Expected false, got true")
    }
    if ba.GetBit(5) {
        t.Error("Expected false, got true")
    }
    if ba.GetBit(27) {
        t.Error("Expected false, got true")
    }
}

func TestBitArrayGetAndSet(t *testing.T) {
    ba1 := BitArray(128)
    ba2 := BitArray(65537)

    ba1.SetBit(0, 1)
    ba1.SetBit(127, 1)
    if !ba1.GetBit(0) {
        t.Error("Expected bit 0 to be set to 1")
    }
    if !ba1.GetBit(127) {
        t.Error("Expected bit 127 to be set to 1")
    }

    ba1.SetBit(127, 0)
    if ba1.GetBit(127) {
        t.Error("Expected bit 127 to be set to 0")
    }

    ba2.SetBit(65530, 1)
    if !ba2.GetBit(65530) {
        t.Error("Expected bit 65530 to be set to 1")
    }
}