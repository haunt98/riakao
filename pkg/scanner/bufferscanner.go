package scanner

import "riakao/pkg/token"

type BufferScanner struct {
	s   *Scanner
	buf buffer
}

type buffer struct {
	tok      token.Token
	text     string
	isCached bool
}

func NewBufferScanner(s *Scanner) *BufferScanner {
	return &BufferScanner{
		s: s,
	}
}

func (bs *BufferScanner) Scan() (token.Token, string) {
	if bs.buf.isCached {
		bs.buf.isCached = false
		return bs.buf.tok, bs.buf.text
	}

	bs.buf.tok, bs.buf.text = bs.s.Scan()
	return bs.buf.tok, bs.buf.text
}

func (bs *BufferScanner) Peek() (token.Token, string) {
	tok, text := bs.Scan()
	bs.unscan()
	return tok, text
}

func (bs *BufferScanner) unscan() {
	bs.buf.isCached = true
}
