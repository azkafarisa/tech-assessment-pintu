# Technical Assessment Pintu (SRE)

## General info
This project is a simple monorepo which contain 2 backend apps
- Go
- NodeJS

## CICD Flow
- sonarqube (code analysis)
- unit-testing (run unit testing)
- build-app (create app artifact)
- push-to-ecr (build image and push to ecr)
- deploy-eks (deploy to eks cluster)

## How to Access
This simple app can be accessed in:
- Go: https://current-vast-lemur.ngrok-free.app/go/
- NodeJS: https://current-vast-lemur.ngrok-free.app/nodejs/
