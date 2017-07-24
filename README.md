# go-guestbook-microservice-example
A very minimal microservice example, featuring a go guestbook backend, a website and Docker and Kubernetes files.

## guestbook-backend
https://travis-ci.org/llb4ll/go-guestbook-microservice-example.svg?branch=master

### Build
`docker build -t guestbook-backend .`

### Run
`docker run -p 8080:8080 -d --rm --name my-guestbook-backend guestbook-backend`


## guestbook-frontend
### Build
`docker build -t guestbook-frontend .`

### Run
`docker run -p 8088:80 -d --link my-guestbook-backend:guestbook-backend --rm --name my-guestbook-frontend guestbook-frontend`
