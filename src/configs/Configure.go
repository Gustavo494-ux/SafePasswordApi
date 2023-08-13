package configs

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
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
	RSAPrivateKey     string
	RSAPublicKey      string
	AESKey            string
	RSAPrivateKeyPath string
	RSAPublicKeyPath  string
	AESKeyPath        string
	SecretKeyJWTPath  string
)

// InitializeConfigurations : Performs the necessary setup for the project to be used
func InitializeConfigurations(Path string) {
	loadEnvironmentVariables(Path)
	loadOrCreateKeys()
}

// loadEnvironmentVariables : Initializes the environment variables
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

// loadOrCreateKeys : Loads the keys that will be used during the execution of the project. if the keys do not exist they will be created.
func loadOrCreateKeys() {
	go LoadKey(&AESKey, AESKeyPath, ValidateAESKey)
	func() {
		LoadKey(&RSAPrivateKey, RSAPrivateKeyPath, ValidateRSAPrivateKey)
		LoadKey(&RSAPublicKey, RSAPublicKeyPath, ValidateKeyPublicaRSA)
	}()
	go LoadKey(nil, SecretKeyJWTPath, LoadSecretKeyJWT)
}

// LoadSecretKeyJWT : Loads the Secret Key JWT variable of type byte
func LoadSecretKeyJWT(Key *string, Path string) {
	fileHandler.CreateDirectoryOrFileIfNotExists(Path)

	JWT_Key := string(SecretKeyJWT)
	LoadKey(&JWT_Key, SecretKeyJWTPath, ValidateSecretKeyJWT)
	SecretKeyJWT = []byte(JWT_Key)
}

// LoadKey : Load a variable from data from a file.
func LoadKey(Key *string, Path string, proximas ...func(key *string, path string)) {
	if Key != nil && len(*Key) == 0 {
		dir, fileName := filepath.Split(Path)
		fileHandler.CreateDirectoryOrFileIfNotExists(Path)
		var err error
		*Key, err = fileHandler.OpenFile(dir, fileName)
		if err != nil {
			log.Fatal("Error opening file: ", err)
		}
	}
	for _, proxima_funcao := range proximas {
		proxima_funcao(Key, Path)
	}
}

// writeQueryAndCheckFileData : Writes the data to the file, reads the same file and checks if the data was written
func writeQueryAndCheckFileData(data string, path string) {
	dir, fileName := filepath.Split(path)
	err := fileHandler.WriteFile(dir, fileName, data)
	if err != nil {
		log.Fatal("Invalid key, please check: ", path)
	}

	Key, err := fileHandler.OpenFile(dir, fileName)
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}

	if len(Key) == 0 {
		log.Fatal("Invalid key, please check: ", path)
	}
}

// ValidateAESKey : Performs AES key validation
func ValidateAESKey(key *string, path string) {
	var err error
	switch {
	case len(*key) == 0:
		{
			*key, err = symmetricEncryp.GenerateRandomAESKey()
			if err != nil {
				log.Fatal("Error generate AES KEY, err: ", err)
			}
			writeQueryAndCheckFileData(*key, AESKeyPath)
			LoadKey(key, AESKeyPath, ValidateAESKey)
		}
	case len(*key) == 32:
	case len(*key) >= 32:
		*key = string(*key)[:32]
	default:
		log.Fatal("AES Key invalida")
	}
}

// ValidateRSAPrivateKey : Performs RSA Private key validation
func ValidateRSAPrivateKey(key *string, path string) {
	err := asymmetrical.ValidatePrivateKey(*key)
	if err != nil {
		PrivateKey, err := asymmetrical.GeneratePrivateKey(2048)
		if err != nil {
			log.Fatal("Error generating RSA private key, please check: ", RSAPrivateKeyPath)
		}

		*key, err = asymmetrical.ExportPrivateKey(PrivateKey)
		if err != nil {
			log.Fatal("Error generating RSA private key, please check: ", RSAPrivateKeyPath)
		}
		writeQueryAndCheckFileData(RSAPrivateKey, path)
		LoadKey(key, path, ValidateRSAPrivateKey)
	}
}

// ValidateRSAPrivateKey : Performs RSA Public key validation
func ValidateKeyPublicaRSA(Key *string, Path string) {
	err := asymmetrical.ValidatePublicKey(*Key)
	if err != nil {
		err = asymmetrical.ValidatePrivateKey(RSAPrivateKey)
		if err != nil {
			log.Fatal("Invalid RSA Private key, please check: ", Path)
		}

		PublicKey, err := asymmetrical.GeneratePublicKey(RSAPrivateKey)
		if err != nil {
			log.Fatal("Error generating RSA public KEY, please check: ", Path)
		}

		*Key, err = asymmetrical.ExportPublicKey(PublicKey)
		if err != nil {
			log.Fatal("Error generating RSA public KEY, please check: ", Path)
		}

		writeQueryAndCheckFileData(*Key, Path)
		LoadKey(Key, Path, ValidateKeyPublicaRSA)
	}
}

// ValidateSecretKeyJWT : Performs SecretKeyJWT validation
func ValidateSecretKeyJWT(Key *string, Path string) {
	if len(strings.Trim(*Key, " ")) == 0 {
		RandomAESKey, err := symmetricEncryp.GenerateRandomAESKey()
		if err != nil {
			log.Fatal("Error generate Secret KEY, err: ", err)
		}
		RandomAESKeyHash, err := hashEncrpt.GenerateSHA512(RandomAESKey)
		if err != nil {
			log.Fatal("Error generate Secret KEY, err: ", err)
		}

		*Key, err = hashEncrpt.GenerateSHA512(fmt.Sprintf("%s,%s,%d", RandomAESKey, RandomAESKeyHash, time.Now().Unix()))
		if err != nil {
			log.Fatal("Error generate Secret KEY, err: ", err)
		}

		writeQueryAndCheckFileData(*Key, Path)
		LoadKey(Key, Path, ValidateSecretKeyJWT)
	}
}
