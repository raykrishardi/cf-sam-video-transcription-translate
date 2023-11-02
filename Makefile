.PHONY: build

build:
	cfn-include --yaml cloudFormation/templateSkeleton.yaml > template.yaml
	cd golambda && go mod download
	rm -rf .aws-sam
	sam build -t template.yaml

deploy:
	sam deploy --no-confirm-changeset