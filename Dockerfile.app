FROM golang:1.21

WORKDIR /koksmat
COPY . .
WORKDIR /koksmat/.koksmat/app
RUN go install

CMD [ "sleep","infinity"]