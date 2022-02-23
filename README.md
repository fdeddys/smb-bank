# smb-bank
## Setting Database
- install docker
- pull Golang postgresql
- User =  user_adm1n
- Password = _pa$sw0rd321.
- port = 5432
- Database = db_smb

## Register
    curl --location --request POST 'localhost:8888/api/1.0-1/account' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "username" :  "{{$randomUserName}}",
        "accountname" :  "{{$randomFullName}}"
    }'

## Login
    curl --location --request GET 'localhost:8888/api/1.0-1/account' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0b2tlbkNyZWF0ZWQiOiIxNjQ1NjMzNzE4NjYyIiwidXNlciI6Ikplc3VzNDciLCJ1c2VySWQiOiIxMiJ9.pFNDQsGWNDL1s328mBUu3-gRZv-VLkDDGmcnwCjTBB8'

## Get balance
    curl --location --request POST 'localhost:8888/api/1.0-1/account/login' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "userName":"deddy",
        "password":"0C91A43F8E1EC5FCBA28F8A5A34532679305CA131302AD2A06218B47F30CED88"
    }'