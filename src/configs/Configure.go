package configs

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"safePasswordApi/src/security/encrypt/asymmetrical"
	hashEncrpt "safePasswordApi/src/security/encrypt/hash"
	symmetricEncryp "safePasswordApi/src/security/encrypt/symmetrical"
	"safePasswordApi/src/utility/fileHandler"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

var (
	StringConnection  = ""
	Port              = 0
	SecretKeyJWT      []byte
	RSAPrivateKeyPath string
	RSAPublicKeyPath  string
	AESKeyPath        string
	SecretKeyJWTPath  string
	RSAPrivateKey     string
	RSAPublicKey      string
	AESKey            string
)

// InitializeConfigurations performs the necessary setup for the project to be used
func InitializeConfigurations(Path string) {
	loadEnvironmentVariables(Path)
	loadOrCreateKeys()
}

// loadOrCreateKeys loads and uses keys or creates keys used in project encryption
func loadOrCreateKeys() {
	loadOrCreateAESKey()
	loadOrCreateRSAPrivateKey()
	loadOrCreateRSAPublicKey()
	loadOrCreateSecretKeyJWT()
}

// loadEnvironmentVariables initializes the environment variables
func loadEnvironmentVariables(Path string) {
	err := godotenv.Load(Path)
	if err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 5000
	}

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
	SecretKeyJWTPath = os.Getenv("SECRET_KEY_JWT_PATH")
}

func loadOrCreateAESKey() {
	var err error
	if len(AESKeyPath) == 0 {
		log.Fatal(errors.New("path key AES empty"))
	}

	createDirectoryOrFileIfNotExists(AESKeyPath)

	AESKey, err = fileHandler.OpenFile(AESKeyPath)
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}

	if len(AESKey) > 32 {
		AESKey = AESKey[:32]
	} else {
		AESKey, err = symmetricEncryp.GenerateRandomAESKey()
		if err != nil {
			log.Fatal("Error generate AES KEY, err: ", err)
		}
		writeQueryAndCheckFileData(AESKey, AESKeyPath)
	}
}

func loadOrCreateRSAPrivateKey() {
	var err error
	if len(RSAPrivateKeyPath) == 0 {
		log.Fatal(errors.New("path private key RSA empty"))
	}
	createDirectoryOrFileIfNotExists(RSAPrivateKeyPath)

	RSAPrivateKey, err = fileHandler.OpenFile(RSAPrivateKeyPath)
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}

	err = asymmetrical.ValidatePrivateKey(RSAPrivateKey)
	if err != nil {
		PrivateKey, err := asymmetrical.GeneratePrivateKey(2048)
		if err != nil {
			log.Fatal("Error generating RSA private key, please check: ", RSAPrivateKeyPath)
		}

		RSAPrivateKey, err = asymmetrical.ExportPrivateKey(PrivateKey)
		if err != nil {
			log.Fatal("Error generating RSA private key, please check: ", RSAPrivateKeyPath)
		}

		writeQueryAndCheckFileData(RSAPrivateKey, RSAPrivateKeyPath)

		RSAPrivateKey, err = fileHandler.OpenFile(RSAPrivateKeyPath)
		if err != nil {
			log.Fatal("Error opening file: ", err)
		}

		err = asymmetrical.ValidatePrivateKey(RSAPrivateKey)
		if err != nil {
			log.Fatal("Invalid RSA Private key, please check: ", RSAPrivateKeyPath)
		}
	}
}

func loadOrCreateRSAPublicKey() {
	if len(RSAPublicKeyPath) == 0 {
		log.Fatal(errors.New("path public key RSA empty"))
	}
	dirPathCreate := getDirectoryPath(RSAPublicKeyPath)

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

	fileInfo, err := fileHandler.GetFileInfo(RSAPublicKeyPath)
	if err != nil {
		log.Fatal("Error getting file info: ", err)
	}
	if fileInfo == nil {
		err = fileHandler.CreateFile(RSAPublicKeyPath)
		if err != nil {
			log.Fatal("Error creating file: ", err)
		}
	}

	RSAPublicKey, err = fileHandler.OpenFile(RSAPublicKeyPath)
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}

	err = asymmetrical.ValidatePublicKey(RSAPublicKey)
	if err != nil {
		err = asymmetrical.ValidatePrivateKey(RSAPrivateKey)
		if err != nil {
			log.Fatal("Invalid RSA Private key, please check: ", RSAPrivateKeyPath)
		}

		PublicKey, err := asymmetrical.GeneratePublicKey(RSAPrivateKey)
		if err != nil {
			log.Fatal("Error generating RSA public KEY, please check: ", RSAPublicKeyPath)
		}

		RSAPublicKey, err := asymmetrical.ExportPublicKey(PublicKey)
		if err != nil {
			log.Fatal("Error generating RSA public KEY, please check: ", RSAPublicKeyPath)
		}

		err = fileHandler.WriteFile(RSAPublicKeyPath, RSAPublicKey)
		if err != nil {
			log.Fatal("Invalid RSA public KEY, please check: ", RSAPublicKeyPath)
		}

		RSAPublicKey, err = fileHandler.OpenFile(RSAPublicKeyPath)
		if err != nil {
			log.Fatal("Error opening file: ", err)
		}

		err = asymmetrical.ValidatePublicKey(RSAPublicKey)
		if err != nil {
			log.Fatal("Invalid RSA public KEY, please check: ", RSAPublicKeyPath)
		}
	}
}

func loadOrCreateSecretKeyJWT() {
	if len(SecretKeyJWTPath) == 0 {
		log.Fatal(errors.New("path secret key JWT empty"))
	}

	dirPathCreate := getDirectoryPath(SecretKeyJWTPath)

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

	fileInfo, err := fileHandler.GetFileInfo(SecretKeyJWTPath)
	if err != nil {
		log.Fatal("Error getting file info: ", err)
	}
	if fileInfo == nil {
		err = fileHandler.CreateFile(SecretKeyJWTPath)
		if err != nil {
			log.Fatal("Error creating file: ", err)
		}
	}

	SecretKeyJWTString, err := fileHandler.OpenFile(SecretKeyJWTPath)
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}

	if SecretKeyJWTString == "" {
		RandomAESKey, err := symmetricEncryp.GenerateRandomAESKey()
		if err != nil {
			log.Fatal("Error generate Secret KEY, err: ", err)
		}
		RandomAESKeyHash, err := hashEncrpt.GenerateSHA512(RandomAESKey)
		if err != nil {
			log.Fatal("Error generate Secret KEY, err: ", err)
		}

		SecretKeyJWTString = randomizeString(fmt.Sprintf("%s,%s", RandomAESKey, RandomAESKeyHash))

		err = fileHandler.WriteFile(SecretKeyJWTPath, SecretKeyJWTString)
		if err != nil {
			log.Fatal("Invalid Secret key, please check: ", SecretKeyJWTPath)
		}

		SecretKeyJWTString, err = fileHandler.OpenFile(SecretKeyJWTPath)
		if err != nil {
			log.Fatal("Error opening file: ", err)
		}
	}

	if SecretKeyJWTString == "" {
		log.Fatal("Invalid Secret key, please check: ", SecretKeyJWTString)
	}
}

func createDirectoryOrFileIfNotExists(path string) {
	createDirectoryIfNotExists(getDirectoryPath(path))
	createFileIfNotExists(path)
}

func writeQueryAndCheckFileData(data string, path string) {
	err := fileHandler.WriteFile(path, data)
	if err != nil {
		log.Fatal("Invalid key, please check: ", path)
	}

	AESKey, err = fileHandler.OpenFile(path)
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}

	if len(AESKey) == 0 {
		log.Fatal("Invalid key, please check: ", path)
	}
}

func createDirectoryIfNotExists(path string) {
	dirInfo, err := fileHandler.GetFileInfo(path)
	if err != nil {
		log.Fatal("Error getting directory info: ", err)
	}

	if dirInfo == nil {
		err = fileHandler.CreateDirectory(path)
		if err != nil {
			log.Fatal("Error creating directory: ", err)
		}
	}
}

func createFileIfNotExists(path string) {
	fileInfo, err := fileHandler.GetFileInfo(path)
	if err != nil {
		log.Fatal("Error getting file info: ", err)
	}
	if fileInfo == nil {
		err = fileHandler.CreateFile(path)
		if err != nil {
			log.Fatal("Error creating file: ", err)
		}
	}
}

func getDirectoryPath(Path string) string {
	dirPath := strings.Split(Path, "/")
	dirPath = append(dirPath[:len(dirPath)-1], dirPath[len(dirPath):]...)
	dirPathCreate := ""
	for i, dir := range dirPath {
		if i > 0 {
			dirPathCreate += "/"
		}
		dirPathCreate += dir
	}
	return dirPathCreate
}

func randomizeString(input string) string {
	// Convert the string to a slice of runes
	runes := []rune(input)

	// Create a new random source with a specific seed
	source := rand.NewSource(time.Now().UnixNano())

	// Create a new random generator using the source
	random := rand.New(source)

	// Shuffle the runes using Fisher-Yates algorithm
	length := len(runes)
	for i := length - 1; i > 0; i-- {
		j := random.Intn(i + 1)
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Convert the slice of runes back to a string
	randomizedString := string(runes)

	return randomizedString
}
