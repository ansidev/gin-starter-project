ARG APP_NAME=gin-starter-project

FROM registry.gitlab.com/ansidev/docker:golang-builder-latest AS builder

FROM registry.gitlab.com/ansidev/docker:golang-prod-latest as production

ENTRYPOINT ["/app/gin-starter-project"]
