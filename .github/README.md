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

* Perform one of the following options:
    * Clone this repository and build the Docker image
    ```bash
    git clone https://github.com/codemicro/githubCommitInfo.git
    cd githubCommitInfo
    docker build -t commitinfo .
    ```
    * Or, provided it's not been removed by the image retention policy, pull a prebuilt image from the Docker Hub.
    ```bash
    docker pull codemicro/github-commit-info
    ```
    and use the image name `codemicro/github-commit-info` in the next step.
* Start the Docker container
  ```bash
  docker run -d --restart unless-stopped -p 8000:80 -e GITHUB_OAUTH_TOKEN="your personal access token" -e GITHUB_USERNAME="codemicro" commitinfo
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

### Reverse proxying

To reverse proxy this using the Apache2 HTTP server:

```
<Location /a/path>
    ProxyPass http://127.0.0.1:8000/
    ProxyPassReverse http://127.0.0.1:8000/
</Location>
```

### Shields.io

This service is designed to work as a JSON endpoint that can be used with Shields.io. Read more [here](https://shields.io/endpoint).