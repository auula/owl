// MIT License

// Copyright (c) 2022 Leon Ding <ding@ibyte.me>

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package scan

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"hash"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	ErrBufferIsEmpty = errors.New("buffer is empty")
	ErrFileIsEmpty   = errors.New("file path is empty")
)

const (
	// When a single file exceeds this size limit
	// the fragmented md5 algorithm will be used
	fileMaxSize = 10 << 10 << 10 // 10MB
)

// IsDir check if the path is a directory
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// Files returns the collection of all file paths under the specified path
func Files(folder string) ([]string, error) {
	var result []string

	filepath.Walk(folder, func(filePath string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !fi.IsDir() {
			// If you want to ignore this directory, return filepath.SkipDir, ie:
			// return filepath.SkipDir
			result = append(result, filePath)
		}

		return nil
	})

	return result, nil
}

// Md5 get the md5 value of the file based on the path
func Md5(filepath string) (string, error) {

	file, err := os.Open(filepath)

	defer file.Close()
	if err != nil {
		return "", err
	}

	fileInfo, err := os.Stat(filepath)

	if err != nil {
		return "", err
	}

	hashed := md5.New()

	if fileInfo.Size() >= fileMaxSize {
		return blockMd5(file, fileInfo.Size(), hashed)
	} else {
		io.Copy(hashed, file)
	}

	// io.Copy(hashed, file)

	return hex.EncodeToString(hashed.Sum(nil)), nil
}

// blockMd5 Built-in sharding md5 algorithm used
func blockMd5(file *os.File, size int64, hashed hash.Hash) (string, error) {
	// Intercept data segment size window
	var (
		head = make([]byte, 500)
		body = make([]byte, 24)
		tail = make([]byte, 500)
	)

	if _, err := file.ReadAt(head, 0); err != nil {
		return "", err
	}

	if _, err := file.ReadAt(body, size/4); err != nil {
		return "", err
	}

	if _, err := file.ReadAt(tail, size-512); err != nil {
		return "", err
	}
	hashed.Write(head)
	hashed.Write(body)
	hashed.Write(tail)
	return hex.EncodeToString(hashed.Sum(nil)), nil
}

// HexDecode hex decryption
func HexDecode(s string) []byte {
	dst := make([]byte, hex.DecodedLen(len(s)))
	n, err := hex.Decode(dst, []byte(s))
	if err != nil {
		return nil
	}
	return dst[:n]
}

// HexEncode hex encoding
func HexEncode(s string) []byte {
	dst := make([]byte, hex.EncodedLen(len(s)))
	n := hex.Encode(dst, []byte(s))
	return dst[:n]
}

// HexDump convert the specified file to hexadecimal
func HexDump(path string) (string, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return NilString, err
	}
	return hex.Dump(bytes), nil
}
