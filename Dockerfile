FROM        golang
RUN         mkdir -p /app
WORKDIR     /app
COPY        . .
#RUN         go mod init CESapi
#RUN         go get -u github.com/gin-gonic/gin
RUN         go mod download
RUN         go build -o app
ENTRYPOINT  ["./app"]