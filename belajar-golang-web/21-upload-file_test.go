package belajar_golang_web

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func UploadForm(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "upload.form.gohtml", nil)
}

func Upload(writer http.ResponseWriter, request *http.Request) {
	// request.ParseMultipartForm(32 << 20) => ini yg dilakukan sama request.FormFile jika dibuka lebih detail

	// ambil file nya dari input data nya
	file, fileHeader, err := request.FormFile("file") // => params "file" itu harus sama dengan name yang ada di template nya, e.g. <label>File : <input type="file" name="file"></label><br>
	if err != nil {
		panic(err)
	}

	// membuat destinasi file nya
	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}

	// simpan semua file nya ke destinasi file nya
	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}

	// ambil inputan form yang bukan file
	name := request.PostFormValue("name")

	myTemplates.ExecuteTemplate(writer, "upload.success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

func TestUploadForm(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/2.png
var uploadFileTest []byte

func TestUploadFile(t *testing.T) {
	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)
	writer.WriteField("name", "Testing upload file")
	file, _ := writer.CreateFormFile("file", "CONTOHUPLOAD.png")
	file.Write(uploadFileTest)
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	Upload(recorder, request)

	bodyResponse, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyResponse))
}
