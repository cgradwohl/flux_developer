package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	fmt.Println("beetle starting on :8080")

	configDir := "/etc/beetle-config"

	// List files in the config directory
	files, err := ioutil.ReadDir(configDir)
	if err != nil {
		log.Fatalf("failed to read config directory: %v", err)
	}

	for _, file := range files {
		if !file.IsDir() {
			path := filepath.Join(configDir, file.Name())
			content, err := ioutil.ReadFile(path)
			if err != nil {
				log.Printf("failed to read file %s: %v", path, err)
				continue
			}
			fmt.Printf("==== %s ====\n%s\n", file.Name(), string(content))
		}
	}

	http.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

