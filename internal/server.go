package internal

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type FileInfo struct {
	Name  string
	IsDir bool
	Size  string
	Path  string
}

func RunServer(dirPath string, port string, recursive bool) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fullPath := filepath.Join(dirPath, filepath.Clean(r.URL.Path))
		if !recursive && r.URL.Path != "/" {
			rel, _ := filepath.Rel(dirPath, fullPath)
			if strings.Contains(rel, string(os.PathSeparator)) {
				http.Error(w, "Subfolder access disabled (-R not set)", http.StatusForbidden)
				return
			}
		}

		if !strings.HasPrefix(fullPath, dirPath) {
			http.Error(w, "Access Denied", http.StatusForbidden)
			return
		}

		if r.Method == http.MethodPost {
			handleUpload(w, r, fullPath)
			return
		}

		info, err := os.Stat(fullPath)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		if info.IsDir() {
			renderDirectory(w, fullPath, r.URL.Path)
		} else {
			http.ServeFile(w, r, fullPath)
		}
	})

	ips := getLocalIPs()

	fmt.Println("--------------------------------------------")
	fmt.Printf("🚀 File Server Started!\n")
	fmt.Printf("📂 Sharing: %s\n", dirPath)
	fmt.Printf("📂 Recursive: %v\n", recursive)
	fmt.Println("--------------------------------------------")
	fmt.Println("Access the server at:")

	fmt.Printf("  🏠 Local:   http://localhost:%s\n", port)
	for _, ip := range ips {
		fmt.Printf("  🌐 Network: http://%s:%s\n", ip, port)
	}
	fmt.Println("--------------------------------------------")
	fmt.Println("Press Ctrl+C to stop.")

	http.ListenAndServe(":"+port, nil)
}
