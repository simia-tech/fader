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

package crypt_test

import (
	"bytes"
	"encoding/hex"
	"math/big"
	"testing"

	. "github.com/posteo/fader/crypt"
)

func TestEncryption(t *testing.T) {
	e := setUp(t)

	buffer := &bytes.Buffer{}
	encrypter, err := NewEncrypter(buffer, e.key)
	e.assertNoError(err)

	nonce := big.NewInt(0)
	plainText := []byte{1, 2, 3, 4, 5, 6, 7, 8}

	n, err := encrypter.Write(nonce, plainText)
	e.assertNoError(err)
	e.assertEquals(8, n)
	e.assertEquals(
		"00180000000000000000000000002e3b1966d4bb71503ec7942a5f4e352735219d268cbdcda0",
		hex.EncodeToString(buffer.Bytes()))
}

func TestNonceAlternation(t *testing.T) {
	e := setUp(t)

	buffer := &bytes.Buffer{}
	encrypter, err := NewEncrypter(buffer, e.key)
	e.assertNoError(err)

	nonce := big.NewInt(2222222)
	plainText := []byte{1, 2, 3, 4, 5, 6, 7, 8}

	n, err := encrypter.Write(nonce, plainText)
	e.assertNoError(err)
	e.assertEquals(8, n)
	e.assertEquals(
		"001800000000000000000021e88e84d211ce6c805f66aa2924c8a4886e81e0d3a287f2dab83a",
		hex.EncodeToString(buffer.Bytes()))
}
