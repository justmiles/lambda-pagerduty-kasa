default: build deploy clean invoke
	
deploy:
	aws lambda update-function-code \
		--region us-east-1 \
		--function-name PagerDutyLight \
		--zip-file fileb://./deployment.zip

init:
	aws lambda create-function \
	 --region us-east-1 \
	 --function-name PagerDutyLight \
	 --zip-file fileb://./deployment.zip \
	 --runtime go1.x \
	 --role arn:aws:iam::965579072529:role/service-role/lambda \
	 --handler main

build:
	GOOS=linux go build -o main main.go types.go
	zip deployment.zip main

test:
	go test

invoke: invoke-resolve

invoke-resolve:
	curl -d '{ "messages": [{ "event": "incident.resolve" }] }' -X POST https://28k5tmaitc.execute-api.us-east-1.amazonaws.com/prd/powerswitch

invoke-acknowledge:
	curl -d '{ "messages": [{ "event": "incident.acknowledge" }] }' -X POST https://28k5tmaitc.execute-api.us-east-1.amazonaws.com/prd/powerswitch

invoke-trigger:
	curl -d '{ "messages": [{ "event": "incident.trigger" }] }' -X POST https://28k5tmaitc.execute-api.us-east-1.amazonaws.com/prd/powerswitch

clean:
	rm -rf main deployment.zip
