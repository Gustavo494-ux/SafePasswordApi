package fileHandler_test

import (
	"fmt"
	"os"
	"path/filepath"
	"safePasswordApi/src/modules/logger"
	"safePasswordApi/src/utility/fileHandler"
	"strings"
	"testing"
)

const (
	fileName      = "testfile.txt"
	directoryName = "directoryTest"
)

func TestCreateFile(t *testing.T) {
	dir, _ := os.Getwd()
	_, err := fileHandler.CreateFile(dir, fileName)
	if err != nil {
		t.Errorf("Falha ao criar o arquivo: %s", err)
		logger.Logger().Error("Falha ao criar o arquivo", err)
	}
	os.Remove(fileName)
	logger.Logger().Info("Teste TestCreateFile executado com sucesso!")
}

func TestOpenFile(t *testing.T) {
	dir, _ := os.Getwd()
	expectedContent := "Conteúdo do arquivo de teste"
	err := os.WriteFile(fileName, []byte(expectedContent), 0644)
	if err != nil {
		t.Fatalf("Falha ao escrever o arquivo de teste: %s", err)
		logger.Logger().Error("Falha ao escrever o arquivo de teste", err)
	}
	defer os.Remove(fileName)

	content, err := fileHandler.OpenFile(dir, fileName)
	if err != nil {
		t.Errorf("Falha ao abrir o arquivo: %s", err)
		logger.Logger().Error("Falha ao abrir o arquivo", err)
	}

	if content != expectedContent {
		t.Errorf("OpenFile retornou conteúdo incorreto. Esperado: %s, Obtido: %s", expectedContent, content)
		logger.Logger().Error(fmt.Sprintf("OpenFile retornou conteúdo incorreto. Esperado: %s, Obtido: %s", expectedContent, content), err)
	}

	logger.Logger().Info("Teste TestOpenFile executado com sucesso!")
}

func TestWriteFile(t *testing.T) {
	dir, _ := os.Getwd()
	content := "Conteúdo do arquivo de teste"
	err := fileHandler.WriteFile(dir, fileName, content)
	if err != nil {
		t.Errorf("Falha ao escrever no arquivo: %s", err)
		logger.Logger().Error("Falha ao escrever no arquivo", err)
	}
	defer os.Remove(fileName)

	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		t.Fatalf("Falha ao ler o arquivo de teste: %s", err)
		logger.Logger().Error("Falha ao ler o arquivo de teste", err)
	}

	if string(fileContent) != content {
		t.Errorf("WriteFile não escreveu o conteúdo esperado. Esperado: %s, Obtido: %s", content, string(fileContent))
		logger.Logger().Error(fmt.Sprintf("WriteFile não escreveu o conteúdo esperado. Esperado: %s, Obtido: %s", content, string(fileContent)), err)
	}

	logger.Logger().Info("Teste TestWriteFile executado com sucesso!")
}

func TestAppendToFile(t *testing.T) {
	initialContent := "Conteúdo inicial"
	appendedContent := "Conteúdo anexado"

	directoryPath, _ := os.Getwd()
	fullPath := strings.ReplaceAll(filepath.Join(directoryPath, fileName), "\\", "/")
	directoryPath = directoryPath + "\\"

	err := os.WriteFile(fullPath, []byte(initialContent), 0644)
	if err != nil {
		t.Fatalf("Falha ao escrever o arquivo de teste: %s", err)
		logger.Logger().Error("Falha ao escrever o arquivo de teste", err)
	}
	defer os.Remove(fullPath)

	err = fileHandler.AppendToFile(directoryPath, fileName, appendedContent)
	if err != nil {
		t.Errorf("Falha ao anexar ao arquivo: %s", err)
		logger.Logger().Error("Falha ao anexar ao arquivo", err)
	}

	fileContent, err := os.ReadFile(fullPath)
	if err != nil {
		t.Fatalf("Falha ao ler o arquivo de teste: %s", err)
		logger.Logger().Error("Falha ao ler o arquivo de teste", err)
	}

	expectedContent := initialContent + appendedContent
	if string(fileContent) != expectedContent {
		t.Errorf("AppendToFile não anexou o conteúdo esperado. Esperado: %s, Obtido: %s", expectedContent, string(fileContent))
		logger.Logger().Error(fmt.Sprintf("AppendToFile não anexou o conteúdo esperado. Esperado: %s, Obtido: %s", expectedContent, string(fileContent)), err)
	}

	logger.Logger().Info("Teste TestAppendToFile executado com sucesso!")
}

func TestDeleteFile(t *testing.T) {
	dir, _ := os.Getwd()
	fullPath := filepath.Join(dir, fileName)
	file, err := os.Create(fullPath)
	if err != nil {
		t.Errorf("Falha ao criar o arquivo: %s", err)
		logger.Logger().Error("Falha ao criar o arquivo", err)
	}
	file.Close()
	err = fileHandler.DeleteFile(dir, fileName)
	if err != nil {
		t.Errorf("Erro ao excluir o arquivo: %s", err)
		logger.Logger().Error("Erro ao excluir o arquivo", err)
	}

	_, err = os.Stat(fileName)
	if !os.IsNotExist(err) {
		t.Errorf("DeleteFile não excluiu o arquivo como esperado")
		logger.Logger().Error("DeleteFile não excluiu o arquivo como esperado", err)
	}

	logger.Logger().Info("Teste TestDeleteFile executado com sucesso!")
}

func TestRenameFile(t *testing.T) {
	dir, _ := os.Getwd()
	newFileName := "newfile.txt"

	fullPathInitial := strings.ReplaceAll(filepath.Join(dir, fileName), "\\", "/")
	fullPathRename := strings.ReplaceAll(filepath.Join(dir, newFileName), "\\", "/")

	err := os.WriteFile(fullPathInitial, []byte("Conteúdo do arquivo de teste"), 0644)
	if err != nil {
		t.Fatalf("Falha ao escrever o arquivo de teste: %s", err)
		logger.Logger().Error("Falha ao escrever o arquivo de teste", err)
	}
	defer os.Remove(newFileName)

	err = fileHandler.RenameFile(dir, fileName, newFileName)
	if err != nil {
		t.Errorf("Falha ao renomear o arquivo: %s", err)
		logger.Logger().Error("Falha ao renomear o arquivo", err)
	}

	_, err = os.Stat(fullPathInitial)
	if !os.IsNotExist(err) {
		t.Errorf("RenameFile não renomeou o arquivo como esperado")
		logger.Logger().Error("RenameFile não renomeou o arquivo como esperado", err)
	}

	_, err = os.Stat(fullPathRename)
	if os.IsNotExist(err) {
		t.Errorf("RenameFile não criou o arquivo renomeado como esperado")
		logger.Logger().Error("RenameFile não criou o arquivo renomeado como esperado", err)
	}

	logger.Logger().Info("Teste TestRenameFile executado com sucesso!")
}

func TestGetFileList(t *testing.T) {
	// Criar um diretório temporário para fins de teste
	currentDirectory, err := os.Getwd()
	if err != nil {
		t.Fatalf("Erro ao buscar o caminho do diretório atual, erro: %s", err)
		logger.Logger().Error("Erro ao buscar o caminho do diretório atual, erro", err)
	}

	fullPath := strings.ReplaceAll(filepath.Join(currentDirectory, "test_directory"), "\\", "/")

	err = os.MkdirAll(fullPath, 0750)
	if err != nil {
		t.Fatalf("Falha ao criar o diretório temporário: %v", err)
		logger.Logger().Error("Falha ao criar o diretório temporário", err)
	}
	defer os.RemoveAll("test_directory")

	// Criar arquivos de teste dentro do diretório
	testFiles := []string{"file1.txt", "file2.txt", "file3.txt"}
	for _, filename := range testFiles {
		filePath := filepath.Join(fullPath, filename)
		file, err := os.Create(filePath)
		if err != nil {
			t.Fatalf("Falha ao criar o arquivo de teste: %v", err)
			logger.Logger().Error("Falha ao criar o arquivo de teste", err)
		}
		defer file.Close()
	}

	// Executar a função GetFileList no diretório de teste
	fileList, err := fileHandler.GetFileList(fullPath)
	if err != nil {
		t.Fatalf("Erro ao obter a lista de arquivos: %v", err)
		logger.Logger().Error("Erro ao obter a lista de arquivos", err)
	}

	// Verificar se todos os arquivos de teste estão presentes na lista de arquivos retornada
	for _, filename := range testFiles {
		found := false
		for _, file := range fileList {
			if file == filename {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Arquivo esperado ausente na lista de arquivos retornada: %s", filename)
			logger.Logger().Error("Arquivo esperado ausente na lista de arquivos retornada", err)
		}
	}

	// Verificar se não há arquivos extras na lista de arquivos retornada
	if len(fileList) != len(testFiles) {
		t.Errorf("O número de arquivos retornados é diferente do esperado. Esperado: %d, Retornado: %d", len(testFiles), len(fileList))
		logger.Logger().Error(fmt.Sprintf("O número de arquivos retornados é diferente do esperado. Esperado: %d, Retornado: %d", len(testFiles), len(fileList)), err)
	}

	logger.Logger().Info("Teste TestGetFileList executado com sucesso!")
}

func TestCreateDirectory(t *testing.T) {
	// Especificar o caminho base para os novos diretórios
	basePath := "./src/utility/teste"

	// Criar uma lista de nomes de pastas
	pastas := []string{"pasta1", "pasta2", "pasta3", "pasta4", "pasta5", "pasta6", "pasta7", "pasta8", "pasta9", "pasta10"}

	// Iterar sobre a lista de pastas
	for _, pasta := range pastas {
		// Construir o caminho completo para cada diretório
		caminho := filepath.Join(basePath, pasta)

		// Chamar a função CreateDirectory
		err := fileHandler.CreateDirectory(caminho)
		if err != nil {
			t.Errorf("Falha ao criar o diretório: %v", err)
			logger.Logger().Error("Falha ao criar o diretório", err)
		}
	}

	sliceBasePath := strings.Split(basePath, "/")
	err := os.RemoveAll("./" + sliceBasePath[1])
	if err != nil {
		t.Errorf("Falha ao remover o diretório: %v", err)
		logger.Logger().Error("Falha ao remover o diretório", err)
	}

	logger.Logger().Info("Teste TestCreateDirectory executado com sucesso!")
}

func TestGetFileInfo(t *testing.T) {
	// Teste para um arquivo
	caminhoDoArquivo := "./src/utility/file.txt"
	infoDoArquivo, err := fileHandler.GetFileInfo(caminhoDoArquivo)
	if err != nil {
		t.Errorf("Falha ao obter informações do arquivo: %v", err)
		logger.Logger().Error("Falha ao obter informações do arquivo", err)
	}

	if infoDoArquivo != nil {
		// Verificar se é um arquivo
		if !infoDoArquivo.Mode().IsRegular() {
			t.Errorf("Esperado um arquivo, obtido um diretório")
			logger.Logger().Error("Esperado um arquivo, obtido um diretório", err)
		}
	}

	// Teste para um diretório
	caminhoDoDiretório := "./src/utility"
	infoDoDiretório, err := fileHandler.GetFileInfo(caminhoDoDiretório)
	if err != nil {
		t.Errorf("Falha ao obter informações do diretório: %v", err)
		logger.Logger().Error("Falha ao obter informações do diretório", err)
	}

	if infoDoDiretório != nil {
		// Verificar se é um diretório
		if !infoDoDiretório.Mode().IsDir() {
			t.Errorf("Esperado um diretório, obtido um arquivo")
			logger.Logger().Error("Esperado um diretório, obtido um arquivo", err)
		}
	}

	logger.Logger().Info("Teste TestGetFileInfo executado com sucesso!")
}
