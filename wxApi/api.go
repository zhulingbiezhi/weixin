package wxApi

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func uploadFile(url, filePath, fileType string) ([]byte, error) {
	// Prepare a form that you will submit to that URL.
	var b bytes.Buffer
	multiWriter := multipart.NewWriter(&b)
	// Add your image file
	f, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("open file error: %s, the file path: %s", err.Error(), filePath))
	}
	defer f.Close()
	fw, err := multiWriter.CreateFormFile(fileType, filepath.Base(filePath))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("multiWriter CreateFormFile error: ", err.Error()))
	}
	if _, err = io.Copy(fw, f); err != nil {
		return nil, errors.New(fmt.Sprintf("Copy file to multiWriter error: %s", err.Error()))
	}
	multiWriter.Close()

	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("make new post request error: %s", err.Error()))
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", multiWriter.FormDataContentType())
	req.Header.Set("Content-Disposition", "attachment; filename=\"ququ.jpg\"")

	// Submit the request
	client := http.DefaultClient
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("response error: %s", err.Error()))
	} else {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("ioutil read response error: %s", err.Error()))
		} else {
			return data, nil
		}
	}
}

func DownloadFile(url, filePath string) ([]byte, error) {
	// Create the file
	fmt.Println(url)
	out, err := os.Create(filePath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("create file error: %s", err.Error()))
	}
	defer out.Close()

	// Get the data
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("make new post request error: %s", err.Error()))
	}
	req.Header.Add("Accept", "*/*")
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("http client do request error: %s", err.Error()))
	}
	defer resp.Body.Close()
	fmt.Print(resp.Header)
	// Writer the body to file
	mediaType := resp.Header.Get("Content-Type")

	if strings.Contains(mediaType, "image") {

	} else if strings.Contains(mediaType, "video") {

	} else /*if strings.Contains(mediaType, "json") */ {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("ioutil read response error: %s", err.Error()))
		} else {
			return data, nil
		}
	}
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("io copy response error: %s", err.Error()))
	}

	return nil, nil
}
