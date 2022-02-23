# smb-bank

curl --location --request GET 'localhost:8888/api/1.0-1/account'


curl --location --request POST 'localhost:8888/api/1.0-1/account' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username" : "deddy2"   
}'


curl --location --request POST 'localhost:8888/api/1.0-1/account/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username":"deddy",
    "password":"0C91A43F8E1EC5FCBA28F8A5A34532679305CA131302AD2A06218B47F30CED88"
}'