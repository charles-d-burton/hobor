// Package wireprotocol The Wire Protocol for Hobor
package wireprotocol

import (
	"errors"
	"io"
	"math"
)

const (
	ContentSize    = 4 // The number of bytes in a uint32
	LenDelimiter   = 8
	MaxMessageSize = 65_536
	ReadBuffer     = 256
	ControlSignal  = math.MaxUint32
)

type ControlMessage byte

const (
	MessageOK ControlMessage = iota
	MessageFail
	MessageError
)

var srSn = []byte{0x5c, 0x72, 0x5c, 0x6e} // \r\n

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
	size := bytesToInt(contentSize)
	if size == 0 {
		return nil, errors.New("invalid message size")
	}
	if size > MaxMessageSize {
		return nil, errors.New("message size exceeds maximum of 64k")
	}
	if size == ControlSignal {
		cm, err := hb.readControlMessage()
		if err != nil {
			return nil, err
		}
		return []byte{byte(cm)}, nil

	}
	return hb.readMessage(size)
}

// WriteMessage writes a heeader with the calculated message size
// and then the corresponding data
func (hb *HoborConn) WriteMessage(msg []byte) error {
	if len(msg) > MaxMessageSize {
		return errors.New("message size exceeds maximum of 64k")
	}
	return hb.writeMessage(msg)
}

func (hb *HoborConn) WriteControlMessage(cm ControlMessage) error {
	payload := intToBytes(ControlSignal)
	payload = append(payload, byte(cm))
	return hb.writePayload(payload)
}

func (hb *HoborConn) Close() error {
	return hb.conn.Close()
}

func (hb *HoborConn) readControlMessage() (ControlMessage, error) {
	m := make([]byte, 1)
	_, err := hb.conn.Read(m)
	if err != nil {
		return MessageError, err
	}
	switch ControlMessage(m[0]) {
	case MessageOK:
		return MessageOK, nil
	case MessageFail:
		return MessageFail, nil
	case MessageError:
		return MessageError, nil
	default:
		return MessageError, nil
	}
}

func (hb *HoborConn) readMessage(size uint32) ([]byte, error) {
	payload := make([]byte, size)
	n, err := io.ReadFull(hb.conn, payload)
	if err != nil {
		return nil, err
	}
	if n != int(size) {
		return nil, errors.New("failed to read all data from socket")
	}
	delim := make([]byte, LenDelimiter)
	n, err = io.ReadFull(hb.conn, delim)
	if err != nil {
		return nil, err
	}
	if n != LenDelimiter {
		return nil, errors.New("invalid message break received")
	}
	if !compareBytes(delim[0:4], srSn) && !compareBytes(delim[4:LenDelimiter-1], srSn) {
		return nil, errors.New("invalid message break received")
	}
	return payload, nil
}

func (hb *HoborConn) writeMessage(msg []byte) error {
	msgSize := uint32(len(msg))
	payload := intToBytes(msgSize)
	payload = append(payload, msg...)
	return hb.writePayload(payload)
}

func (hb *HoborConn) writePayload(msg []byte) error {
	delim := append(srSn, srSn...)
	msg = append(msg, delim...)

	n, err := hb.conn.Write(msg)
	if err != nil {
		return err
	}
	if n != len(msg) {
		return errors.New("unable to write all data to connection")
	}
	return nil
}

// func (hb *HoborConn) reset() error {
// 	delim := append(srSn, srSn...)
// 	reader := make([]byte, LenDelimiter)
// 	n, err := hb.conn.Read(reader)
//
// 	return nil
// }

// Functions to reduce overall imports for tinygo.  Converts a [4]byte to an integer
func bytesToInt(b []byte) uint32 {
	if len(b) != 4 {
		return 0
	}
	result := int32(b[0])<<24 | int32(b[1])<<16 | int32(b[2])<<8 | int32(b[3])
	return uint32(result)
}

func intToBytes(num uint32) []byte {
	b := make([]byte, 4)
	b[0] = byte(num >> 24)
	b[1] = byte(num >> 16)
	b[2] = byte(num >> 8)
	b[3] = byte(num)
	return b
}

func compareBytes(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
