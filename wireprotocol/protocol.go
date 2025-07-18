// Package wireprotocol The Wire Protocol for Hobor
package wireprotocol

import (
	"bytes"
	"errors"
	"io"
	"strconv"
)

const (
	Header = "content-size"
	// Delimiter = `\r\n\r\n`
	hsize = len(Header) + 4
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

func (hb *HoborConn) ReadMessage() ([]byte, error) {
	header := make([]byte, hsize)
	// not tracking how many bytes read since it is fixed
	_, err := hb.conn.Read(header)
	if err != nil {
		return nil, err
	}
	hvals := bytes.Split(header, []byte(":"))
	if string(hvals[0]) != Header {
		return nil, errors.New("content size not sent in message")
	}
	size, err := strconv.Atoi(string(hvals[1]))
	if err != nil {
		return nil, err
	}
	data := make([]byte, size)
	_, err = hb.conn.Read(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (hb *HoborConn) WriteMessage(msg []byte) error {
	size := strconv.Itoa(len(msg))
	header := Header + ":" + size
	n, err := hb.conn.Write([]byte(header))
	if err != nil {
		return err
	}
	if n != len(header) {
		return errors.New("header write mismatch")
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
