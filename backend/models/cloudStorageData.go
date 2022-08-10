package models

/**
 * TODO: remove this comment after some implementation
 * The idea of this object is to have a better way to track where data for each game is.
 * In terms of some env variables that we will need (this is after doing some testing with GCS)
 * We will need the following:
 * GOOGLE_APPLICATION_CREDENTIALS (path to json that contains credentials)
 * GOOGLE_CLOUD_PROJECT (the project id that we will be working with)
 * 
 */
type CloudStorageData struct {
	ObjectName	string
	ObjectData 	string
}

func(d *CloudStorageData) SetData(data string) {
	d.ObjectData = data
}