// Package wireprotocol The Wire Protocol for Hobor
package wireprotocol

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

const (
	// Header     = "content-size:"
	ContentSize = 4 //The number of bytes in a uin32
	// Delimiter   = `\r\n\r\n`
	MessageSize = 65_536
)

// type message struct{}

type HoborConn struct {
	conn io.ReadWriteCloser
}

func NewHoborConn(conn io.ReadWriteCloser) (*HoborConn, error) {
	if conn != nil {
		return &HoborConn{conn: conn}, nil
	}
	return nil, errors.New("connection not set")
}

// ReadMessage reads the data from the connection in.
// Probably needs some thought around message size limitations
func (hb *HoborConn) ReadMessage() ([]byte, error) {
	contentSize := make([]byte, ContentSize)
	// not tracking how many bytes read since it is fixed
	n, err := hb.conn.Read(contentSize)
	if err != nil {
		return nil, err
	}
	if n != ContentSize {
		return nil, errors.New("failed to read first 4 bytes of data stream for size computation")
	}
	var size int32
	err = binary.Read(bytes.NewReader(contentSize), binary.BigEndian, &size)
	if err != nil {
		return nil, errors.New("unable to convert first 4 bytes into size integer")
	}
	if size > MessageSize {
		return nil, errors.New("message size exceeds maximum of 64k")
	}
	payload := make([]byte, size)
	n, err = hb.conn.Read(payload)
	if err != nil {
		return nil, err
	}
	if n != int(size) {
		return nil, errors.New("failed to read all data from socket")
	}
	return payload, nil
}

// WriteMessage writes a heeader with the calculated message size
// and then the corresponding data
func (hb *HoborConn) WriteMessage(msg []byte) error {
	if len(msg) > MessageSize {
		return errors.New("message size exceeds maximum of 64k")
	}
	msgSizeArr := make([]byte, ContentSize)
	msgSize := uint32(len(msg))
	binary.BigEndian.PutUint32(msgSizeArr, msgSize)

	payload := make([]byte, ContentSize+len(msg))
	copy(payload, msgSizeArr)
	copy(payload[ContentSize-1:], msg)
	n, err := hb.conn.Write(payload)
	if err != nil {
		return err
	}
	if n != len(payload) {
		return errors.New("unable to write all data to connection")
	}
	return nil
}

func (hb *HoborConn) Close() error {
	return hb.conn.Close()
}
