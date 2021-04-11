
### AWS SAM Package Init
`$ sam init --name NAME --runtime GO1.xx --app-template TEMPLATE`
`$ sam build && sam local invoke`
`$ sam build && sam package --s3-bucket S3BUCKET`
`$ sam deploy`

### aws sam deploy routine
To build and package for deployment
`$ sam build && sam package --s3-bucket <bucketname>`

// https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-getting-started-hello-world.html
deploy:
`$ sam deploy --guided`

```
$sam deploy --template-file .aws-sam/build/template.yaml --s3-bucket lvthillo-sam-upload-bucket --parameter-overrides ParameterKey=Environment,ParameterValue=aws ParameterKey=DDBTableName,ParameterValue=documentTable --stack-name aws-lambda-sam-demo --capabilities CAPABILITY_NAMED_IAM
```

### environment vars
* firebase (env file)
* aws sam (env file)
* private key for endpoints (pem files)
* dynamodb table url (template.yml)

### test endpoints
* POST /signin
* POST /signup
* GET /check_session
* GET /profile
* GET /citylist
* POST /user/edit

### data storage
* import `data.json` file

### localhost
* start dynamodb local
* local db file `./shared-local-instance.db`

```
java -Djava.library.path=./DynamoDBLocal_lib -jar ~/Downloads/dynamodb_local_latest/DynamoDBLocal.jar -sharedDb -dbPath .
```

* start dashboard for dynamodb-admin
```
export DYNAMO_ENDPOINT=http://localhost:8000

dynamodb-admin -p 8001

make local
```

