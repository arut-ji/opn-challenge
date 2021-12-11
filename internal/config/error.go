package config

import (
	"fmt"
	"strings"
)

type MissingRequiredFlagsError struct {
	isPublicKeyMissing bool
	isSecretKeyMissing bool
}

func (e *MissingRequiredFlagsError) Error() string {
	missingParams := make([]string, 0)
	if e.isSecretKeyMissing {
		missingParams = append(missingParams, "secretKey")
	}

	if e.isPublicKeyMissing {
		missingParams = append(missingParams, "publicKey")
	}
	return fmt.Sprintf("missing required flags: %v", strings.Join(missingParams, ","))
}

type FileDoesNotExistError struct {
	FilePath string
}

func (e *FileDoesNotExistError) Error() string {
	return fmt.Sprintf("incorrect file path: %v does not exist", e.FilePath)
}

type IncorrectFileFormatError struct {
	FilePath string
}

func (e *IncorrectFileFormatError) Error() string {
	return fmt.Sprintf("incorrect file format: %v is not a supported format, please pass a CSV file", e.FilePath)
}
