package main

import (
	"bytes"
	"encoding/hex"
	"log"
	"strconv"
	"strings"

	"github.com/jacobsa/go-serial/serial"
)

const (
	cmd      = "7A 7A 73 6E 3B"
	comPort  = "COM3"
	baud     = "9600"
	dataBits = 8
	stopBits = 1
	readSize = 4
)

func main() {
	hexCmd := strings.Split(cmd, " ")
	byteCmd := strings.Join(hexCmd, "")

	baudInt, err := strconv.ParseInt(baud, 10, 32)
	if err != nil {
		panic(err)
	}

	data, err := hex.DecodeString(byteCmd)
	if err != nil {
		panic(err)
	}

	portOptions := serial.OpenOptions{
		PortName:        comPort,
		BaudRate:        uint(baudInt),
		DataBits:        dataBits,
		StopBits:        stopBits,
		MinimumReadSize: readSize,
	}

	port, err := serial.Open(portOptions)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	defer port.Close()

	_, err = port.Write(data)
	if err != nil {
		log.Fatalf("port.Write: %v", err)
	}

	readBuf := make([]byte, 100)

	if _, err := port.Read(readBuf); err != nil {
		log.Panic(err)
	} else {
		outputBuf, _ := FormatBuf(readBuf)

		log.Println(outputBuf)
	}
}

// FormatBuf returns a trimmed string of the byte array
// as well as a trimmed version of the byte array itself
func FormatBuf(readBuf []byte) (string, []byte) {
	trimmedBuf := bytes.Trim(readBuf, "\x00")
	outputBuf := string(trimmedBuf)

	return outputBuf, trimmedBuf
}
