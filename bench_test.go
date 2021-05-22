package main

import (
	"io"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/klauspost/pgzip"
	"github.com/planetscale/pargzip"
)

type writer struct{}

func (w *writer) Write(b []byte) (int, error) { return len(b), nil }

var w = &writer{}

func listFiles(b *testing.B, dir string) []string {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		b.Fatal(err)
	}

	filenames := make([]string, 0, len(entries))
	for _, f := range entries {
		if f.IsDir() {
			continue
		}

		filenames = append(filenames, f.Name())
	}

	return filenames
}

const canterburyCorpusDir = "testdata/canterbury"

var implementations = []struct {
	name      string
	NewWriter func(w io.Writer) io.WriteCloser
}{
	{
		name:      "pargzip",
		NewWriter: func(w io.Writer) io.WriteCloser { return pargzip.NewWriter(w) },
	},
	{
		name:      "pgzip",
		NewWriter: func(w io.Writer) io.WriteCloser { return pgzip.NewWriter(w) },
	},
}

func BenchmarkCompression(b *testing.B) {
	filenames := listFiles(b, canterburyCorpusDir)

	for _, filename := range filenames {
		b.Run(filename, func(b *testing.B) {
			data, err := ioutil.ReadFile(filepath.Join(canterburyCorpusDir, filename))
			if err != nil {
				b.Fatal(err)
			}

			b.ResetTimer()
			for _, impl := range implementations {
				b.Run(impl.name, func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						compressor := impl.NewWriter(w)
						_, err := compressor.Write(data)
						compressor.Close()

						if err != nil {
							b.Error(err)
							continue
						}
					}
				})
			}
		})
	}
}
