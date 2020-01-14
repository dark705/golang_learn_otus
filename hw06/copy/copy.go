package copy

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

func Copy(from string, to string, limit int, offset int) error {
	var iFile *os.File
	var oFile *os.File

	iFile, err := os.OpenFile(from, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer func() { _ = iFile.Close() }()

	iFileStat, err := iFile.Stat()
	if err != nil {
		return err
	}
	iFileLen := int(iFileStat.Size())

	if offset >= iFileLen {
		return errors.New("offset is bigger or equal len of input file")
	}

	if limit == 0 || limit > iFileLen {
		limit = iFileLen
	}

	if offset+limit > iFileLen {
		limit = iFileLen - offset
	}

	_, err = iFile.Seek(int64(offset), 0)
	if err != nil {
		return err
	}

	oFile, err = os.OpenFile(to, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer func() { _ = iFile.Close() }()

	bar := pb.StartNew(limit)
	bar.Set(pb.Bytes, true)
	bar.Set(pb.SIBytesPrefix, true)

	bSize := 2
	buff := make([]byte, bSize)
	var totalReadBytes int

	for totalReadBytes < limit {
		var r int

		if totalReadBytes+bSize > limit { //last read, for not get unexpected eol
			_, err = io.ReadFull(iFile, buff[:limit-totalReadBytes])
			r = limit - totalReadBytes
		} else {
			r, err = io.ReadFull(iFile, buff)
		}
		if err != nil {
			return err
		}

		w, err := oFile.Write(buff[:r])
		if r != w || err != nil {
			return err
		}
		totalReadBytes += r
		bar.Add(r)
	}
	bar.Finish()

	return nil
}
