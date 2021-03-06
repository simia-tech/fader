// Copyright 2014 The fader authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"errors"
	"io"
	"math/big"

	"github.com/juju/errgo"
)

type decrypter struct {
	parent io.Reader
	aesGCM cipher.AEAD
}

var (
	ErrInvalidNonce = errors.New("tried to decrypt with a previouly used nonce")
)

func NewDecrypter(parent io.Reader, key []byte) (Reader, error) {
	aes, err := aes.NewCipher(key)
	if err != nil {
		return nil, errgo.Mask(err)
	}

	aesGCM, err := cipher.NewGCM(aes)
	if err != nil {
		return nil, errgo.Mask(err)
	}

	return &decrypter{
		parent: parent,
		aesGCM: aesGCM,
	}, nil
}

func (d *decrypter) Read(nonce *big.Int, data []byte) (int, error) {
	length := uint16(0)
	if err := binary.Read(d.parent, binary.BigEndian, &length); err != nil {
		return 0, errgo.Mask(err)
	}

	nonceBytes := make([]byte, d.aesGCM.NonceSize())
	if _, err := d.parent.Read(nonceBytes); err != nil {
		return 0, errgo.Mask(err)
	}

	cipherText := make([]byte, length)
	if _, err := d.parent.Read(cipherText); err != nil {
		return 0, errgo.Mask(err)
	}

	nonce.SetBytes(nonceBytes)

	plainText, err := d.aesGCM.Open(nil, nonceBytes, cipherText, []byte{})
	if err != nil {
		return 0, errgo.Mask(err)
	}
	copy(data, plainText)

	return len(plainText), nil
}
