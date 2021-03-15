package encode

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"
)

// TextMsg - message struct
type TextMsg struct {
	Str       string
	Timestamp time.Time
}

// Encoder return encoded string as bytes buffer
func Encode(str string) bytes.Buffer {
	msg := TextMsg{
		Str:       str,
		Timestamp: time.Now(),
	}

	fmt.Printf("Encode msg: %v\n", msg)

	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(msg)
	fmt.Printf("Encoded msg: (err=%v) - %v\n", err, buff)

	return buff
}

// Decoder return decoded struct
func Decode(buff bytes.Buffer) TextMsg {
	msg := TextMsg{}

	fmt.Printf("Decode msg\n")
	dec := gob.NewDecoder(&buff)
	dec.Decode(&msg)
	fmt.Printf("Decoded msg: (err=%v) - %v\n", err, msg)
	return msg
}