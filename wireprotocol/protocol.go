// Package wireprotocol The Wire Protocol for Hobor
package wireprotocol

import (
	"bytes"
	"errors"
	"io"
	"strconv"
)

const (
	Header      = "content-size"
	Delimiter   = `\r\n\r\n`
	bufferSize  = 4096
	MessageSize = 65_536
)

// type message struct{}

type HoborConn struct {
	conn   io.ReadWriteCloser
	closed bool
}

func NewHoborConn(conn io.ReadWriteCloser) (*HoborConn, error) {
	if conn != nil {
		return &HoborConn{conn: conn, closed: false}, nil
	}
	return nil, errors.New("connection not set")
}

// ReadMessage reads the data from the connection in.
// Probably needs some thought around message size limitations
func (hb *HoborConn) ReadMessage() ([]byte, error) {
	buffer := make([]byte, bufferSize)
	// not tracking how many bytes read since it is fixed
	_, err := hb.conn.Read(buffer)
	if err != nil {
		return nil, err
	}
	headerAndData := bytes.Split(buffer, []byte(Delimiter))
	headerAndValue := bytes.Split(headerAndData[0], []byte(":"))
	if string(headerAndValue[0]) != Header && len(headerAndValue[1]) <= 0 {
		return nil, errors.New("content size not sent in message")
	}
	size, err := strconv.Atoi(string(headerAndValue[1]))
	if err != nil {
		return nil, err
	}
	if size > MessageSize {
		return nil, errors.New("message size exceeds maximum of 64k")
	}
	size = size - len(headerAndData[1])
	data := make([]byte, size)
	_, err = hb.conn.Read(data)
	if err != nil {
		return nil, err
	}
	return append(headerAndData[1], data...), nil
}

// WriteMessage writes a heeader with the calculated message size
// and then the corresponding data
func (hb *HoborConn) WriteMessage(msg []byte) error {
	size := strconv.Itoa(len(msg))
	if len(msg) > MessageSize {
		return errors.New("message size exceeds maximum of 64k")
	}
	header := Header + ":" + size
	n, err := hb.conn.Write([]byte(header))
	if err != nil {
		return err
	}
	if n != len(header) {
		return errors.New("header write mismatch")
	}
	_, err = hb.conn.Write([]byte(Delimiter))
	if err != nil {
		return err
	}
	n, err = hb.conn.Write(msg)
	if err != nil {
		return err
	}
	if n != len(msg) {
		return errors.New("body write mismatch")
	}
	return nil
}

func (hb *HoborConn) Close() error {
	return hb.conn.Close()
}
