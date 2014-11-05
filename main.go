package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

const PAYLOAD_SIZE = 1000
const PACKET_COUNT = 1000

var t = flag.String("t", "ip, tcp or udp", "Protocol")
var d = flag.String("d", "127.0.0.1:8080", "Destination IP Address and Port Number")
var p = flag.String("p", "test.txt", "Packet Payload(HEX) file size < "+fmt.Sprintf("%d", PAYLOAD_SIZE)+"byte")
var c = flag.Int("c", 10, "Packet Count < "+fmt.Sprintf("%d", PACKET_COUNT))

func main() {
	flag.Parse()
	protocol, destination, payload, count := "", "", "", 0
	var payloadSize int64
	checkCount := 0

	flag.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "t":
			protocol = fmt.Sprintf("%v", f.Value)
			switch protocol {
			case "ip", "tcp", "udp":
				checkCount++
			default:
				fmt.Println("Error")
			}
		case "d":
			destination = fmt.Sprintf("%v", f.Value)
			checkCount++
		case "p":
			payload = fmt.Sprintf("%v", f.Value)
			stat, err := os.Stat(payload)
			if err != nil {
				log.Fatal(err)
			}
			if stat.Size() <= PAYLOAD_SIZE {
				checkCount++
			}
		case "c":
			count = *c
			if count <= PACKET_COUNT {
				checkCount++
			}
		}
	})
	if checkCount == 4 {
		hexdataB, hexdataA := "", ""
		hexdataBuf := make([]byte, PAYLOAD_SIZE)
		file, err := os.Open(payload)
		if err != nil {
			log.Fatal(err)
		}
		for {
			n, _ := file.Read(hexdataBuf)
			if 0 == n {
				break
			}
			hexdataB += string(hexdataBuf)
		}
		for i := 0; i < len(hexdataB); i++ {
			switch string(hexdataB[i]) {
			case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F":
				hexdataA += string(hexdataB[i])
				payloadSize++
			default:
			}
		}
		bufTmp := make([]byte, 2)
		hexBuf := make([]byte, PAYLOAD_SIZE)
		y := 0
		for z := 0; z < len(hexdataA); z += 2 {
			if z+1 < len(hexdataA) {
				bufTmp[0] = hexdataA[z]
				bufTmp[1] = hexdataA[z+1]
				hexBuf[y] = hexByte(bufTmp)
			} else {
				bufTmp[0] = hexdataA[z]
				bufTmp[1] = 0
				hexBuf[y] = hexByte(bufTmp)
				payloadSize++
			}
			y++
		}
		defer file.Close()

		for i := 0; i < count; i++ {
			conn, err := net.Dial(protocol, destination)
			if err != nil {
				log.Fatal(err)
			}
			conn.Write(hexBuf[0 : payloadSize/2])
			defer conn.Close()
		}
		fmt.Println("Success!!")
	} else {
		flag.PrintDefaults()
	}
}
func hexByte(htob []byte) uint8 {
	byteUint := map[string]uint8{
		"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8,
		"9": 9, "A": 10, "B": 11, "C": 12, "D": 13, "E": 14, "F": 15,
	}
	var tmpu, tmpuA, tmpuB uint8
	tmpuA = byteUint[string(htob[0])]
	tmpuB = byteUint[string(htob[1])]
	tmpu = tmpuA << 4
	tmpu = tmpu | tmpuB
	return tmpu
}
