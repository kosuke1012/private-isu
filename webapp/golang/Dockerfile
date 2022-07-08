FROM golang:1.18

RUN mkdir -p /home/webapp/image
COPY . /home/webapp
WORKDIR /home/webapp
RUN go build -o app
CMD ./app
