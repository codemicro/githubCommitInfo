# commitInfo

A small microservice to pull the latest number of commits from the GitHub API for a specific user.

Designed to work as a custom endpoint for [Shields.io](https://shields.io).

### Usage
```bash
$ curl http://127.0.0.1:8000/
{"schemaVersion":1,"label":"Commits","message":"1578","colour":"green"}
```

### Prerequisites
* Docker or Go 1.14 or later
* A GitHub personal access token

### Setup (Docker)

* Clone this repository
  ```bash
  git clone https://github.com/codemicro/githubCommitInfo.git
  cd githubCommitInfo
  ```
* Build the Docker image
  ```bash
  docker build -t commitInfo .
  ```
* Start the Docker container
  ```bash
  docker run -d --restart unless-stopped -p 8000:80 -e GITHUB_OAUTH_TOKEN="your personal access token" -e GITHUB_USERNAME="codemicro" commitInfo
  ```
  This will start the server on port 8000 of your local machine for the user `codemicro`.

### Setup (compilation)

This should run on any platform that can run Go 1.14 or later.

* Clone this repository
  ```bash
  git clone https://github.com/codemicro/githubCommitInfo.git
  cd githubCommitInfo
  ```
* Build using Go
  ```bash
  go build github.com/codemicro/githubCommitInfo/cmd/githubCommitInfo
  sudo chmod +x githubCommitInfo
  ```
* Run
  ```bash
  GITHUB_OAUTH_TOKEN="your personal access token"
  GITHUB_USERNAME="codemicro"
  ./githubCommitInfo
  ```
  This will start the server on port 80 for the user `codemicro`.

### Shields.io

This service is designed to work as a JSON endpoint that can be used with Shields.io. Read more [here](https://shields.io/endpoint).