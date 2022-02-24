docker build -t app_be_image .

# docker run -d -p 5432:5432 --name db_smb -e POSTGRES_USER=user_adm1n -e POSTGRES_PASSWORD=_pa$sw0rd321. -e POSTGRES_DB=db_smb --network
# app_be_network postgres

# docker run -d -p 8888:8888 --name app_be -e DB_HOST=db_smb --network app_be_network app_be_image