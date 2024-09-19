package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func copyFile(src, dst string) error {
	fmt.Printf("Copying file %s to %s\n", src, dst)
	_ = os.MkdirAll(filepath.Dir(dst), 0750)
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = srcFile.WriteTo(dstFile)
	if err != nil {
		return err
	}

	return nil
}
