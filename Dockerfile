FROM scratch

COPY gopath/bin/greeting /greeting

ENTRYPOINT ["/greeting"]