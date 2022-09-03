package zip

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func ZipIntoWriter(dir string, writer io.Writer) error {
	archive := zip.NewWriter(writer)
	defer archive.Close()

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

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = strings.TrimPrefix(path, dir+string(os.PathSeparator))
		header.Method = zip.Deflate

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(writer, file)
		if err != nil {
			return err
		}
		return nil
	})
}

func ZipDir(dir string, file string) error {
	zf, err := os.Create(file)
	if err != nil {
		return err
	}
	defer zf.Close()

	return ZipIntoWriter(dir, zf)
}

func UnzipFile(zipFile string, dir string) error {
	zf, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer zf.Close()

	for _, f := range zf.File {
		filename := filepath.Join(dir, f.Name)
		if f.FileInfo().IsDir() {
			_ = os.MkdirAll(filename, os.ModePerm)
		} else {
			_ = os.MkdirAll(filepath.Dir(filename), os.ModePerm)

			reader, err := f.Open()
			if err != nil {
				return err
			}
			defer reader.Close()

			file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(file, reader)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func Unzip(reader io.ReaderAt, size int64, dir string) error {
	zf, err := zip.NewReader(reader, size)
	if err != nil {
		return err
	}
	for _, f := range zf.File {
		filename := filepath.Join(dir, f.Name)
		if f.FileInfo().IsDir() {
			_ = os.MkdirAll(filename, os.ModePerm)
		} else {
			_ = os.MkdirAll(filepath.Dir(filename), os.ModePerm)

			reader, err := f.Open()
			if err != nil {
				return err
			}
			defer reader.Close()

			file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(file, reader)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
