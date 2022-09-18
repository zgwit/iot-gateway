
# 整体编译
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GOPRIVATE=*.gitlab.com,*.gitee.com
go env -w GOSUMDB=off

app="iot-master-gateway"
pkg="iot-master-gateway"
version="1.0.0"

read -t 5 -p "please input version(default:$version)" ver
if [ -n "${ver}" ];then
	version=$ver
fi


gitHash=$(git show -s --format=%H)
buildTime=$(date -d today +"%Y-%m-%d %H:%M:%S")

# -w -s
ldflags=" -w -s -X '$pkg/args.Version=$version' \
-X '$pkg/args.gitHash=$gitHash' \
-X '$pkg/args.buildTime=$buildTime'"

#export CGO_ENABLED=1
#CC=x86_64-w64-mingw32-gcc

export GOARCH=amd64

export GOOS=windows
name="iot-master-gateway-windows-amd64.exe"
go build -ldflags "$ldflags" -o iot-master-gateway-windows-amd64.exe main.go
tar -zvcf iot-master-gateway-windows-amd64.tar.gz iot-master-gateway-windows-amd64.exe
rm iot-master-gateway-windows-amd64.exe

export GOOS=linux
#CC=gcc
name="iot-master-gateway-linux-amd64"
go build -ldflags "$ldflags" -o iot-master-gateway-linux-amd64 main.go
tar -zvcf iot-master-gateway-linux-amd64.tar.gz iot-master-gateway-linux-amd64
rm iot-master-gateway-linux-amd64

#export CC=arm-linux-gnueabihf-gcc

export GOARCH=arm64
name="iot-master-gateway-linux-arm64"
go build -ldflags "$ldflags" -o iot-master-gateway-linux-arm64 main.go
tar -zvcf iot-master-gateway-linux-arm64.tar.gz iot-master-gateway-linux-arm64
rm iot-master-gateway-linux-arm64

export GOARCH=arm

export GOARM=6
name="iot-master-gateway-linux-armv6l"
go build -ldflags "$ldflags" -o iot-master-gateway-linux-armv6l main.go
tar -zvcf iot-master-gateway-linux-armv6l.tar.gz iot-master-gateway-linux-armv6l
rm iot-master-gateway-linux-armv6l

export GOARM=7
name="iot-master-gateway-linux-armv7l"
go build -ldflags "$ldflags" -o iot-master-gateway-linux-armv7l main.go
tar -zvcf iot-master-gateway-linux-armv7l.tar.gz iot-master-gateway-linux-armv7l
rm iot-master-gateway-linux-armv7l

