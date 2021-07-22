FROM golang:1.12.0-alpine3.9
LABEL maintainer="Josiah A"
RUN mkdir /app
ADD . /app
WORKDIR /app
COPY . .
#RUN go mod download
CMD ["go","run","main.go","-random"]