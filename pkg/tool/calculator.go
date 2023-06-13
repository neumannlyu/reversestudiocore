package tool

import (
    "encoding/binary"
    "errors"
    "fmt"
    "math"
    "strings"
)

const (
    LITTLE_ENDIAN  = 0
    BIG_ENDIAN     = 1
    BINARY         = 2
    OCTAL          = 8
    DECIMAL        = 10
    HEXADECIMAL    = 16
    UseBaseType    = "v is a base type like int float"
    UnknownVarSize = "size of var is not 32 or 64"
)

// Convert uint32 value to ram arrangement. Using litter endian.
// @param value: uint32 value
// @return string of ram arrangement
func ramLayout(v interface{}, endian int) (string, error) {
    bytes := binary.Size(v)
    if bytes < 0 {
        return "", errors.New(UseBaseType)
    }

    // convert uint32 values to byte arrangement(using litter endian)
    if bytes == 4 {
        data := make([]byte, 4)
        if endian == LITTLE_ENDIAN {
            binary.LittleEndian.PutUint32(data, v.(uint32))
        } else if endian == BIG_ENDIAN {
            binary.BigEndian.PutUint32(data, v.(uint32))
        }
        // ram arrangement
        var ram string
        for _, d := range data {
            ram += fmt.Sprintf("%02X ", d)
        }
        ram = strings.TrimSpace(ram)
        return ram, nil
    } else if bytes == 8 {
        data := make([]byte, 8)
        if endian == LITTLE_ENDIAN {
            binary.LittleEndian.PutUint64(data, v.(uint64))
        } else if endian == BIG_ENDIAN {
            binary.BigEndian.PutUint64(data, v.(uint64))
        }
        // ram arrangement
        var ram string
        for _, d := range data {
            ram += fmt.Sprintf("%02X ", d)
        }
        ram = strings.TrimSpace(ram)
        return ram, nil
    }

    return "", errors.New(UnknownVarSize)
}

// ram data(4 or 8 bytes) convert to float data
// @ram the data of ram
// @endian little or big endian
// @return interface{} float32 or float64
func ramData2Float(ram []byte, endian int) (interface{}, error) {
    if len(ram) != 4 && len(ram) != 8 {
        return -1, errors.New(UnknownVarSize)
    }
    if len(ram) == 4 {
        if endian == LITTLE_ENDIAN {
            return math.Float32frombits(binary.LittleEndian.Uint32(ram)), nil
        } else if endian == BIG_ENDIAN {
            return math.Float32frombits(binary.BigEndian.Uint32(ram)), nil
        }
    } else if len(ram) == 8 {
        if endian == LITTLE_ENDIAN {
            return math.Float64frombits(binary.LittleEndian.Uint64(ram)), nil
        } else if endian == BIG_ENDIAN {
            return math.Float64frombits(binary.BigEndian.Uint64(ram)), nil
        }
    }
    return "", errors.New(UnknownVarSize)
}

// Integer Conversion following base
// @value integer number value
// @base base number
// @zeroPadding whether padding using zero
// @return string of after converted
func baseConvert(v interface{}, base int) (string, error) {
    bits := binary.Size(v) * 8
    if bits < 0 {
        return "", errors.New(UseBaseType)
    }

    switch base {
    case BINARY:
        return fmt.Sprintf("%0*b", bits, v), nil
    case OCTAL:
        return fmt.Sprintf("%0*o", bits/3+1, v), nil
    case DECIMAL:
        return fmt.Sprintf("%d", v), nil
    case HEXADECIMAL:
        return fmt.Sprintf("%0*X", bits/4, v), nil
    }

    return "", errors.New("Unknown base")
}

// float convert to binary
func baseConvertFloat(v interface{}) (interface{}, error) {
    bits := binary.Size(v) * 8
    if bits < 0 {
        return nil, errors.New(UseBaseType)
    }
    if bits == 32 {
        return math.Float32bits(v.(float32)), nil
    } else if bits == 64 {
        return math.Float64bits(v.(float64)), nil
    }
    return nil, errors.New(UnknownVarSize)
}

// calculate complement code. invert and add one
// @v value
// @base base
// @return string of complement code
func neg(v interface{}) (interface{}, error) {
    bits := binary.Size(v) * 8
    if bits < 0 {
        return nil, errors.New(UseBaseType)
    }

    if bits == 32 {
        if v.(int32) >= 0 {
            // the complememt of a positive number is equal to the original code
            return v.(uint32), nil
        } else {
            // 1 << uint(32) set the highest to 1 and the rest to 0.
            // 100000000000000000000000000000000 （1 + 32 zeros） The maximum representation range.
            // The complement of negative numbers = maximum representation range - the positive value
            // The complement code = maximum representaion range + value.
            complementCode := uint32((1 << bits) + v.(int32))
            return complementCode, nil
        }
    } else if bits == 64 {
        if v.(int64) >= 0 {
            // the complememt of a positive number is equal to the original code
            return v.(uint64), nil
        } else {
            // 1 << uint(32) set the highest to 1 and the rest to 0.
            // 100000000000000000000000000000000 （1 + 32 zeros） The maximum representation range.
            // The complement of negative numbers = maximum representation range - the positive value
            // The complement code = maximum representaion range + value.
            complementCode := (uint64)((1 << bits) + v.(int64))
            return complementCode, nil
        }
    }
    return nil, errors.New(UnknownVarSize)
}
