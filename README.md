# lambda-common
Simple set of go modules that can be used with aws lambda development

## To publish a new version

Run tests
>go test

Run build
>go build

Update dependencies
>go get -u

Tidy dependencies
>go mod tidy

Tag branch
>git tag v1.x.x

Push Tag
>git push origin v1.x.x

Update go indices
>GOPROXY=proxy.golang.org go list -m github.com/gksoftware/lambda-common@v1.x.x

Create new release in GitHub