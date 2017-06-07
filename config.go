package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"unicode"
)

const (
	CONFIG_FILE string = ".harpoon"
	Tag_EOF     int32  = 0
	Tag_EQUAL   int32  = int32('=')
	Tag_WORD    int32  = 256 + iota
	Tag_REPO
	Tag_IMAGE
)

var (
	Token_EOF = token{Tag_EOF, "EOF"}
	words     = map[string]*token{
		"repo":  &token{Tag_REPO, "repo"},
		"image": &token{Tag_IMAGE, "image"},
	}
)

type config struct {
	Repo  string
	Image string
}

type token struct {
	Tag int32
	V   string
}

func loadConfig() *config {
	f, err := os.Open(CONFIG_FILE)
	if err != nil {
		if err == os.ErrNotExist {
			fmt.Fprintln(os.Stderr, ".harpoon file not exists, please init first")
		} else {
			panic(err)
		}
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	var c config
	err = parseConfig(&c, b)
	if err != nil {
		panic(err)
	}

	return &c
}

func scan(r io.ByteScanner) *token {
	var (
		b   byte
		err error
	)
	// Skip whitespaces
	for {
		b, err = r.ReadByte()

		if err != nil {
			if err == io.EOF {
				return &Token_EOF
			} else {
				panic(err)
			}
		}

		if !unicode.IsSpace(rune(b)) {
			break
		}
	}

	if unicode.IsLetter(rune(b)) {
		buf := new(bytes.Buffer)
		err = buf.WriteByte(b)
		if err != nil {
			panic(err)
		}

		for {
			b, err = r.ReadByte()
			if err != nil {
				panic(err)
			}
			if unicode.IsDigit(rune(b)) || unicode.IsLetter(rune(b)) || b == '.' {
				buf.WriteByte(b)
			} else {
				r.UnreadByte()
				break
			}
		}

		t, ok := words[buf.String()]
		if ok {
			return t
		}
		t = &token{Tag_WORD, buf.String()}
		words[t.V] = t
		return t
	}

	if b == '=' {
		return &token{Tag_EQUAL, "="}
	}

	return &token{int32(b), fmt.Sprintf("%c", b)}
}

func match(t *token, tag int32) {
	if t.Tag != tag {
		panic("Syntax error")
	}
}

func parseConfig(dst *config, bs []byte) error {
	r := bytes.NewReader(bs)
	for {
		t := scan(r)

		if t.Tag == Tag_REPO {
			t = scan(r)
			match(t, Tag_EQUAL)
			t = scan(r)
			match(t, Tag_WORD)
			dst.Repo = t.V
		}

		if t.Tag == Tag_IMAGE {
			t = scan(r)
			match(t, Tag_EQUAL)
			t = scan(r)
			match(t, Tag_WORD)
			dst.Image = t.V
		}

		if t.Tag == Tag_EOF {
			break
		}
	}

	return nil
}
