package main

import (
	"io"
	"os"
)

// writeFile writes an interface state cache to file
func writeFile(f string, r io.Reader) error {
	w, err := os.Create(f)
	if err != nil {
		return err
	}

	_, err = io.Copy(w, r)
	w.Close()
	return err
}
