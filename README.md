# GO Packet Send

## About

This project is unique packet send for Go language.

## How To Use
### Make Packet Payload File
`
$ cat test.txt 
EF EF EF EF EF EF EF EF
00 11 22 33 44 55 66 77
88 99 AA BB CC DD EE FF 
00 11 22 33 44 55 66 77
88 99 AA BB CC DD EE FF 
00 11 22 33 44 55 66 77
88 99 AA BB CC DD EE FF 
00 11 22 33 44 55 66 77
88 99 AA BB CC DD EE FF 
11 22 33 4
`
### IP Packet Send
`go run main.go -t ip -c 1 -d 192.168.11.17 -p test.txt`

### TCP Packet Send
`go run main.go -t tcp -c 1 -d 192.168.11.17:60000 -p test.txt`

### UDP Packet Send
`go run main.go -t udp -c 1 -d 192.168.11.17:60000 -p test.txt`

## Parameter
$ go run main.go
  -c=10: Packet Count < 1000
  -d="127.0.0.1:8080": Destination IP Address and Port Number
  -p="test.txt": Packet Payload(HEX) file size < 1000byte
  -t="ip, tcp or udp": Protocol
