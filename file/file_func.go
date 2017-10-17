package file

import (
	"crypto/md5"
	"encoding/hex"
	"math"
	"os"
)

const fileChunk = 8192

//FileHash struct for MD5's
type FileHash struct {
	Filename string
	Size     int64
	MD5      string
}

//GetFileHash Get takes a file and returns the FileHash struct
func GetFileHash(fileLocation string) (*FileHash, error) {
	var fileHash = FileHash{}

	var err error

	file, err := os.Open(fileLocation)
	if err != nil {
		return &fileHash, err
	}

	defer file.Close()

	fileInfo, _ := file.Stat()
	fileHash.Filename = fileInfo.Name()
	fileHash.Size = fileInfo.Size()

	hash := md5.New()

	blocks := uint64(math.Ceil(float64(fileHash.Size) / float64(fileChunk)))

	for i := uint64(0); i < blocks; i++ {
		blocksize := int(math.Min(fileChunk, float64(fileHash.Size-int64(i*fileChunk))))
		buf := make([]byte, blocksize)

		file.Read(buf)
		hash.Write(buf) // append into the hash
	}

	fileHash.MD5 = hex.EncodeToString(hash.Sum(nil))

	return &fileHash, nil
}
