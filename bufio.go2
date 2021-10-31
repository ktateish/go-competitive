package main

// vim:set ft=go:

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
)

// snip ------------------------------------------------------------------------

// IO helper

type FullBufferedReader struct {
	b   []byte // holds input bytes
	cur int    // holds read position
	r   io.Reader
}

func NewFullBufferedReader(r io.Reader) (*FullBufferedReader, error) {
	res := &FullBufferedReader{
		r: r,
	}
	return res, nil
}

func (r *FullBufferedReader) ensureBuffered() error {
	if r.b == nil {
		b, err := ioutil.ReadAll(r.r)
		if err != nil {
			return err
		}
		r.b = b
		r.cur = 0
	}
	return nil
}

func (r *FullBufferedReader) NextToken() ([]byte, error) {
	if err := r.ensureBuffered(); err != nil {
		return nil, err
	}
	for r.cur < len(r.b) && bio_is_space(r.b[r.cur]) {
		r.cur++
	}
	if r.cur == len(r.b) {
		return nil, io.ErrUnexpectedEOF
	}
	j := r.cur
	for r.cur < len(r.b) && !bio_is_space(r.b[r.cur]) {
		r.cur++
	}
	return r.b[j:r.cur], nil
}

func (r *FullBufferedReader) NextLine() ([]byte, error) {
	if err := r.ensureBuffered(); err != nil {
		return nil, err
	}
	for r.cur < len(r.b) && bio_is_new_line(r.b[r.cur]) {
		r.cur++
	}
	if r.cur == len(r.b) {
		return nil, io.ErrUnexpectedEOF
	}
	j := r.cur
	for r.cur < len(r.b) && !bio_is_new_line(r.b[r.cur]) {
		r.cur++
	}
	return r.b[j:r.cur], nil
}

// -----------------------------------------------------------------------------

type LineBufferedReader struct {
	b   []byte        // holds input line
	r   *bufio.Reader // read buffer
	cur int           // holds read position
}

func NewLineBufferedReader(r io.Reader) (*LineBufferedReader, error) {
	return &LineBufferedReader{
		r: bufio.NewReader(r),
	}, nil
}

func (r *LineBufferedReader) ensureBuffered() error {
	if r.cur == len(r.b) {
		b, err := r.r.ReadBytes('\n')
		if err != nil {
			return err
		}
		r.cur = 0
		r.b = b
	}
	return nil
}

func (r *LineBufferedReader) NextToken() ([]byte, error) {
	for ok := true; ok; ok = r.cur == len(r.b) || bio_is_space(r.b[r.cur]) {
		if err := r.ensureBuffered(); err != nil {
			return nil, err
		}
		for r.cur < len(r.b) && bio_is_space(r.b[r.cur]) {
			r.cur++
		}
	}
	j := r.cur
	for r.cur < len(r.b) && !bio_is_space(r.b[r.cur]) {
		r.cur++
	}
	if r.cur == j {
		return nil, io.ErrUnexpectedEOF
	}
	return r.b[j:r.cur], nil
}

func (r *LineBufferedReader) NextLine() ([]byte, error) {
	if err := r.ensureBuffered(); err != nil {
		return nil, err
	}
	b := r.b[r.cur:]
	r.cur = len(r.b)
	return b, nil
}

// -----------------------------------------------------------------------------

type LineBufferedWriter struct {
	w io.Writer
	b []byte
}

func NewLineBufferedWriter(w io.Writer) (*LineBufferedWriter, error) {
	return &LineBufferedWriter{
		w: w,
		b: make([]byte, 0, 128),
	}, nil
}

func (w *LineBufferedWriter) Write(data []byte) (int, error) {
	for i, c := range data {
		w.b = append(w.b, c)
		if bio_is_new_line(c) && (i+1 == len(data) || !bio_is_space(data[i+1])) {
			err := w.Flush()
			if err != nil {
				return 0, err
			}
		}
	}
	return len(data), nil
}

func (w *LineBufferedWriter) Flush() error {
	_, err := w.w.Write(w.b)
	if err != nil {
		return err
	}
	w.b = w.b[:0]
	return nil
}

// -----------------------------------------------------------------------------

// Returns true if c is a white space
func bio_is_space(c byte) bool {
	switch c {
	case '\t', '\n', '\v', '\f', '\r', ' ':
		return true
	}
	return false
}

func bio_is_new_line(c byte) bool {
	switch c {
	case '\n', '\r':
		return true
	}
	return false
}

// -----------------------------------------------------------------------------

type BufferedReader interface {
	NextToken() ([]byte, error)
	NextLine() ([]byte, error)
}

type BufferedWriter interface {
	Write([]byte) (int, error)
	Flush() error
}

type BufferedIO struct {
	BufferedReader
	BufferedWriter
}

func NewFullBufferedIO(r io.Reader, w io.Writer) (*BufferedIO, error) {
	res := &BufferedIO{}
	br, err := NewFullBufferedReader(r)
	if err != nil {
		return nil, err
	}
	res.BufferedReader = br
	res.BufferedWriter = bufio.NewWriter(w)
	return res, nil
}

func NewLineBufferedIO(r io.Reader, w io.Writer) (*BufferedIO, error) {
	res := &BufferedIO{}
	br, err := NewLineBufferedReader(r)
	if err != nil {
		return nil, err
	}
	res.BufferedReader = br

	bw, err := NewLineBufferedWriter(w)
	if err != nil {
		return nil, err
	}
	res.BufferedWriter = bw
	return res, nil
}

// ReadBytes reads next token and return it as []byte
func (bio *BufferedIO) ReadBytes() ([]byte, error) {
	b, err := bio.NextToken()
	// Use b[::] for avoiding overridden by appending to the returning []byte
	return b[:len(b):len(b)], err
}

// ReadString reads next token and return it as string
func (bio *BufferedIO) ReadString() (string, error) {
	b, err := bio.ReadBytes()
	return string(b), err
}

// ReadLine reads next line as []byte.  Trailing '\n' will not be included.
// See also comments on ReadBytes()
func (bio *BufferedIO) ReadLine() ([]byte, error) {
	b, err := bio.NextLine()
	if err != nil {
		return nil, err
	}
	return b[:len(b):len(b)], nil
}

// ReadStringLine reads next line as string
func (bio *BufferedIO) ReadStringLine() (string, error) {
	b, err := bio.ReadLine()
	return string(b), err
}

// LeadInt64 reads next token and return it as int64
func (bio *BufferedIO) ReadInt64() (int64, error) {
	s, err := bio.ReadString()
	if err != nil {
		return 0, fmt.Errorf("reading string: %w", err)
	}

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("parsing int: %w", err)
	}

	return i, nil
}

// ReadInt reads next token and return it as int
func (bio *BufferedIO) ReadInt() (int, error) {
	i, err := bio.ReadInt64()
	return int(i), err
}

// ReadFloat64 reads next token and return it as float64
func (bio *BufferedIO) ReadFloat64() (float64, error) {
	s, err := bio.ReadString()
	if err != nil {
		return 0, fmt.Errorf("reading string: %w", err)
	}

	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("parsing float: %w", err)
	}

	return f, nil
}
