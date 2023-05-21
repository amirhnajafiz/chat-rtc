# Chat RTC

In this project we are going to use ```fiber``` websockets and Golang
```channels``` to create a Real Time Server in order to make
a Real Time Communication system.

<br />

## What is RTC?

Real-Time Communication (RTC) is a 
development platform for delivering 
real-time audio and video content, or
any kind of media around the world.

## Connections

Each user connects to server by calling ```/ws``` endpoint
to create a websocket. After that each data is being processed 
by our rtc server.

<br />

<div align="center">
    <img src="./assets/schema.png" alt="schema" width="500" />
</div>

<br />

## Run

In order to start application you can use in terminal build
commands or Docker.

### In terminal

```shell
go build -o main
chmod +x ./main
./main --port 8080
```

### Docker

```shell
docker build . --build-args PORT=8080 -t chat-rtc:v0.1
docker run -d -it --port 8080:8080 char-rtc:v0.1
```

In your browser open ```localhost:8080```.
