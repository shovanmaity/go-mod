FROM golang:1.14 AS builder
LABEL type=build-container
ARG GITHUB_USERNAME
ARG GITHUB_TOKEN
ARG GITLAB_USERNAME
ARG GITLAB_TOKEN
ARG BITBUCKET_USERNAME
ARG BITBUCKET_TOKEN
RUN git config --global \
  url."https://${GITHUB_USERNAME}:${GITHUB_TOKEN}@github.com".insteadOf \
  "https://github.com"
RUN git config --global \
  url."https://${GITLAB_USERNAME}:${GITLAB_TOKEN}@gitlab.com".insteadOf \
  "https://gitlab.com"
RUN git config --global \
  url."https://${BITBUCKET_USERNAME}:${BITBUCKET_TOKEN}@bitbucket.org".insteadOf \
  "https://bitbucket.org"
WORKDIR /app
COPY . .
RUN make vendor
RUN go run main.go
