# linux
export DB_HOST='localhost'
export DB_USERNAME='newuser'
export DB_NAME='testdb'
export DB_PASSWORD='123456'
export DB_PORT=4040
#  Database MAX IDLE_CONN  : Your server count with CPU Core 
export DB_MAX_IDLE_CONN=8
#  Database MAX OPEN_CONN  : Your server count with Logic CPU Core 
export DB_MAX_OPEN_CONN=16

export JWT_RSA_KEY_LOCATION='ArVm0yy4Qvi67CqfV6kf_mb-Tic' # You should use private_key.pem
export JWT_OLD_RSA_KEY_LOCATION=''

# 180 minute  ; 3 hours  
export JWT_TOKEN_LIFETIME=180
#PagedList 50 
export DEFAULT_RECORD_PER_PAGE=50
