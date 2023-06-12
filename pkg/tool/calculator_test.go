package tool

import (
    "fmt"
    "testing"
)

func TestUint32ToRamLittleEndian(t *testing.T) {
    fmt.Printf("uint32ToRamLittleEndian(999): %v\n", uint32RamLayout(999, LITTLE_ENDIAN))
    fmt.Printf("uint32Convert(999, DECIMAL): %v\n", uint32Convert(0o1111, BINARY, false))
}
