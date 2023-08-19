### from: https://dev.to/aurelievache/learning-go-by-examples-part-3-create-a-cli-app-in-go-1h43

1. go mod init github.com/tsemach/go-gopher-cli
2. go get -u github.com/spf13/cobra@latest
3. go install github.com/spf13/cobra-cli@latest
  - cobra-cli installed on ~/go/bin if $GOPATH is not defined

4. init cobra by: `cobra-cli init`
5. get viper package: `go get github.com/spf13/viper@v1.8.1`
6. cobra-cli add get
7. cobra-cli add unit
8. cobra-cli add install -p unit

9. go install: this command create an executable in $GOPATH/bin
  export GOBIN=$GOPATH/bin
  export PATH=$PATH:$GOBIN

## IOT Cli Commands

### Environment Variable
- IOTCLI_ENVIRONMENT - if defined and --env is provide that use it for selecting the environment

1. iot unit create --env dev | --pid 12345678901 --tid &lt;tls-uuid&#62;<br>
   --pid and --tid are opetional

2. iot unit install --env dev | --pid 12345678901 --tid &lt;tls-uuid&#62; --data @install.json

3. iot unit assign --env dev | --pid 12345678901 --tid &lt;tls-uuid&#62; --data @install.json
