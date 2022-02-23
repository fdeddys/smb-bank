# smb-bank

## register
curl --location --request POST 'localhost:8888/api/1.0-1/account' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username" :  "{{$randomUserName}}",
    "accountname" :  "{{$randomFullName}}"
}'

## login
curl --location --request GET 'localhost:8888/api/1.0-1/account' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0b2tlbkNyZWF0ZWQiOiIxNjQ1NjMzNzE4NjYyIiwidXNlciI6Ikplc3VzNDciLCJ1c2VySWQiOiIxMiJ9.pFNDQsGWNDL1s328mBUu3-gRZv-VLkDDGmcnwCjTBB8'

## get balance
curl --location --request POST 'localhost:8888/api/1.0-1/account/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "userName":"Jesus47",
    "password":"d03822d4efab331d962067da6fd81204e0ce6fab7886ba9786d383c818ec6ec2"
}'