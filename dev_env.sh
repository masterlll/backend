export DB_HOST='localhost'
export DB_USERNAME='test_user'
export DB_NAME='test_db'
export DB_PASSWORD='user_password'
export DB_PORT=700
#  Database MAX IDLE_CONN  : Your server count with CPU Core 
export DB_MAX_IDLE_CONN=8
#  Database MAX OPEN_CONN  : Your server count with Logic CPU Core 
export DB_MAX_OPEN_CONN=16

export JWT_RSA_KEY_LOCATION='/your/private_key/addr/private_key.pem'
export JWT_OLD_RSA_KEY_LOCATION='/yourdata/old_key.pem'

# 180 minute  ; 3 hours  
export JWT_TOKEN_LIFETIME=180
#PagedList 50 
export DEFAULT_RECORD_PER_PAGE=50
