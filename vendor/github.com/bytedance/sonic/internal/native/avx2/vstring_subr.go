// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package avx2

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__vstring = 96
)

const (
    _stack__vstring = 72
)

const (
    _size__vstring = 1808
)

var (
    _pcsp__vstring = [][2]uint32{
        {0x1, 0},
        {0x6, 8},
        {0x8, 16},
        {0xa, 24},
        {0xc, 32},
        {0xd, 40},
        {0x11, 48},
        {0x65f, 72},
        {0x660, 48},
        {0x662, 40},
        {0x664, 32},
        {0x666, 24},
        {0x668, 16},
        {0x669, 8},
        {0x66d, 0},
        {0x710, 72},
    }
)

var _cfunc_vstring = []loader.CFunc{
    {"_vstring_entry", 0,  _entry__vstring, 0, nil},
    {"_vstring", _entry__vstring, _size__vstring, _stack__vstring, _pcsp__vstring},
}
