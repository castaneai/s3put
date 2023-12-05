# s3put

Simple CLI to upload files to AWS S3 and S3-compatible storage.

## Install


```
go install github.com/castaneai/s3put/cmd/s3put@latest
```

## Usage

```sh
# Show usage
s3put -h

# example usage
s3put -region us-east1 /path/to/file s3://bucket/key
```


Set credentials in the following environment variables.

```
export AWS_ACCESS_KEY_ID=AKIAIOSFODNN7EXAMPLE
export AWS_SECRET_ACCESS_KEY=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
```

## License

MIT

## Author

[castaneai](https://github.com/castaneai)
