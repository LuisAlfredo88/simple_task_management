FROM debian:bullseye

ENV GO111MODULE=on
ENV PORT=8080
WORKDIR /app

COPY ./build/main /app
COPY ./cmd/bin/main.db /app

RUN chmod u+x ./main

RUN ls

CMD ["./main"]