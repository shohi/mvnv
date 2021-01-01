package check

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"hash"
	"io"
	"os"
	"strings"

	"github.com/shohi/mvnv/pkg/download"
)

var (
	// ErrUnsupportedChecksumAlgorithm 不支持的校验和算法
	ErrUnsupportedChecksumAlgorithm = errors.New("unsupported checksum algorithm")
	// ErrChecksumNotMatched 校验和不匹配
	ErrChecksumNotMatched = errors.New("file checksum does not match the computed checksum")
)

const (
	// SHA1 校验和算法-sha1
	SHA1 = "SHA1"
	// SHA256 校验和算法-sha256
	SHA256 = "SHA256"
	// SHA512 校验和算法-sha512
	SHA512 = "SHA512"
)

// VerifyChecksum checks whether SHA1 sum of the content of filename equals the
// one posted on the sumURL.
func VerifyChecksum(filename string, sumURL string, algorithm string) (bool, error) {
	f, err := os.Open(filename)
	if err != nil {
		return false, err
	}
	defer f.Close()

	var h hash.Hash
	switch algorithm {
	case SHA1:
		h = sha1.New()
		sumURL += ".sha1"
	case SHA256:
		h = sha256.New()
		sumURL += ".sha256"
	case SHA512:
		h = sha512.New()
		sumURL += ".sha512"
	default:
		return false, ErrUnsupportedChecksumAlgorithm
	}

	if _, err := io.Copy(h, f); err != nil {
		return false, err
	}

	remoteChecksum, err := download.Get(sumURL)
	remoteChecksum = strings.Split(remoteChecksum, " ")[0]
	if err != nil {
		return false, fmt.Errorf("failed to retrieve remote checksum: %v", err)
	}

	// fileChecksum := fmt.Sprintf("%x", h.Sum(nil))
	fileChecksum := hex.EncodeToString(h.Sum(nil))

	if remoteChecksum != fileChecksum {
		// TODO: add debug log?
		// fmt.Println("url: ", sumURL)
		// fmt.Println("remoteChecksum: ", remoteChecksum, " local:", fileChecksum)
		return false, nil
	}

	return true, nil
}
