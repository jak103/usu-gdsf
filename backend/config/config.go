package config

import (
	"os"
	"strconv"
	"strings"
	
	"github.com/jak103/usu-gdsf/log"
)

// Environment
var DbType string
var RunEnv string
var MongoUri string
var FirestoreProjectId string

// Authentication
var TokenHashingKey string
var AccessTokenLifetimeMins int
var RefeshTokenLifetimeDays int

func init() {
	log.Info("Running configuration init")
	
	DbType = strings.ToLower(getEnvVarString("DB_TYPE"))
	RunEnv = getEnvVarString("RUN_ENV")

	if DbType == "mongo" {
		MongoUri = getEnvVarString("MONGO_URI")
	} else if DbType == "firestore" {
		FirestoreProjectId = getEnvVarString("FIRESTORE_PROJECT_ID")
	}
	
	TokenHashingKey = getEnvVarString("TOKEN_HASHING_KEY")
	AccessTokenLifetimeMins = getEnvVarInt("ACCESS_TOKEN_LIFETIME_MINS")
	RefeshTokenLifetimeDays = getEnvVarInt("REFRESH_TOKEN_LIFETIME_DAYS")
}

func getEnvVarString(key string) string {
	envVar := os.Getenv(key)

	if envVar == "" {
		log.Error("Environment variable ", key, " not set!")
		os.Exit(1)
	}

	return envVar
}

func getEnvVarInt(key string) int {
	envVarStr := getEnvVarString(key)
	envVarInt, err := strconv.Atoi(envVarStr)

	if err != nil {
		log.Error("Environment variable ", key, " must be an integer")
		os.Exit(1)
	}

	return envVarInt
}
