# [WIP] gin-starter-project

# Description

This project aims to introduce a project structure for a Go project.

Articles:
- [SOLID Principle](https://www.digitalocean.com/community/conceptual_articles/s-o-l-i-d-the-first-five-principles-of-object-oriented-design).
- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).

# Features

- [x] Gin.
- [x] Gorm.
- [x] GitLab CI/CD.
- [x] Conventional commit.
- [x] Docker.

# Project structure

```
|- app: Main app directory
|  |- config: load env variables into config struct
|  |- constant: app constants
|  |- domain
|     |- <domain_name>: Domain package
|        |- <domain_model>.go: Domain model(s).
|        |- <domain_model>_repository.go: Domain repository interface.
|  |- migration: migration files
|  |- pkg: shared packages
|  |- app.env.example: Template file for `app.env`.
|  |- app.go: Main app
|  |- wire.go: Wire DI file.
|  |- wire_gen.go: Wire DI generated file.
|  |- <domain_name>: Domain package (app-specific)
|     |- <protocol>: http/amqp/...
|        |- <domain_name>_controller.go: Controller
|     |- repository:
|        |- <domain_name>_repository.go: Implementation for repository
|     |- service:
|        |- <domain_name>_service.go: Interfaces and implementations for service
|- .husky.yaml: Husky configs.
|- Dockerfile.app: App's Docker image.
|- Dockerfile.builder: Builder's Docker image.
|- Dockerfile.prod: Production Docker image.
```

**Notes**:

- For large-scale projects, you can split service interface into 2 parts: domain service & app service.
- This starter project is using:
  - Gin (controller layer)
  - Gorm (repository layer)
  
  If you want to use another framework/library, feel free to fork this project and customize them.
- Because this is a starter project, so the logic of repository, service and controller only cover simple cases.

# Instructions

## Coding rules

 - Use singular form for names. Ex: `model`, `author`, `post`.
 - Interface name for Repository and Service starts with `I`. Ex: `IPostRepository`, `IPostService`.
 - Receiver var for repository is `r`.
 - Receiver var for service is `s`.
 - Receiver var for controller is `ctrl`.

## Checklist

After cloning this project, please do the following things:

- [ ] Run `cd ./app && make prepare` to install prerequisite tools.
- [ ] Update email regex (`.husky.yaml`) if necessary.
- [ ] Review Makefile for supported commands.
- [ ] Update environment variables (`app.env`).
- [ ] Run `make start` to start Docker container(s).
- [ ] Run `make migup` to update database structure.
- [ ] Connect to database and insert data for tables `author` and `post`.
- [ ] Run `make rundev` to run the API server.
- [ ] You can test 2 routes: `http://localhost:8080/post/v1/posts/1`, `http://localhost:8080/author/v1/authors/1` (The routes can be changed depends on your environment).

## Build

This build instructions is tested on GitLab CI/CD. Other CI/CD platforms might not be guaranteed.

### Prerequisites

1. Update variables
- [ ] Change value of GOPRIVATE (Dockerfile.builder)
- [ ] Change id of OS's user. Default value: `ansidev`, (Dockerfile.builder, Dockerfile.prod)
- [ ] Change Docker base image names (Dockerfile.app)

### Build base Docker images

- `cd` to this folder and run commands:

```shell
docker build -t registry.gitlab.com/ansidev/docker:golang-builder-latest -f Dockerfile.builder .
docker build -t registry.gitlab.com/ansidev/docker:golang-prod-latest -f Dockerfile.prod .
```

**Notes**: If you change Docker base images in Dockerfile.app, please adjust above commands.

- Run `docker images` and check if you have the docker image with repository names: `registry.gitlab.com/ansidev/docker:golang-builder`, `registry.gitlab.com/ansidev/docker:golang-prod`.

### Build app's Docker Images

#### Prepare Information

- Login to your GitLab account.
- Go to https://gitlab.com/-/profile/account and copy your GitLab username.
- Go to https://gitlab.com/-/profile/personal_access_tokens.
- Create new personal access token with scopes: `read_api`.
- Copy the generated personal access token, note that you need to store it by yourself because GitLab will not display it anymore. If you forgot the personal access token, you must revoke the old personal access token and create a new one.

### Build Docker image

1. Run
```
export DOCKER_NETRC="machine gitlab.com login <gitlab_username> password <gitlab_personal_access_token>"
```

Note:
- `gitlab_username`: Get from step 2.
- `gitlab_personal_access_token`: Get from step 5.

Example command:
```
export DOCKER_NETRC="machine gitlab.com login gitlab_user password glpat-6ADHo3yM2hBTBZ5U3kGk"
```

2. Build Docker image

```
docker build . --build-arg DOCKER_NETRC -t registry.gitlab.com/ansidev/gin-starter-project -f Dockerfile.app .
```
