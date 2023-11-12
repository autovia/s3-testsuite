# s3-testsuite

Unofficial Amazon AWS S3 compatibility tests for S3 servers like [s3-go](https://github.com/autovia/s3-go)

## Run tests

Example aws config for local testing

cat $HOME/.aws/config

```shell
[default]
endpoint_url=http://localhost:3000
```

cat $HOME/.aws/config

```shell
[default]
aws_access_key_id = user
aws_secret_access_key = password
region = us-east-1
```

Run tests

```shell
go run main.go
```

## License

MIT