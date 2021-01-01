package base

import "github.com/shohi/mvnv/pkg/check"

func ChecksumAlgorithm(version string) string {
	if version >= "3.6.0" {
		return check.SHA512

	}

	return check.SHA1
}
