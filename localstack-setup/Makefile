# Variables
LAMBDA_FUNCTION=lambda_function.py
ZIP_FILE=lambda_function_payload.zip
AWS_REGION=us-east-1
ROLE_NAME=lambda_exec_role
FUNCTION_NAME=localstack_lambda
LOCALSTACK_ENDPOINT=http://localhost:4566

# Default target
all: clean package deploy

# Package the Lambda function
package:
	@echo "Packaging the Lambda function..."
	zip -r $(ZIP_FILE) $(LAMBDA_FUNCTION)

# Deploy the Lambda function using AWS CLI (assumes Localstack is running)
deploy:
	@echo "Deploying the Lambda function to Localstack..."
	awslocal lambda create-function \
		--function-name $(FUNCTION_NAME) \
		--runtime python3.8 \
		--role arn:aws:iam::000000000000:role/$(ROLE_NAME) \
		--handler lambda_function.handler \
		--zip-file fileb://$(ZIP_FILE) || \
		awslocal lambda update-function-code \
		--function-name $(FUNCTION_NAME) \
		--zip-file fileb://$(ZIP_FILE)

# Invoke the Lambda function to test
invoke:
	@echo "Invoking the Lambda function..."
	awslocal lambda invoke \
		--function-name $(FUNCTION_NAME) \
		response.json
	@cat response.json

# Clean up generated files
clean:
	@echo "Cleaning up..."
	rm -f $(ZIP_FILE) response.json

# Delete the Lambda function
delete:
	@echo "Deleting the Lambda function from Localstack..."
	awslocal lambda delete-function \
		--function-name $(FUNCTION_NAME)

.PHONY: all package deploy invoke clean delete
