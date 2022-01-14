GOOS=darwin \
GOARCH=amd64 \
CGO_ENABLED=1 \
CGO_CFLAGS="-DAPL=1 -DIBM=0 -DLIN=0" \
CGO_LDFLAGS="-F/System/Library/Frameworks/ -F$(pwd)/Libraries/Mac -framework XPLM" \
code .