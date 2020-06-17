package splitio

import (
	"bytes"
	"io"
)

const chunkSize = 128

type each struct {
	data []byte
	err  error
}

type gate chan each

func (g gate) Next() ([]byte, error) {
	e := <-g
	return e.data, e.err
}

//Read interface for splitio
type Read interface {
	//Next is retrun value before next seperate
	Next() ([]byte, error)
}

//New func is start read from reader
func New(r io.Reader, sep []byte) Read {
	g := make(gate)

	go func(r io.Reader, sep []byte, g gate) {
		defer close(g)

		chunk, cache := make([]byte, chunkSize), []byte{}
		for {
			n, err := r.Read(chunk)
			if err != nil && err != io.EOF {
				g <- each{data: nil, err: err}
				return
			}

			subs := bytes.Split(append(cache, chunk[:n]...), sep)
			ln := len(subs) - 1

			for i := 0; i < ln; i++ {
				g <- each{data: subs[i], err: nil}
			}
			cache = subs[ln]

			if err == io.EOF {
				g <- each{data: cache, err: err}
				break
			}
		}
	}(r, sep, g)

	return g
}
