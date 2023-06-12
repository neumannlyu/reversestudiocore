package tool

import (
    "fmt"
    "testing"
)

func TestUint32ToRamLittleEndian(t *testing.T) {
    fmt.Printf("uint32ToRamLittleEndian(999): %v\n", uint32RamLayout(999, BigEndian))
}
