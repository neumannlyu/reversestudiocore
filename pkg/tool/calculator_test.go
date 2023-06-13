package tool

import (
    "fmt"
    "testing"
)

func TestUint32ToRamLittleEndian(t *testing.T) {
    // fmt.Printf("uint32ToRamLittleEndian(999): %v\n", ramLayoutUint32(52, LITTLE_ENDIAN))

    // test read ram to float

    if v, err := ramData2Float([]byte{0x00, 0x68, 0x45, 0x44}, LITTLE_ENDIAN); err != nil {
        fmt.Printf("err.Error(): %v\n", err.Error())
    } else {
        fmt.Printf("v: %v\n", v)
    }

    // fmt.Printf("uint32Convert(999, DECIMAL): %v\n", baseConvertUint32(int32(100), BINARY))

    // test neg
    /*
       if v, err := neg(int32(-100)); err != nil {
           fmt.Printf("err.Error(): %v\n", err.Error())
       } else {
           if s, err := baseConvert(v, BINARY); err != nil {
               fmt.Printf("err.Error(): %v\n", err.Error())
           } else {
               fmt.Printf("s: %v\n", s)
           }
       }
    */

    // test convert float
    /*
       if v, err := baseConvertFloat(float32(895.75)); err != nil {
           fmt.Printf("err.Error(): %v\n", err.Error())
       } else {
           if s, err := ramLayout(v, LITTLE_ENDIAN); err != nil {
               fmt.Printf("err.Error(): %v\n", err.Error())
           } else {
               fmt.Printf("s: %v\n", s)
           }
           // fmt.Printf("v: %v\n", v)
       }
    */
}
