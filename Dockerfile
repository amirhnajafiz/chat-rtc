# using alpine image
FROM golang:1.19-alpine

# setting labels
LABEL app=chat-rtc
LABEL dockerfile_version=0.1.1

# setting the port
ARG PORT=5000

# copy files
COPY . .

# build the executable file
RUN go build -o main

# delete dockerfile
RUN rm Dockerfile \
    && rm -rf assets \
    && rm README.md \
    && rm .gitignore

# expose the port
EXPOSE $PORT

# start server
CMD ./main --port $PORT
