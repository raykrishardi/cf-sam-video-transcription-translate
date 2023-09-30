.PHONY: build

build:
	cfn-include --yaml cloudFormation/templateSkeleton.yaml > template.yaml
	cd golambda && go mod download
	sam build -t template.yaml
