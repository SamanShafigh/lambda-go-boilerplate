# Lambda Go Boilerplate
This repo provides some utility an a boilerplate to implement a Lambda function with Go and aws-lambda-go

## Build and deploy

As Go is a compiled language, build the application and create a Lambda deployment package. To do this, build a binary that runs on Linux, and zip it up into a deployment package.

```bash
$ GOOS=linux go build -o main
$ zip deployment.zip main conf.json
```

## Env variavle
You should define an env varibale for your labmda function and call it `dbCredentials` and it should be KMS Encrypt of the follwoing json data

```
{"dbUser": "", "dbPassword": "", "dbHost": "", "dbName": ""}
```

You can use `KmsEncrypt` function in `util` package to do this
```
util.KmsEncrypt(`{"dbUser": "", "dbPassword": "", "dbHost": "", "dbName": ""}`)
```

## Resources
Some useful resources

### Lambda & Go
Announcing Go Support for AWS Lambda
- https://aws.amazon.com/blogs/compute/announcing-go-support-for-aws-lambda/

The AWS SDK for Go provides APIs and utilities that developers can use to build Go applications that use AWS services
- https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/welcome.html

AWS SDK for Go
- https://aws.amazon.com/sdk-for-go/

Libraries, samples and tools to help Go developers develop AWS Lambda functions.
- https://github.com/aws/aws-lambda-go

### Go
Go by Example
- https://gobyexample.com/

Golang MySQL Tutorial
- https://tutorialedge.net/golang/golang-mysql-tutorial/

A MySQL-Driver for Go's database/sql package
- https://github.com/go-sql-driver/mysql