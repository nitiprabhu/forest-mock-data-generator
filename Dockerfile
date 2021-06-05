FROM golang:1.12-alpine
RUN mkdir /app
COPY forest-mock-data-generator /app
WORKDIR /app

CMD ["./forest-mock-data-generator"]
