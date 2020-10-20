FROM golang:1.14-alpine AS build

WORKDIR /src/
COPY main.go main_test.go /src/
RUN CGO_ENABLED=0 go build -o /bin/greeting

FROM scratch
COPY --from=build /bin/greeting /bin/greeting
ENTRYPOINT ["/bin/greeting"]