if [ -z "$1" ]; then
    echo "No argument provided. Exiting."
    exit 1
fi
rm -rf ./bin
mkdir ./bin
env GOOS=windows GOARCH=amd64 go build -o  bin/GoDDNSClient_amd64_$1.exe
env GOOS=linux GOARCH=amd64 go build -o  bin/GoDDNSClient_amd64_$1
env GOOS=linux GOARCH=arm64 go build -o  bin/GoDDNSClient_arm64_$1