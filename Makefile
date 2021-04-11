clean:
	rm -rf bin/
build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/main ./*go
	# env GOOS=linux go build -ldflags="-s -w" -o bin/jwt_test jwt.go jwt_test.go
	sam build
local:
	# https://stackoverflow.com/questions/48926260/connecting-aws-sam-local-with-dynamodb-in-docker
	# https://www.slideshare.net/AmazonWebServices/local-testing-and-deployment-best-practices-for-serverless-applications-aws-online-tech-talks
	sam local start-api --static-dir data.json