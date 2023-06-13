package tool

import (
    "encoding/binary"
    "fmt"
    "strings"
)

const (
    LITTLE_ENDIAN = 0
    BIG_ENDIAN    = 1
    BINARY        = 2
    OCTAL         = 8
    DECIMAL       = 10
    HEXADECIMAL   = 16
)

// Convert uint32 value to ram arrangement. Using litter endian.
// @param value: uint32 value
// @return string of ram arrangement
func ramLayoutUint32(value uint32, endian int) string {
    // convert uint32 values to byte arrangement(using litter endian)
    data := make([]byte, 4)
    if endian == LITTLE_ENDIAN {
        binary.LittleEndian.PutUint32(data, value)
    } else if endian == BIG_ENDIAN {
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

// Integer Conversion following base
// @value integer number value
// @base base number
// @zeroPadding whether padding using zero
// @return string of after converted
func baseConvertUint32(value uint32, base int, zeroPadding bool) string {
    if base == BINARY {
        if zeroPadding {
            return fmt.Sprintf("%032b", value)
        } else {
            return fmt.Sprintf("%b", value)
        }
    } else if base == OCTAL {
        if zeroPadding {
            return fmt.Sprintf("%011o", value)
        } else {
            return fmt.Sprintf("%o", value)
        }
    } else if base == DECIMAL {
        if zeroPadding {
            return fmt.Sprintf("%010d", value)
        } else {
            return fmt.Sprintf("%d", value)
        }
    } else if base == HEXADECIMAL {
        if zeroPadding {
            return fmt.Sprintf("%08X", value)
        } else {
            return fmt.Sprintf("%X", value)
        }
    }
    return ""
}

// calculate complement code. invert and add one
// @v value
// @base base
// @return string of complement code
func neg32(v int32, base int) string {
    bits := binary.Size(v) * 8
    if v >= 0 {
        // the complememt of a positive number is equal to the original code
        return baseConvertUint32(uint32(v), base, true)
    } else {
        // 1 << uint(32) set the highest to 1 and the rest to 0.
        // 100000000000000000000000000000000 （1 + 32 zeros） The maximum representation range.
        // The complement of negative numbers = maximum representation range - the positive value
        // The complement code = maximum representaion range + value.
        complementCode := (1 << bits) + v // 按位取反再加 1 得到补码, 最大表示范围-x
        return baseConvertUint32(uint32(complementCode), base, true)
    }
}

// 给内存值，计算真值
