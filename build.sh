version=$1
registry=registry.cn-hangzhou.aliyuncs.com/flytrap
echo "build: $registry/go-template:$version"
docker build -t $registry/go-template:$version .

docker push $registry/go-template:$version
