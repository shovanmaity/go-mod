.PHONY: vendor
vendor:
	@GO111MODULE=on \
		GOPRIVATE=bitbucket.org/shovanmaity/go-pvt-module,github.com/shovanmaity/go-pvt-module,gitlab.com/shovanmaity/go-pvt-module \
		GIT_TERMINAL_PROMPT=1 \
		go mod tidy
	@GO111MODULE=on \
		GOPRIVATE=bitbucket.org/shovanmaity/go-pvt-module,github.com/shovanmaity/go-pvt-module,gitlab.com/shovanmaity/go-pvt-module \
		GIT_TERMINAL_PROMPT=1 \
		go mod download
	@GO111MODULE=on \
		GOPRIVATE=bitbucket.org/shovanmaity/go-pvt-module,github.com/shovanmaity/go-pvt-module,gitlab.com/shovanmaity/go-pvt-module \
		GIT_TERMINAL_PROMPT=1 \
		go mod vendor

.PHONY: image
.DEFAULT_GOAL := image
image:
	docker build -t test-private-go-mod:ci \
		--build-arg GITHUB_USERNAME=$(GITHUB_USERNAME) \
		--build-arg GITHUB_TOKEN=$(GITHUB_TOKEN) \
		--build-arg GITLAB_USERNAME=$(GITLAB_USERNAME) \
		--build-arg GITLAB_TOKEN=$(GITLAB_TOKEN) \
		--build-arg BITBUCKET_USERNAME=$(BITBUCKET_USERNAME) \
		--build-arg BITBUCKET_TOKEN=$(BITBUCKET_TOKEN) \
		-f Dockerfile .
	docker image prune -f --filter label=type=build-container