package tool

import (
    "fmt"
    "testing"
)

func TestUint32ToRamLittleEndian(t *testing.T) {
    // fmt.Printf("uint32ToRamLittleEndian(999): %v\n", ramLayoutUint32(52, LITTLE_ENDIAN))
    // fmt.Printf("uint32Convert(999, DECIMAL): %v\n", numConvertUint32(0o1111, BINARY, false))
    fmt.Printf("neg(100, BINARY, true): %v\n", neg32(-128, BINARY))
}
