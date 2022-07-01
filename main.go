package main

import (
	"fmt"
	"github.com/sigurn/crc8"
  "encoding/hex"
  "log"
)

// settings extracted from documentation
var CRC8_SPS30 = crc8.Params{0x31, 0xFF, false, false, 0x00, 0x00,"CRC-8/SPS30"}

func main() {	
  var (
	  table = crc8.MakeTable(CRC8_SPS30)
	  inStr = "BEEF"
    in  []byte
    err error
  )
  
  if in, err = hex.DecodeString(inStr); err != nil {
    log.Fatalf("Error decoding: %v", inStr)
  }

	crc := crc8.Checksum(in, table)
	fmt.Printf("source  : %04X \nchecksum: %04X\n", in, crc)
}