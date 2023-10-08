package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestIntegration(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "integration-test")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	producerDir := filepath.Join(tempDir, "producer")
	consumerDir := filepath.Join(tempDir, "consumer")
	if err := copyDir("c:/Users/HP/Documents/Assignment/producer", producerDir); err != nil {
		t.Fatalf("Failed to copy producer directory: %v", err)
	}
	if err := copyDir("c:/Users/HP/Documents/Assignment/consumer", consumerDir); err != nil {
		t.Fatalf("Failed to copy consumer directory: %v", err)
	}

	producerServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, world!"))
	}))
	defer producerServer.Close()

	configFile := filepath.Join(consumerDir, "config.json")
	configData := []byte(`{"producer_url": "` + producerServer.URL + `"}`)
	if err := writeFile(configFile, configData); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	consumerServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get(producerServer.URL)
		if err != nil {
			t.Fatalf("Failed to make request to producer server: %v", err)
		}
		defer resp.Body.Close()

		_, err = io.Copy(w, resp.Body)
		if err != nil {
			t.Fatalf("Failed to copy response body to consumer: %v", err)
		}
	}))
	defer consumerServer.Close()

	resp, err := http.Get(consumerServer.URL)
	if err != nil {
		t.Fatalf("Failed to make request to consumer server: %v", err)
	}
	defer resp.Body.Close()

	var body bytes.Buffer
	_, err = io.Copy(&body, resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body from consumer server: %v", err)
	}

	expectedBody := []byte("Hello, world!")
	if !bytes.Equal(body.Bytes(), expectedBody) {
		t.Errorf("Response body is incorrect: expected %q, got %q", expectedBody, body.Bytes())
	}
}

func copyDir(src, dst string) error {
	if err := os.MkdirAll(dst, 0755); err != nil {
		return err
	}

	files, err := os.ReadDir(src)
	if err != nil {
		return err
	}
	for _, file := range files {
		srcPath := filepath.Join(src, file.Name())
		dstPath := filepath.Join(dst, file.Name())
		if file.IsDir() {
			if err := copyDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			if err := copyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}

	return nil
}

func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}

func writeFile(filename string, data []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	return err
}
