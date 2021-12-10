package config

import "fmt"

type MissingRequiredFlagsError struct {
	isPublicKeyMissing bool
	isSecretKeyMissing bool
}

func (e *MissingRequiredFlagsError) Error() string {
	missingParams := ""
	if e.isSecretKeyMissing {
		missingParams += "secretKey"
	}

	if e.isPublicKeyMissing {
		missingParams += "publicKey"
	}
	return fmt.Sprintf("missing required flags: %v", missingParams)
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
