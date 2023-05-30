TEST_FLAGS := -v

# Targets
.PHONY: all clean test node-deps deploy-dev build dev generate gen rds-local

# Default target
all: build

# Clean target for cleaning up generated files
clean:
	@go clean -testcache -x

# Test target for running the test suite
test: clean
	@pnpm sst bind 'go test $(shell go list ./... | grep -v integratedfinance_legacy) $(TEST_FLAGS)'

test-enfuce: clean
	@pnpm sst bind 'go test github.com/science-card/science-card-api/internal/enfuce/api/... $(TEST_FLAGS)'

test-sumsub: clean
	@pnpm sst bind 'go test github.com/science-card/science-card-api/internal/sumsub/api/... $(TEST_FLAGS)'

node-deps:
	@pnpm i --frozen-lockfile

# Deploy target for deploying the application to AWS
deploy-dev: node-deps
	@pnpm run deploy --stage dev
local-dev: node-deps
	@pnpm run dev

# Build target for building the application stack with AWS CDK
build: node-deps
	@pnpm build

# Dev target for developing the functions locally
dev:
	@pnpm dev

generate, gen:
	@go generate ./...

PORT ?= 5433
rds-local:
ifeq ($(STAGE),)
	$(error You must define the STAGE variable)
endif

	@aws ssm start-session \
		--target $(shell aws ec2 describe-instances --filters "Name=tag:Name,Values=${STAGE}-sc-vpc-bastion" | jq -r '.Reservations[].Instances[] | select(.State.Name != "terminated") | .InstanceId') \
		--document-name AWS-StartPortForwardingSessionToRemoteHost \
		--parameters host="$(shell aws rds describe-db-clusters --db-cluster-identifier ${STAGE}-sc-rds --query "DBClusters[0].Endpoint" --output text)",portNumber="5432",localPortNumber="$(PORT)"

apply-schema:
ifeq ($(STAGE),)
	$(error You must define the STAGE variable)
endif

	@atlas schema apply \
		--env default \
		--var secret_name=$(shell aws secretsmanager list-secrets --query "SecretList[?Tags[?Key == 'aws:cloudformation:stack-name' && Value == '${STAGE}-sc-rds']] | [0].Name")
