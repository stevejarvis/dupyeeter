ACCT_ID = 613475857488
GOARCH = amd64 
GOOS = linux

build_container:
	GOOS=$(GOOS) GOARCH=$(GOARCH) docker build -t dupyeeter . --platform=linux/amd64

docker_login:
	aws --profile dupyeeter-deploy ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin $(ACCT_ID).dkr.ecr.us-east-1.amazonaws.com

# TODO delete old images
deploy_ecr:
	docker tag  dupyeeter:latest $(ACCT_ID).dkr.ecr.us-east-1.amazonaws.com/dupyeeter:latest
	docker push $(ACCT_ID).dkr.ecr.us-east-1.amazonaws.com/dupyeeter:latest
	docker push $(ACCT_ID).dkr.ecr.us-east-1.amazonaws.com/dupyeeter:latest

update_lambda:
	aws --profile dupyeeter-deploy lambda update-function-code --region us-east-1 --function-name dupyeeter \
   --image-uri $(ACCT_ID).dkr.ecr.us-east-1.amazonaws.com/dupyeeter:latest   \

infra:
	cd terraform/; terraform apply; cd -

deploy: build_container deploy_ecr update_lambda

test:
	go test ./...
