package fileHandler

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// CreateFile cria um novo arquivo com o nome especificado
func CreateFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

// OpenFile abre um arquivo existente para leitura
func OpenFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// WriteFile grava o conteúdo fornecido em um arquivo
func WriteFile(filename string, content string) error {
	err := ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

// AppendToFile anexa o conteúdo fornecido a um arquivo existente
func AppendToFile(filename string, content string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
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

// DeleteFile exclui um arquivo
func DeleteFile(filename string) error {
	err := os.Remove(filename)
	if err != nil {
		return err
	}
	return nil
}

// RenameFile renomeia um arquivo
func RenameFile(oldFilename, newFilename string) error {
	err := os.Rename(oldFilename, newFilename)
	if err != nil {
		return err
	}
	return nil
}

// GetFileList retorna a lista de arquivos no diretório especificado
func GetFileList(directory string) ([]string, error) {
	fileList := []string{}

	files, err := ioutil.ReadDir(directory)
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

// CreateDirectory creates a directory at the specified path
func CreateDirectory(path string) error {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return fmt.Errorf("error creating directory: %v", err)
	}
	return nil
}

// GetFileInfo returns information about a file or directory specified by the given path.
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
