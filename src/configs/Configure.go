package configs

import (
	"fmt"
	"log"
	"os"
	"safePasswordApi/src/security/encrypt/asymmetrical"
	"safePasswordApi/src/utility/fileHandler"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var (
	StringConnection  = ""
	Port              = 0
	SecretKey         []byte
	RSAPrivateKeyPath string
	RSAPublicKeyPath  string
	AESKeyPath        string
	RSAPrivateKey     string
	RSAPublicKey      string
	AESKey            string
)

// InitializeConfigurations performs the necessary setup for the project to be used
func InitializeConfigurations() {
	loadEnvironmentVariables()
	loadOrCreateKeys()
}

// loadOrCreateKeys loads and uses keys or creates keys used in project encryption
func loadOrCreateKeys() {
	loadOrCreateAESKey()
	loadOrCreateRSAPrivateKey()
}

// loadEnvironmentVariables initializes the environment variables
func loadEnvironmentVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 5000
	}

	SecretKey = []byte(os.Getenv("SECRET_KEY"))

	StringConnection = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	RSAPrivateKeyPath = os.Getenv("RSA_PRIVATE_KEY_PATH")
	RSAPublicKeyPath = os.Getenv("RSA_PUBLIC_KEY_PATH")
	AESKeyPath = os.Getenv("AES_KEY_PATH")
}

func loadOrCreateAESKey() {
	dirPath := strings.Split(AESKeyPath, "/")
	dirPath = append(dirPath[:len(dirPath)-1], dirPath[len(dirPath):]...)
	dirPathCreate := ""
	for i, dir := range dirPath {
		if i > 0 {
			dirPathCreate += "/"
		}
		dirPathCreate += dir
	}

	dirInfo, err := fileHandler.GetFileInfo(dirPathCreate)
	if err != nil {
		log.Fatal("Error getting directory info: ", err)
	}

	if dirInfo == nil {
		err = fileHandler.CreateDirectory(dirPathCreate)
		if err != nil {
			log.Fatal("Error creating directory: ", err)
		}
	}

	fileInfo, err := fileHandler.GetFileInfo(AESKeyPath)
	if err != nil {
		log.Fatal("Error getting file info: ", err)
	}
	if fileInfo == nil {
		err = fileHandler.CreateFile(AESKeyPath)
		if err != nil {
			log.Fatal("Error creating file: ", err)
		}
	}

	AESKey, err := fileHandler.OpenFile(AESKeyPath)
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}

	if AESKey == "" {
		err = fileHandler.WriteFile(AESKeyPath, "temporaryKey")
		if err != nil {
			log.Fatal("Invalid AES key, please check: ", AESKeyPath)
		}

		AESKey, err = fileHandler.OpenFile(AESKeyPath)
		if err != nil {
			log.Fatal("Error opening file: ", err)
		}
	}

	if AESKey == "" {
		log.Fatal("Invalid AES key, please check: ", AESKeyPath)
	}
}

func loadOrCreateRSAPrivateKey() {
	dirPath := strings.Split(RSAPrivateKeyPath, "/")
	dirPath = append(dirPath[:len(dirPath)-1], dirPath[len(dirPath):]...)
	dirPathCreate := ""
	for i, dir := range dirPath {
		if i > 0 {
			dirPathCreate += "/"
		}
		dirPathCreate += dir
	}

	dirInfo, err := fileHandler.GetFileInfo(dirPathCreate)
	if err != nil {
		log.Fatal("Error getting directory info: ", err)
	}
	if dirInfo == nil {
		err = fileHandler.CreateDirectory(dirPathCreate)
		if err != nil {
			log.Fatal("Error creating directory: ", err)
		}
	}

	fileInfo, err := fileHandler.GetFileInfo(RSAPrivateKeyPath)
	if err != nil {
		log.Fatal("Error getting file info: ", err)
	}
	if fileInfo == nil {
		err = fileHandler.CreateFile(RSAPrivateKeyPath)
		if err != nil {
			log.Fatal("Error creating file: ", err)
		}
	}

	RSAPrivateKey, err := fileHandler.OpenFile(RSAPrivateKeyPath)
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}

	err = asymmetrical.ValidatePrivateKey(RSAPrivateKey)
	if err != nil {
		PrivateKey, err := asymmetrical.GeneratePrivateKey(2048)
		if err != nil {
			log.Fatal("Error generating RSA private key, please check: ", RSAPrivateKeyPath)
		}

		RSAPrivateKey, err := asymmetrical.ExportPrivateKey(PrivateKey)
		if err != nil {
			log.Fatal("Error generating RSA private key, please check: ", RSAPrivateKeyPath)
		}

		err = fileHandler.WriteFile(RSAPrivateKeyPath, RSAPrivateKey)
		if err != nil {
			log.Fatal("Invalid AES key, please check: ", RSAPrivateKeyPath)
		}

		RSAPrivateKey, err = fileHandler.OpenFile(RSAPrivateKeyPath)
		if err != nil {
			log.Fatal("Error opening file: ", err)
		}

		err = asymmetrical.ValidatePrivateKey(RSAPrivateKey)
		if err != nil {
			log.Fatal("Invalid AES key, please check: ", RSAPrivateKeyPath)
		}
	}
}
