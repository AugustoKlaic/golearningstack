## Golang testing

### Simple unit tests in golang
Example in file: [learning_mapper_test.go](https://github.com/AugustoKlaic/golearningstack/blob/master/test/mapper/learning_mapper_test.go)

### Unit testing that needs repository mocking
Example in file: [learning_service_test.go](https://github.com/AugustoKlaic/golearningstack/blob/master/test/service/learning_service_test.go)

### Unit testing controller http calls
Example in file: [learning_controller_test.go](https://github.com/AugustoKlaic/golearningstack/blob/master/test/controller/learning_controller_test.go)

### Lib used for generating mocks and doing asserts
- Mockgen
- setup
  - run: ``go get -u github.com/golang/mock/mockgen``
  - run: ``go install github.com/golang/mock/mockgen@latest``
  - Add this to the file that the mock will be generated (an interface probably):
    ``//go:generate mockgen -source={fileName}.go -destination={pathYouWant}/{fileNameYouWant}.go``
  - run: ``go generate -v ./... to generate mocks``


## Go comands to run tests

- Simple test run ``go test ./..``
  - This will show on terminal all tests that passed with and ``ok``
- To see coverage report ``go test ./... -cover``
  - This will show on terminal a report of coverage
- To generate a profile file of coverage ``go test ./... -coverprofile={pathYouWantHere}/coverage``
  - This will generate a coverage file in the designated folder that will be required for the next commands
  - To exclude packages:
    - Windows: ``go test @(go list ./... | Where-Object {$_ -notmatch "{packageNameX}|{packageNameY}"}) -coverprofile={pathYouWantHere}\coverage``
    - Linux: ``go test $(go list ./... | grep -v -E "{packageNameX}|{packageNameY}") -coverprofile={pathYouWantHere}/coverage``
- To see the coverage on terminal by function ``go tool cover -func={pathOfCoverageFile}/coverage`` 
- To see an HTML report of the coverage ``go tool cover -html={pathOfCoverageFile}/coverage``