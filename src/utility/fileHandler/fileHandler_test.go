package fileHandler_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"safePasswordApi/src/utility/fileHandler"
	"testing"
)

const (
	fileName      = "testfile.txt"
	directoryName = "directoryTest"
	directoryPath = "E:/Projects/go/src/SafePasswordApi/src/test"
)

func TestCreateFile(t *testing.T) {
	err := fileHandler.CreateFile(fileName)
	if err != nil {
		t.Errorf("CreateFile failed with error: %s", err)
	}
	os.Remove(fileName)
}

func TestOpenFile(t *testing.T) {
	expectedContent := "Test file content"
	err := ioutil.WriteFile(fileName, []byte(expectedContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write test file: %s", err)
	}
	defer os.Remove(fileName)

	content, err := fileHandler.OpenFile(fileName)
	if err != nil {
		t.Errorf("OpenFile failed with error: %s", err)
	}

	if content != expectedContent {
		t.Errorf("OpenFile returned incorrect content. Expected: %s, Got: %s", expectedContent, content)
	}
}

func TestWriteFile(t *testing.T) {
	content := "Test file content"
	err := fileHandler.WriteFile(fileName, content)
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

	err := ioutil.WriteFile(fileName, []byte(initialContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write test file: %s", err)
	}
	defer os.Remove(fileName)

	err = fileHandler.AppendToFile(fileName, appendedContent)
	if err != nil {
		t.Errorf("AppendToFile failed with error: %s", err)
	}

	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		t.Fatalf("Failed to read test file: %s", err)
	}

	expectedContent := initialContent + appendedContent
	if string(fileContent) != expectedContent {
		t.Errorf("AppendToFile did not append the expected content. Expected: %s, Got: %s", expectedContent, string(fileContent))
	}
}

func TestDeleteFile(t *testing.T) {
	file, err := os.Create(fileName)
	if err != nil {
		t.Errorf("CreateFile failed with error: %s", err)
	}
	file.Close()
	fileHandler.DeleteFile(fileName)

	_, err = os.Stat(fileName)
	if !os.IsNotExist(err) {
		t.Errorf("DeleteFile did not delete the file as expected")
	}
}

func TestRenameFile(t *testing.T) {
	newFileName := "newfile.txt"
	err := ioutil.WriteFile(fileName, []byte("Test file content"), 0644)
	if err != nil {
		t.Fatalf("Failed to write test file: %s", err)
	}
	defer os.Remove(newFileName)

	err = fileHandler.RenameFile(fileName, newFileName)
	if err != nil {
		t.Errorf("RenameFile failed with error: %s", err)
	}

	_, err = os.Stat(fileName)
	if !os.IsNotExist(err) {
		t.Errorf("RenameFile did not rename the file as expected")
	}

	_, err = os.Stat(newFileName)
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
