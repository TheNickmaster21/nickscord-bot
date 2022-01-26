# Nickscord Bot

This is a simple Discord bot.

## Setup

Copy the example.env file as .env and set your BOT_SECRET

## Test

`go test . -v`

## Run without Docker

`go run .`

## Run with Docker

Build the docker image:
`docker build -t nickscord-bot .`

Run the docker image (detached):
`docker run -d --env-file ./.env nickscord-bot`

## Run with Kubernetes

Build the docker image:
`docker build -t nickscord-bot .`

Tag the docker image:
`docker tag nickscord-bot registry.digitalocean.com/nmain/nickscord-bot`

Push the docker image:
`docker push registry.digitalocean.com/nmain/nickscord-bot`

Create the Kubernetes secret:
`kubectl create secret generic nickscord-bot --from-literal='bot_token=***'`

Apply the Kubernetes resources (service and deployment):
`kubectl apply -f kubernetes.yml`