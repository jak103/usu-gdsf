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
