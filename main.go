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
	// Example of READ USB MODE
	cmd      = "7A 7A 6D 64 3B"
	comPort  = "COM4"
	baud     = "9600"
	dataBits = 8
	stopBits = 1
	readSize = 4
)

func main() {
	baudInt, data := CreateBaudRateAndCmdData()

	portOptions := serial.OpenOptions{
		PortName:        comPort,
		BaudRate:        baudInt,
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

	if count, err := port.Read(readBuf); err != nil {
		log.Panic(err)
	} else {
		outputBuf, _ := FormatBuf(readBuf)

		log.Println("-- bytes received: ", count)
		log.Println("-- msg received: ", outputBuf)
	}
}

// FormatBuf returns a trimmed string of the byte array
// as well as a trimmed version of the byte array itself
func FormatBuf(readBuf []byte) (string, []byte) {
	trimmedBuf := bytes.Trim(readBuf, "\x00")
	outputBuf := string(trimmedBuf)

	return outputBuf, trimmedBuf
}

// CreateBaudRateAndCmdData grabs user defined cmd and baud
// to convert for serial port library
func CreateBaudRateAndCmdData() (uint, []byte) {
	hexCmd := strings.Split(cmd, " ")

	log.Println("-- pre parsed cmd: ", hexCmd)

	byteCmd := strings.Join(hexCmd, "")

	baudInt, err := strconv.ParseInt(baud, 10, 32)
	if err != nil {
		panic(err)
	}

	data, err := hex.DecodeString(byteCmd)
	if err != nil {
		panic(err)
	}

	log.Println("-- sending decimal: ", data)

	return uint(baudInt), data
}
