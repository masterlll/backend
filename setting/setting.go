//a centralized location to store the configration from the environment variable

package setting

//a centralized location containing all the configuration needed by the system.
//so that it would avoid collusion of the config parameter name
const (

	// DATABASE
	DBHost        string = `DB_HOST`
	DBUserName    string = `DB_USERNAME`
	DBName        string = `DB_NAME`
	DBPassword    string = `DB_PASSWORD`
	DBPort        string = `DB_PORT`
	DBMaxIdleConn string = `DB_MAX_IDLE_CONN`
	DBMaxOpenConn string = `DB_MAX_OPEN_CONN`
	// JWT
	JwtRsaKeyLocation    string = `JWT_RSA_KEY_LOCATION`
	JwtOldRsaKeyLocation string = `JWT_OLD_RSA_KEY_LOCATION`
	JwtToekenLifeTime    string = `JWT_TOKEN_LIFETIME`
	// PageList
	DefaultRecordPerPage string = `DEFAULT_RECORD_PER_PAGE`
)
