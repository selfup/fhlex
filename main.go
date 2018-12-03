package main

import (
	"encoding/hex"
	"log"
	"strconv"
	"strings"

	"github.com/jacobsa/go-serial/serial"
)

func main() {
	icomCmds := strings.Split("7A7A6D643B", " ")
	byteCmd := strings.Join(icomCmds, "")

	baudInt, err := strconv.ParseInt("9600", 10, 32)

	if err != nil {
		panic(err)
	}

	data, err := hex.DecodeString(byteCmd)
	if err != nil {
		panic(err)
	}

	options := serial.OpenOptions{
		PortName:        "COM4",
		BaudRate:        uint(baudInt),
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	defer port.Close()

	_, err = port.Write(data)
	if err != nil {
		log.Fatalf("port.Write: %v", err)
	}

	readBuf := make([]byte, 1000)

	if c, err := port.Read(readBuf); err != nil {
		log.Panic(err)
	} else {
		log.Println(string(readBuf))
		log.Print(c)
	}
}
