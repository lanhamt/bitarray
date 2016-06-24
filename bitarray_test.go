/*
 * File: bitarray_test
 * ------------------------------
 * Automated unit testing for bitarray. 
 *
 * Note: does not provide coverage for 
 * invalid position for set or get, and 
 * invalid value for set since testing 
 * does not support panic. 
 */

package bitarray

import (
    "testing"
)

/*
 * Test: TestBitArrayInit
 * ------------------------------
 * Tests that all bits are initialized to 
 * 0 in array. 
 */
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

/*
 * Test: TestBitArrayGetAndSet
 * ------------------------------
 * Tests get and set functions for bit 
 * array on arrays of different sizes 
 * including one aligned to whole number of 
 * bytes and another not. 
 */
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