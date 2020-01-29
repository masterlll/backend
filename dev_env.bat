:: windowns server   
set DB_HOST=localhost
set DB_USERNAME=newuser
set DB_NAME=testdb
set DB_PASSWORD=123456
set DB_PORT=4040
:: Database MAX IDLE_CONN  : Your server count with CPU Core 
set DB_MAX_IDLE_CONN=8
::  Database MAX OPEN_CONN  : Your server count with Logic CPU Core 
set DB_MAX_OPEN_CONN=16
:: You should use private_key.pem
set JWT_RSA_KEY_LOCATION=private_key.pem
set JWT_OLD_RSA_KEY_LOCATION=''

 :: 180 minute  ; 3 hours  
set JWT_TOKEN_LIFETIME=180
 ::PagedList 50 
set DEFAULT_RECORD_PER_PAGE=50
