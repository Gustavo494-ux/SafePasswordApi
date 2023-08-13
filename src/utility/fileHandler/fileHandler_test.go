package fileHandler_test

import (
	// "fmt"
	"io/ioutil"
	"os"
	"path/filepath"
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
	err := fileHandler.CreateFile(dir, fileName)
	if err != nil {
		t.Errorf("CreateFile failed with error: %s", err)
	}
	os.Remove(fileName)
}

func TestOpenFile(t *testing.T) {
	dir, _ := os.Getwd()
	expectedContent := "Test file content"
	err := ioutil.WriteFile(fileName, []byte(expectedContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write test file: %s", err)
	}
	defer os.Remove(fileName)

	content, err := fileHandler.OpenFile(dir, fileName)
	if err != nil {
		t.Errorf("OpenFile failed with error: %s", err)
	}

	if content != expectedContent {
		t.Errorf("OpenFile returned incorrect content. Expected: %s, Got: %s", expectedContent, content)
	}
}

func TestWriteFile(t *testing.T) {
	dir, _ := os.Getwd()
	content := "Test file content"
	err := fileHandler.WriteFile(dir, fileName, content)
	if err != nil {
		t.Errorf("WriteFile failed with error: %s", err)
	}
	defer os.Remove(fileName)

	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		t.Fatalf("Failed to read test file: %s", err)
	}

	if string(fileContent) != content {
		t.Errorf("WriteFile did not write the expected content. Expected: %s, Got: %s", content, string(fileContent))
	}
}

func TestAppendToFile(t *testing.T) {
	initialContent := "Initial content"
	appendedContent := "Appended content"

	directoryPath, _ := os.Getwd()
	// directoryPath = fmt.Sprintf("%s/fileHandler", directoryPath)
	fullPath := strings.ReplaceAll(filepath.Join(directoryPath, fileName), "\\", "/")
	directoryPath = directoryPath + "\\"

	err := ioutil.WriteFile(fullPath, []byte(initialContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write test file: %s", err)
	}
	defer os.Remove(fullPath)

	err = fileHandler.AppendToFile(directoryPath, fileName, appendedContent)
	if err != nil {
		t.Errorf("AppendToFile failed with error: %s", err)
	}

	fileContent, err := ioutil.ReadFile(fullPath)
	if err != nil {
		t.Fatalf("Failed to read test file: %s", err)
	}

	expectedContent := initialContent + appendedContent
	if string(fileContent) != expectedContent {
		t.Errorf("AppendToFile did not append the expected content. Expected: %s, Got: %s", expectedContent, string(fileContent))
	}
}

func TestDeleteFile(t *testing.T) {
	dir, _ := os.Getwd()
	fullPath := filepath.Join(dir, fileName)
	file, err := os.Create(fullPath)
	if err != nil {
		t.Errorf("CreateFile failed with error: %s", err)
	}
	file.Close()
	fileHandler.DeleteFile(dir, fileName)

	_, err = os.Stat(fileName)
	if !os.IsNotExist(err) {
		t.Errorf("DeleteFile did not delete the file as expected")
	}
}

func TestRenameFile(t *testing.T) {
	dir, _ := os.Getwd()
	newFileName := "newfile.txt"

	fullPath_initial := strings.ReplaceAll(filepath.Join(dir, fileName), "\\", "/")
	fullPath_rename := strings.ReplaceAll(filepath.Join(dir, newFileName), "\\", "/")

	err := ioutil.WriteFile(fullPath_initial, []byte("Test file content"), 0644)
	if err != nil {
		t.Fatalf("Failed to write test file: %s", err)
	}
	defer os.Remove(newFileName)

	err = fileHandler.RenameFile(dir, fileName, newFileName)
	if err != nil {
		t.Errorf("RenameFile failed with error: %s", err)
	}

	_, err = os.Stat(fullPath_initial)
	if !os.IsNotExist(err) {
		t.Errorf("RenameFile did not rename the file as expected")
	}

	_, err = os.Stat(fullPath_rename)
	if os.IsNotExist(err) {
		t.Errorf("RenameFile did not create the renamed file as expected")
	}
}

func TestGetFileList(t *testing.T) {
	// Create a temporary directory for testing purposes
	tempDir, err := ioutil.TempDir("", "test_directory")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test files inside the directory
	testFiles := []string{"file1.txt", "file2.txt", "file3.txt"}
	for _, filename := range testFiles {
		filePath := filepath.Join(tempDir, filename)
		file, err := os.Create(filePath)
		if err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}
		defer file.Close()
	}

	// Execute the GetFileList function on the test directory
	fileList, err := fileHandler.GetFileList(tempDir)
	if err != nil {
		t.Fatalf("Error getting file list: %v", err)
	}

	// Check if all the test files are present in the returned file list
	for _, filename := range testFiles {
		found := false
		for _, file := range fileList {
			if file == filename {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected file missing in the returned file list: %s", filename)
		}
	}

	// Check if there are no extra files in the returned file list
	if len(fileList) != len(testFiles) {
		t.Errorf("Number of returned files is different than expected. Expected: %d, Returned: %d", len(testFiles), len(fileList))
	}
}

func TestCreateDirectory(t *testing.T) {
	// Specify the base path for the new directories
	basePath := "./src/utility/teste"

	// Create a slice of folder names
	folders := []string{"folder1", "folder2", "folder3", "folder4", "folder5", "folder6", "folder7", "folder8", "folder9", "folder10"}

	// Slice to store paths of created directories
	createdDirectories := []string{}

	// Iterate over the folders slice
	for _, folder := range folders {
		// Construct the full path for each directory
		path := filepath.Join(basePath, folder)

		// Call the CreateDirectory function
		err := fileHandler.CreateDirectory(path)
		if err != nil {
			t.Errorf("Failed to create directory: %v", err)
		} else {
			// Append the path to the created directories slice
			createdDirectories = append(createdDirectories, path)
		}
	}

	sliceBasePath := strings.Split(basePath, "/")
	err := os.RemoveAll("./" + sliceBasePath[1])
	if err != nil {
		t.Errorf("Failed to remove directory: %v", err)
	}
}

func TestGetFileInfo(t *testing.T) {
	// Test for a file
	filePath := "./src/utility/file.txt"
	fileInfo, err := fileHandler.GetFileInfo(filePath)
	if err != nil {
		t.Errorf("Failed to get file info: %v", err)
	}

	if fileInfo != nil {
		// Check if it's a file
		if !fileInfo.Mode().IsRegular() {
			t.Errorf("Expected a file, got directory")
		}
	}

	// Test for a directory
	dirPath := "./src/utility"
	dirInfo, err := fileHandler.GetFileInfo(dirPath)
	if err != nil {
		t.Errorf("Failed to get directory info: %v", err)
	}

	if dirInfo != nil {
		// Check if it's a directory
		if !dirInfo.Mode().IsDir() {
			t.Errorf("Expected a directory, got file")
		}
	}
}
