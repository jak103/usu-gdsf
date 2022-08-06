package db

import (
	"io/ioutil"
	"log"
	"net/http"
	"github.com/jak103/usu-gdsf/config"
)

// uploadFile uploads an object.
func UploadFile(filePath string, fileName string) error {
	storageBucket := "BUCKET_NAME"
	url := "https://storage.googleapis.com/upload/storage/v1/b/"+ storageBucket +"/o?uploadType=media&name=" + fileName
	var bearer = "Bearer " + config.GoogleCloudStorageToken
	log.Println(bearer)
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

