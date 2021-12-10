package config

import (
	"errors"
	"flag"
	"os"
	"regexp"
)

type Config struct {
	OmiseClientConfig OmiseClientConfig
	FileSourceConfig  FileSourceConfig
}

type OmiseClientConfig struct {
	OmisePublicKey string
	OmiseSecretKey string
}

type FileSourceConfig struct {
	FilePath string
}

func TryGetConfig() (*Config, error) {
	pubKeyPtr := flag.String("publicKey", "", "Omise public key")
	secretKeyPtr := flag.String("secretKey", "", "Omise secret Key")

	flag.Parse()

	if *pubKeyPtr == "" || *secretKeyPtr == "" {
		return nil, &MissingRequiredFlagsError{
			isPublicKeyMissing: *pubKeyPtr == "",
			isSecretKeyMissing: *secretKeyPtr == "",
		}
	}

	filePath := flag.Args()[0]
	err := validateFileExtension(filePath)
	if err != nil {
		return nil, err
	}

	err = validateFileExistence(filePath)
	if err != nil {
		return nil, err
	}

	return &Config{
		OmiseClientConfig: OmiseClientConfig{
			OmisePublicKey: *pubKeyPtr,
			OmiseSecretKey: *secretKeyPtr,
		},
		FileSourceConfig: FileSourceConfig{
			FilePath: filePath,
		},
	}, nil
}

func validateFileExtension(filePath string) error {
	match, err := regexp.Match("^.*\\.(csv)$", []byte(filePath))
	if err != nil {
		return err
	}
	if match {
		return nil
	} else {
		return &IncorrectFileFormatError{filePath}
	}
}

func validateFileExistence(filePath string) error {
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		return &FileDoesNotExistError{
			FilePath: filePath,
		}
	}
	return nil
}
