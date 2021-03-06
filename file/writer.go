package file

import (
	"bufio"
	"log"
	"os"

	"github.com/gravetii/diztl/diztl"
	"github.com/gravetii/diztl/util"
)

// Writer : The file writer.
type Writer struct {
	metadata *diztl.FileMetadata
	buf      *bufio.Writer
	f        *os.File
}

// Close : Closes the resources held by this writer and returns the created file.
func (obj *Writer) Close() *os.File {
	obj.buf.Flush()
	obj.f.Close()
	return obj.f
}

// CreateWriter : Returns an instance of the Writer for the given file metadata.
func CreateWriter(m *diztl.FileMetadata) (*Writer, error) {
	f, err := createFile(m)
	if err != nil {
		return nil, err
	}

	buf := bufio.NewWriter(f)
	return &Writer{metadata: m, buf: buf, f: f}, nil
}

func createFile(metadata *diztl.FileMetadata) (*os.File, error) {
	fname := util.GetFilename(metadata)
	fpath := util.GetOutputPath(fname)
	f, err := os.Create(fpath)
	if err != nil {
		log.Printf("Unable to create file %s: %v\n", fpath, err)
		return nil, err
	}

	return f, nil
}

// Write : Writes the given set of bytes to the underlying buffer.
func (obj *Writer) Write(data []byte) error {
	_, err := obj.buf.Write(data)
	if err != nil {
		return err
	}

	return nil
}

// Name : Returns the name of the file that will be created by this file writer.
func (obj *Writer) Name() string {
	return obj.f.Name()
}

// Chunks : Returns the total number of chunks that need to be written to create the file.
func (obj *Writer) Chunks() int32 {
	return obj.metadata.GetChunks()
}
