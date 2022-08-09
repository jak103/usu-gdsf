package db

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/jak103/usu-gdsf/config"
)

// uploadFile uploads an object (.exe to google cloud storage)
func UploadFile(filePath string, fileName string) error {
	storageBucket := "BUCKET_NAME"
	url := "https://storage.googleapis.com/upload/storage/v1/b/"+ storageBucket +"/o?uploadType=media&name=" + fileName
	var bearer = "Bearer " + config.GoogleCloudStorageToken
	req, err := http.NewRequest("POST", url, nil)
	req.Header.Set("Authorization", bearer)
	req.Header.Set("Content-Type", "exe")
	req.Header.Set("Object_Location", filePath)
	req.Header.Set("Connection", "keep-alive")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	log.Println(string([]byte(body)))
	return nil
}

func deleteFileFromCloudStorage(fileName string) error {
	storageBucket := "BUCKET_NAME"
	var bearer = "Bearer " + config.GoogleCloudStorageToken
	req, err := http.NewRequest("DELETE", "https:\\" + storageBucket + ".storage.googleapis.com?"+ fileName, nil)
	req.Header.Set("Host", storageBucket + ".storage.googleapis.com")
	req.Header.Set("Authorization", bearer)
	req.Header.Set("Date", time.Now().String())
	req.Header.Set("Content Length", "0")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	log.Println(string([]byte(body)))
	return nil
}

