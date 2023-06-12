package tool

import (
    "encoding/binary"
    "fmt"
    "strings"
)

const LittleEndian = 0
const BigEndian = 1

// Convert uint32 value to ram arrangement. Using litter endian.
// @param value: uint32 value
// @return string of ram arrangement
func uint32RamLayout(value uint32, endian int) string {
    // convert uint32 values to byte arrangement(using litter endian)
    data := make([]byte, 4)
    if endian == LittleEndian {
        binary.LittleEndian.PutUint32(data, value)
    } else if endian == BigEndian {
        binary.BigEndian.PutUint32(data, value)
    }

    // ram arrangement
    var ram string
    for _, d := range data {
        ram += fmt.Sprintf("%02X ", d)
    }
    ram = strings.TrimSpace(ram)
    return ram
}
