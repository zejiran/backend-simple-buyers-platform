FROM go:1.7

COPY . .

RUN go install

CMD [ "go", "run"]