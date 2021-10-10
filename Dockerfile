FROM golang:1.16-alpine

WORKDIR /home/estudiante/snap/go/src/github.com/backend-simple-buyers-platform

COPY . .

#RUN go install

EXPOSE 3000

CMD ["go", "run"]