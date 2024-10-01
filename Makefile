.PHONY: build

build:
	sam build

build-and-run:
	rm -rf .aws-sam/build
	rm -rf .aws-sam/cache
	sam build
	sam local start-api

build-and-deploy:
	rm -rf .aws-sam/build
	rm -rf .aws-sam/cache
	sam build
	sam deploy