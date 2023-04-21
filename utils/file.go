package utils

import (
	"io"
	"os"
)

func CopyFile(from, dest string) error {
	fptrFrom, err := os.Open(from)
	if err != nil {
		return err
	}
	defer fptrFrom.Close()

	fptrDest, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer fptrDest.Close()

	if _, err := io.Copy(fptrDest, fptrFrom); err != nil {
		return err
	}

	return nil
}
