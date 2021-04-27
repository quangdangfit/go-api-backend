package errors

// Prefix EM = ErrorMessage

// common errors
const (
	// EMOK : successful
	EMOK string = "no error"
	// EMCanceled : operation was canceled
	EMCanceled string = "operation was canceled"
	// EMUnknown : Unknown error
	EMUnknown string = "unknown error"
	// EMDeadlineExceeded : operation expired
	EMDeadlineExceeded string = "operation expired"
)

// user errors
const (
	// EMInvalidEmail : error message invalid email
	EMInvalidEmail string = "invalid email"
	// EMInvalidPassword : error message invalid password
	EMInvalidPassword string = "invalid password"
	// EMEmailNotExists : error message email doesn't exist
	EMEmailNotExists string = "email doesn't exist"
	// EMEmailAlreadyExists : error message email already exists
	EMEmailAlreadyExists string = "email already exists"
	// EMPasswordMismatch : perror message assword and confirm password must match
	EMPasswordMismatch string = "password and confirm password must match"
	// EMInvalidCredentials : error message invalid credentials
	EMInvalidCredentials string = "invalid credentials"
	// EMInvalidName : error message invalid name
	EMInvalidName string = "invalid name"
	// EMUpdateUser : error message cannot update user profile
	EMUpdateUser string = "cannot update user profile"
	// EMChangePasswordSame : new password and old password are the same
	EMChangePasswordSame string = "new password and old password are the same"
	// EMChangePasswordOldPwdNotSame : old password doesn't match
	EMChangePasswordOldPwdNotSame string = "old password doesn't match"
	// EMForgotPasswordCode : invalid forgot password code
	EMForgotPasswordCode string = "invalid forgot password code"
	// EMPOSServerNotFound : POS server not found
	EMPOSServerNotFound string = "POS server not found"
	// EMPOSTokenInvalid : POS's token is invalid
	EMPOSTokenInvalid string = "POS's token is invalid"
	// ECPOSTokenParamsNotFound : POS's token params not found
	EMPOSTokenParamsNotFound string = "POS's token params not found"
	// EMInvalidVerificationCode : invalid verification code
	EMInvalidVerificationCode string = "invalid verification code"
	// EMIdentityDocumentAlreadyExists : identity document already exists
	EMIdentityDocumentAlreadyExists string = "identity document already exists"
	// EMIdentityDocumentNotFound : identity document not found
	EMIdentityDocumentNotFound string = "identity document not found"
	// ECIdentityDocumentRejected24h : identity document rejected
	EMIdentityDocumentRejected24h string = "identity document rejected, please retry after 24h"
	// EMUserVerificationAlreadyExists : user verification already exists
	EMUserVerificationAlreadyExists string = "user verification already exists"
	// ECUserVerificationNotFound : user verification not found
	EMUserVerificationNotFound string = "user verification not found"
	// ECUserVerificationStatusInvalid : user verification status invalid
	EMUserVerificationStatusInvalid string = "user verification status invalid"
	// EMInvalidPhone : invalid phone
	EMInvalidPhone string = "invalid phone"
	// EMPhoneAlreadyExists : phone already exists
	EMPhoneAlreadyExists string = "phone already exists"
	// EMNotAllowUpdateEmail : not allow update email
	EMNotAllowUpdateEmail string = "not allow update email"
	// EMVersionKeyInvalid : version key invalid
	EMVersionKeyInvalid string = "version key invalid"
)

// server errors
const (
	// EMInvalidMessage : invalid message
	EMInvalidMessage string = "invalid message"
	// EMInvalidArgument : invalid argument
	EMInvalidArgument string = "invalid argument"
	// EMInternalServerError : internal server error
	EMInternalServerError string = "internal server error"
	// EMInvalidLimit : invalid limit
	EMInvalidLimit string = "invalid limit"
	// EMInvalidPage : invalid page
	EMInvalidPage string = "invalid page"
	// EMSystemError : system error
	EMSystemError string = "system error"
	// EMPermissionDenied : permission denied
	EMPermissionDenied string = "permission denied"

	// EMInvalidAdminKey : admin key is invalid
	EMInvalidAdminKey string = "admin key is invalid"

	// ECInvalidJSON : JSON is invalid
	EMInvalidJSON string = "JSON is invalid"
	// EMInvalidJSON : file is not found
	EMFileNotFound string = "file is not found"
	// EMInvalidJSON : invalid MIME type
	EMInvalidMIME string = "invalid MIME type"
	// EMUploadFileFail : upload file is fail
	EMUploadFileFail string = "upload file is fail"
	// EMMarshal : marshal fail
	EMMarshal string = "upload file is fail"
	// EMUnmarshal : unmarshal fail
	EMUnmarshal string = "unmarshal fail"

	// EMInvalidDateTime : DateTime is invalid
	EMInvalidDateTime string = "dateTime is invalid"
	// ECHashPasswordFail : hash password fail
	EMHashPasswordFail string = "hash password fail"
)

// third - party errors
const (
	// EMMongoConnection : mongodb connection error
	EMMongoConnection string = "mongodb connection error"
	// EMMongoCreate : mongodb create model error
	EMMongoCreate string = "mongodb create model error"
	// EMMongoRead : mongodb read model error
	EMMongoRead string = "mongodb read model error"
	// EMMongoUpdate : mongodb update model error
	EMMongoUpdate string = "mongodb update model error"
	// EMMongoDelete : mongodb delete model error
	EMMongoDelete string = "mongodb delete model error"

	// EMMySQLConnection : mysql connection error
	EMMySQLConnection string = "mysql connection error"
	// EMMySQLDBEmpty : mysql db is empty
	EMMySQLDBEmpty string = "mysql db is empty error"
	// EMMySQLDBAutoMigrate : mysql db auto migrate error
	EMMySQLDBAutoMigrate string = "mysql db auto migrate error"
	// EMMySQLCreate : mysql create model error
	EMMySQLCreate string = "mysql create model error"
	// EMMySQLRead : mysql read model error
	EMMySQLRead string = "mysql read model error"
	// EMMySQLUpdate : mysql update model error
	EMMySQLUpdate string = "mysql update model error"
	// EMMySQLDelete : mysql delete model error
	EMMySQLDelete string = "mysql delete model error"
	// EMMySQLAfterSave : mysql after save hook model error
	EMMySQLAfterSave string = "mysql after save hook model error"

	// EMSolrConnection : solr connection error
	EMSolrConnection string = "solr connection error"
	// EMSolrRead : solr read error
	EMSolrRead string = "solr read error"
	// EMSolrCreate : solr create error
	EMSolrCreate string = "solr create error"
	// EMSolrUpdate : solr update error
	EMSolrUpdate string = "solr update error"
	// EMSolrDelete : solr delete error
	EMSolrDelete string = "solr delete error"
	// EMSolrCommit : solr commit error
	EMSolrCommit string = "solr commit error"

	// EMRedisConnection : redis connection error
	EMRedisConnection string = "redis connection error"
	// EMRedisGet : redis get error
	EMRedisGet string = "redis get error"
	// EMRedisSet : redis set error
	EMRedisSet string = "redis set error"
	// EMRedisRemove : redis remove error
	EMRedisRemove string = "redis remove error"
)

var errorMessages map[ErrorCode]string

func init() {
	errorMessages = make(map[ErrorCode]string, 0)

	initDefaultCommonErrors()
	initDefaultUserErrors()
	initDefaultServerErrors()
	initDefaultThirdPartyErrors()
}

func initDefaultCommonErrors() {
	errorMessages[ECOK] = EMOK
	errorMessages[ECCanceled] = EMCanceled
	errorMessages[ECUnknown] = EMUnknown
	errorMessages[ECDeadlineExceeded] = EMDeadlineExceeded
}

func initDefaultUserErrors() {
	errorMessages[ECInvalidEmail] = EMInvalidEmail
	errorMessages[ECInvalidPassword] = EMInvalidPassword
	errorMessages[ECEmailNotExists] = EMEmailNotExists
	errorMessages[ECEmailAlreadyExists] = EMEmailAlreadyExists
	errorMessages[ECPasswordMismatch] = EMPasswordMismatch
	errorMessages[ECInvalidCredentials] = EMInvalidCredentials
	errorMessages[ECInvalidName] = EMInvalidName
	errorMessages[ECUpdateUser] = EMUpdateUser
	errorMessages[ECChangePasswordSame] = EMChangePasswordSame
	errorMessages[ECChangePasswordOldPwdNotSame] = EMChangePasswordOldPwdNotSame
	errorMessages[ECForgotPasswordCode] = EMForgotPasswordCode
	errorMessages[ECPOSServerNotFound] = EMPOSServerNotFound
	errorMessages[ECPOSTokenInvalid] = EMPOSTokenInvalid
	errorMessages[ECPOSTokenParamsNotFound] = EMPOSTokenParamsNotFound
	errorMessages[ECInvalidVerificationCode] = EMInvalidVerificationCode
	errorMessages[ECIdentityDocumentAlreadyExists] = EMIdentityDocumentAlreadyExists
	errorMessages[ECIdentityDocumentNotFound] = EMIdentityDocumentNotFound
	errorMessages[ECIdentityDocumentRejected24h] = EMIdentityDocumentRejected24h
	errorMessages[ECUserVerificationAlreadyExists] = EMUserVerificationAlreadyExists
	errorMessages[ECUserVerificationNotFound] = EMUserVerificationNotFound
	errorMessages[ECUserVerificationStatusInvalid] = EMUserVerificationStatusInvalid
	errorMessages[ECInvalidPhone] = EMInvalidPhone
	errorMessages[ECPhoneAlreadyExists] = EMPhoneAlreadyExists
	errorMessages[ECNotAllowUpdateEmail] = EMNotAllowUpdateEmail
	errorMessages[ECVersionKeyInvalid] = EMVersionKeyInvalid
}

func initDefaultServerErrors() {
	errorMessages[ECInvalidMessage] = EMInvalidMessage
	errorMessages[ECInvalidArgument] = EMInvalidArgument
	errorMessages[ECInternalServerError] = EMInternalServerError
	errorMessages[ECInvalidLimit] = EMInvalidLimit
	errorMessages[ECInvalidPage] = EMInvalidPage
	errorMessages[ECSystemError] = EMSystemError
	errorMessages[ECPermissionDenied] = EMPermissionDenied

	errorMessages[ECInvalidAdminKey] = EMInvalidAdminKey
	errorMessages[ECInvalidJSON] = EMInvalidJSON
	errorMessages[ECFileNotFound] = EMFileNotFound
	errorMessages[ECInvalidMIME] = EMInvalidMIME
	errorMessages[ECUploadFileFail] = EMUploadFileFail
	errorMessages[ECMarshal] = EMMarshal
	errorMessages[ECUnmarshal] = EMUnmarshal

	errorMessages[ECInvalidDateTime] = EMInvalidDateTime
	errorMessages[ECHashPasswordFail] = EMHashPasswordFail
}

func initDefaultThirdPartyErrors() {
	errorMessages[ECMongoConnection] = EMMongoConnection
	errorMessages[ECMongoCreate] = EMMongoCreate
	errorMessages[ECMongoRead] = EMMongoRead
	errorMessages[ECMongoUpdate] = EMMongoUpdate
	errorMessages[ECMongoDelete] = EMMongoDelete

	errorMessages[ECMySQLConnection] = EMMySQLConnection
	errorMessages[ECMySQLDBEmpty] = EMMySQLDBEmpty
	errorMessages[ECMySQLDBAutoMigrate] = EMMySQLDBAutoMigrate
	errorMessages[ECMySQLCreate] = EMMySQLCreate
	errorMessages[ECMySQLRead] = EMMySQLRead
	errorMessages[ECMySQLUpdate] = EMMySQLUpdate
	errorMessages[ECMySQLDelete] = EMMySQLDelete
	errorMessages[ECMySQLAfterSave] = EMMySQLAfterSave

	errorMessages[ECSolrConnection] = EMSolrConnection
	errorMessages[ECSolrRead] = EMSolrRead
	errorMessages[ECSolrCreate] = EMSolrCreate
	errorMessages[ECSolrUpdate] = EMSolrUpdate
	errorMessages[ECSolrDelete] = EMSolrDelete
	errorMessages[ECSolrCommit] = EMSolrCommit

	errorMessages[ECRedisConnection] = EMRedisConnection
	errorMessages[ECRedisGet] = EMRedisGet
	errorMessages[ECRedisSet] = EMRedisSet
	errorMessages[ECRedisRemove] = EMRedisRemove

}
