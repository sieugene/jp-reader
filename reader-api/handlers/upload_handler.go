package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type UploadResponse struct {
	Message string `json:"message"`
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(10 << 20)

	paramTitle := r.FormValue("title")
	if paramTitle == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	files := r.MultipartForm.File["file"]
	if len(files) == 0 {
		http.Error(w, "No files uploaded", http.StatusBadRequest)
		return
	}

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	for _, handler := range files {
		file, err := handler.Open()
		if err != nil {
			http.Error(w, "Error retrieving file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		part, err := writer.CreateFormFile("file", handler.Filename)
		if err != nil {
			fmt.Println("Error creating form file:", err)
			return
		}

		if _, err = io.Copy(part, file); err != nil {
			fmt.Println("Error copying file:", err)
			return
		}

	}

	if err := writer.Close(); err != nil {
		fmt.Println("Error closing writer:", err)
		return
	}

	requestEndpoint := os.Getenv("MOKURO_SERVICE") + "/upload/" + paramTitle
	req, err := http.NewRequest("POST", requestEndpoint, &buf)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Make request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending files:", err)
		return
	}
	defer resp.Body.Close()

	var uploadResp UploadResponse
	if err := json.NewDecoder(resp.Body).Decode(&uploadResp); err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	w.Write([]byte(uploadResp.Message))
}
