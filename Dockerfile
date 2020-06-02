FROM golang:latest as dev

WORKDIR /app

COPY . /app/

RUN go install
RUN CGO_ENABLED=0 GOOS=linux go build -o api

EXPOSE 5000

CMD ["fresh"]


FROM golang:latest as prod

WORKDIR /app

COPY --from=dev /app/api .

CMD [ "./api" ]