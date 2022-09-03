package zip

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Zipper struct {
	//writer io.Writer
	archive *zip.Writer
}

func NewZipper(writer io.Writer) *Zipper {
	return &Zipper{
		archive: zip.NewWriter(writer),
	}
}

func (z *Zipper) Close() error {
	return z.archive.Close()
}

func (z *Zipper) CompressFile(filename string, base string) error {
	info, err := os.Stat(filename)
	if err != nil {
		return err
	}
	return z.CompressFileWithInfo(filename, base, info)
}

func (z *Zipper) CompressFileWithInfo(filename string, base string, info os.FileInfo) error {
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Name = strings.TrimPrefix(filename, base)
	header.Method = zip.Deflate

	writer, err := z.archive.CreateHeader(header)
	if err != nil {
		return err
	}

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(writer, file)
	if err != nil {
		return err
	}

	return z.archive.Flush()
}

func (z *Zipper) CompressFileWithInfoAndReader(name string, info os.FileInfo, reader io.Reader) error {
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Name = name
	header.Method = zip.Deflate

	writer, err := z.archive.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, reader)
	if err != nil {
		return err
	}

	return z.archive.Flush()
}

func (z *Zipper) CompressFileInfoAndContent(name string, info os.FileInfo, data []byte) error {
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Name = name
	header.Method = zip.Deflate

	writer, err := z.archive.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = writer.Write(data)
	if err != nil {
		return err
	}

	return z.archive.Flush()
}

func (z *Zipper) CompressDir(dir string) error {
	base := dir + string(os.PathSeparator)
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path == dir {
			return nil
		}

		if info.IsDir() {
			return nil
		}

		return z.CompressFileWithInfo(path, base, info)
	})
}
