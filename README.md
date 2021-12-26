## Upload file to S3 buket

Upload file `myfile.txt` in S3 bucket `my-bucket/myfiles/` with timeout of 10 minutes

```
export AWS_REGION=region-name
export AWS_ACCESS_KEY_ID=aws-access-key-id
export AWS_SECRET_ACCESS_KEY=aws-secret-key
go run upload2s3.go -b "mybucket/myfiles/" -k "myfile.txt" -d 10m < myfile.txt
```

Key must have s3:PutObject permission.

## List instances

List instances with specific tag

```
export AWS_REGION=region-name
export AWS_ACCESS_KEY_ID=aws-access-key-id
export AWS_SECRET_ACCESS_KEY=aws-secret-key
go run listec2instance.go --tag "Name" --value "myserver"
```

## Start or stop an instances

Start an instance

```
export AWS_REGION=region-name
export AWS_ACCESS_KEY_ID=aws-access-key-id
export AWS_SECRET_ACCESS_KEY=aws-secret-key
go run start-stopec2.go --id "i-ab54a98efc35e659" --action "start"
```

Stop an instance

```
export AWS_REGION=region-name
export AWS_ACCESS_KEY_ID=aws-access-key-id
export AWS_SECRET_ACCESS_KEY=aws-secret-key
go run start-stopec2.go --id "i-ab54a98efc35e659" --action "stop"
```
