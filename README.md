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
`docker run -d nickscord-bot`