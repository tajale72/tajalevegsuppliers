package imageprocessing

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func SendImageToPythonService(imagePath string) {
	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", file.Name())
	if err != nil {
		fmt.Println("Error creating form file:", err)
		return
	}
	io.Copy(part, file)
	writer.Close()

	req, err := http.NewRequest("POST", "http://localhost:5000/process-image", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Process the response
	responseBody, _ := io.ReadAll(resp.Body)
	fmt.Println("Response from Python Service:", string(responseBody))
}
