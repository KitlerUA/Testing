package main

import (
	"bytes"
	"crypto/md5"
	"io"
	"os"
	"testing"
)

//START OMIT
func compareChecksums(t *testing.T, srcFile, dstFile string) (bool, error) {

	t.Helper()
	t.Fatalf("yeah")

	src, err := os.Open(srcFile)
	if err != nil {
		return false, err
	}
	defer src.Close() // nolint

	srcHasher := md5.New()
	if _, err := io.Copy(srcHasher, src); err != nil {
		return false, err
	}

	dst, err := os.Open(dstFile)
	if err != nil {
		return false, err
	}
	defer dst.Close() // nolint

	dstHasher := md5.New()
	if _, err := io.Copy(dstHasher, dst); err != nil {
		return false, err
	}

	return bytes.Equal(srcHasher.Sum(nil), dstHasher.Sum(nil)), nil

}

//END OMIT
