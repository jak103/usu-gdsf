[![Go](https://github.com/jak103/usu-gdsf/actions/workflows/go.yml/badge.svg)](https://github.com/jak103/usu-gdsf/actions/workflows/go.yml)

# usu-gdsf

DevOps is awesome!

To start the project with docker-compose, simply run `docker-compose up -d`

Once everything is up and running, simply try hitting the `/game` endpoint at `localhost:8080`


# Front end testing info

* We are using the jest testing framework with the vue/test-utils to mount a component
* To create a test just have the file extention be `<filename>.test.spec`
* Use the command `npm test` to run the tests.
* To test elements of a component the data-test attribute will need to be added to find the element eg. `data-test="listItem1`
* 

## Backend Server Configuration

Server settings are configured using environment variables. The following environment variables are *required* to be defined for the server to function properly:

### Environment

* `DB_TYPE`

  Specifies which database to use. Valid values are `firestore` and `mongo`
  
* `RUN_ENV`

  Specifies which database to use. Valid values are `firestore`, `mongo`, and `mock`

* `MONGO_URI` (only if `DB_TYPE` is `mongo`)

  The URI used to connect to Mongo DB
  
* `FIRESTORE_PROJECT_ID` (only if `DB_TYPE` is `firestore`)

  The ID of the Firestore project

### Authentication

* `TOKEN_HASHING_KEY`

  The secret key used for generating and HMAC for authentication tokens
      
* `ACCESS_TOKEN_LIFETIME_MINS`

  Specifies the length of time, in minutes, that generated access tokens will be valid before they expire.
  
* `REFRESH_TOKEN_LIFETIME_DAYS`

  Specifies the length of time, in days, that generated refresh tokens will be valid before they expire.
