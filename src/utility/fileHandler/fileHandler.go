package fileHandler

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// CreateFile : creates a new file with the specified name
func CreateFile(directoryPath string, filename string) (file *os.File, err error) {
	file, err = os.Create(strings.ReplaceAll(filepath.Join(directoryPath, filename), "\\", "/"))
	if err != nil {
		return
	}
	defer file.Close()
	return
}

// OpenFile : open an existing file for reading
func OpenFile(directoryPath string, filename string) (string, error) {
	data, err := os.ReadFile(strings.ReplaceAll(filepath.Join(directoryPath, filename), "\\", "/"))
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// CarregarArquivo: Abre o arquivo e retorna o mesmo
func CarregarArquivo(nomeArquivo string) (file *os.File, err error) {
	file, err = os.OpenFile(nomeArquivo, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	return
}

// WriteFile : writes the given content to a file
func WriteFile(directoryPath string, filename string, content string) error {
	fullPath := strings.ReplaceAll(filepath.Join(directoryPath, filename), "\\", "/")
	err := os.WriteFile(fullPath, []byte(content), 0600)
	if err != nil {
		return err
	}
	return nil
}

// AppendToFile : appends the provided content to an existing file
func AppendToFile(directoryPath string, filename string, content string) error {
	file, err := os.OpenFile(
		strings.ReplaceAll(filepath.Join(directoryPath, filename), "\\", "/"),
		os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

// DeleteFile : delete a file
func DeleteFile(directoryPath string, filename string) error {
	fullPath := strings.ReplaceAll(filepath.Join(directoryPath, filename), "\\", "/")
	err := os.Remove(fullPath)
	if err != nil {
		return err
	}
	return nil
}

// RenameFile : rename a file
func RenameFile(directoryPath string, oldFilename, newFilename string) error {
	fullPathOldFilename := strings.ReplaceAll(filepath.Join(directoryPath, oldFilename), "\\", "/")
	fullPathNewFilename := strings.ReplaceAll(filepath.Join(directoryPath, newFilename), "\\", "/")
	err := os.Rename(fullPathOldFilename, fullPathNewFilename)
	if err != nil {
		return err
	}
	return nil
}

// GetFileList : returns the list of files in the specified directory
func GetFileList(directory string) ([]string, error) {
	fileList := []string{}

	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() {
			fileList = append(fileList, file.Name())
		}
	}

	return fileList, nil
}

// CreateDirectory : creates a directory at the specified path
func CreateDirectory(path string) error {
	err := os.MkdirAll(path, 0750)
	if err != nil {
		return fmt.Errorf("error creating directory: %v", err)
	}
	return nil
}

// GetFileInfo : returns information about a file or directory specified by the given path.
func GetFileInfo(path string) (os.FileInfo, error) {
	// Convert the path to an absolute path if it is relative
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("error resolving absolute path: %v", err)
	}

	// Check if the file or directory exists
	_, err = os.Stat(absPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil // Return nil when the directory is not found
		}
		return nil, fmt.Errorf("error getting file info: %v", err)
	}

	// Retrieve the file info
	fileInfo, err := os.Stat(absPath)
	if err != nil {
		return nil, fmt.Errorf("error getting file info: %v", err)
	}

	return fileInfo, nil
}

// CreateDirectoryIfNotExists : Verifica se o Diretorio Existe, caso não exista o mesmo será criado
func CreateDirectoryIfNotExists(path string) (err error) {
	dirInfo, err := GetFileInfo(path)
	if err != nil {
		err = fmt.Errorf("error getting directory info: %s", err)
	}

	if dirInfo == nil {
		err = CreateDirectory(path)
		if err != nil {
			err = fmt.Errorf("error creating directory: %s", err)
		}
	}
	return
}

// CreateFileIfNotExists : Verifica se o arquivo Existe, caso não exista o mesmo será criado
func CreateFileIfNotExists(path string) (err error) {
	dir, fileName := filepath.Split(path)
	fileInfo, err := GetFileInfo(path)
	if err != nil {
		err = fmt.Errorf("error getting file info: %s", err)
	}
	if fileInfo == nil {
		_, err = CreateFile(dir, fileName)
		if err != nil {
			err = fmt.Errorf("error creating file: %s", err)
		}
	}
	return
}

// GetDirectoryPath : Receive the path of a file and extract the path of the directory where this file will be created
func GetDirectoryPath(Path string) string {
	dirPath := strings.Split(strings.ReplaceAll(Path, "\\", "/"), "/")
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

// CreateDirectoryOrFileIfNotExists : It receives the path of a file and if it doesn't exist it will create all the directories and the file itself.
func CreateDirectoryOrFileIfNotExists(path string) (err error) {
	err = CreateDirectoryIfNotExists(GetDirectoryPath(path))
	if err != nil {
		err = fmt.Errorf("error CreateDirectoryIfNotExists : %s", err)
		return
	}

	err = CreateFileIfNotExists(path)
	if err != nil {
		err = fmt.Errorf("error CreateFileIfNotExists : %s", err)
		return
	}
	return
}

// GetPathUntilFolder : retorna o caminho até o nível da pasta desejada em um caminho completo.
func GetPathUntilFolder(path string, folderName string) (string, error) {
	// Procurar a posição da última ocorrência do nome da pasta no caminho
	index := strings.LastIndex(strings.ToLower(path), folderName)
	if index == -1 {
		return "", fmt.Errorf("a pasta '%s' não foi encontrada no caminho '%s'", folderName, path)
	}

	// Obter o caminho até a última ocorrência do nome da pasta
	dirUntilFolder := path[:index+len(folderName)]
	return dirUntilFolder, nil
}

// GetSourceDirectory : returns the absolute path to the root directory of the projectreturns the absolute path to the root directory of the project
func GetSourceDirectory(RootDirectory string) (rootDirectoryPath string, err error) {
	RootDirectory = strings.ToLower(RootDirectory)
	currentDirectoryPath, err := os.Getwd()
	if err != nil {
		return
	}
	rootDirectoryPath, err = GetPathUntilFolder(fmt.Sprintf("%s%s", currentDirectoryPath, string(filepath.Separator)), RootDirectory)
	if err != nil {
		rootDirectoryPath, err = GetSourceDirectory("app")
		return
	}
	return
}

// GetAbsoluteOrRootConcatenatedPath : function to check if the given path is absolute. case will not be set to the root directory of the project + the last directory of the given path
func GetAbsoluteOrRootConcatenatedPath(Path string, RootDirectory string) (string, error) {
	RootDirectory = strings.ToLower(RootDirectory)
	Path = filepath.FromSlash(Path)
	if filepath.IsAbs(Path) {
		return Path, nil
	}

	RootDirectory, err := GetSourceDirectory(RootDirectory)
	if err != nil {
		return "", err
	}

	// Get last directory and file name from path
	dir, file := filepath.Split(Path)
	dir = strings.TrimSuffix(dir, string(filepath.Separator))
	file = strings.TrimPrefix(file, string(filepath.Separator))

	// Concatenate the last directory and filename with the root path
	absolutePath := filepath.Join(RootDirectory, dir, file)
	return absolutePath, nil
}
