// SEE: https://stackoverflow.com/a/13657862 for "Convert [8]byte to a uint64"
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	array := []byte{0x00, 0x01, 0x08, 0x00, 0x08, 0x01, 0xab, 0x01}
	var num uint64
	_ = binary.Read(bytes.NewBuffer(array[:]), binary.LittleEndian, &num)
	fmt.Printf("%v, %x", array, num)
}
