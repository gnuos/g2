// Copyright 2013 Xing Xing <mikespook@gmail.com>.
// All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package client

import (
	"encoding/binary"
)

// request
type request struct {
	DataType uint32
	Data     []byte
}

// Encode a Request to byte slice
func (req *request) Encode() (data []byte) {
	l := len(req.Data)       // length of data
	tl := l + MIN_PACKET_LEN // add 12 bytes head
	data = getBuffer(tl)
	copy(data[:4], REQ_STR)
	binary.BigEndian.PutUint32(data[4:8], req.DataType)
	binary.BigEndian.PutUint32(data[8:12], uint32(l))
	copy(data[MIN_PACKET_LEN:], req.Data)
	return
}

func getRequest() (req *request) {
	// TODO add a pool
	req = &request{}
	return
}

func getJob(funcname, id string, data []byte) (req *request) {
	req = getRequest()
	a := len(funcname)
	b := len(id)
	c := len(data)
	l := a + b + c + 2
	req.Data = getBuffer(l)
	copy(req.Data[0:a], funcname)
	copy(req.Data[a+1:a+b+1], []byte(id))
	copy(req.Data[a+b+1:a+b+c+1], data)
	return
}
