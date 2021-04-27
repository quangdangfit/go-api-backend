package errors

// Prefix EC = ErrorCode

type (
	// ErrorCode : error code
	ErrorCode int32
)

// common errors
const (
	// ECOK is returned on success.
	// Error code starts from 100 in order not to duplicate grpc error codes which starts from 0 to 15.
	ECOK ErrorCode = iota + 100
	// ECCanceled indicates the operation was canceled (typically by the caller).
	ECCanceled
	// ECUnknown error. An example of where this error may be returned is
	// if a Status value received from another address space belongs to
	// an error-space that is not known in this address space. Also
	// errors raised by APIs that do not return enough error information
	// may be converted to this error.
	ECUnknown
	// ECDeadlineExceeded means operation expired before completion.
	// For operations that change the state of the system, this error may be
	// returned even if the operation has completed successfully. For
	// example, a successful response from a server could have been delayed
	// long enough for the deadline to expire.
	ECDeadlineExceeded
)

// user errors
const (
	// ECInvalidEmail : invalid email
	ECInvalidEmail ErrorCode = iota + 200
	// ECInvalidPassword : invalid password
	ECInvalidPassword
	// ECEmailNotExists : email doesn't exist
	ECEmailNotExists
	// ECEmailAlreadyExists : email already exists
	ECEmailAlreadyExists
	// ECPasswordMismatch : password and confirm password must match
	ECPasswordMismatch
	// ECInvalidCredentials : invalid credentials
	ECInvalidCredentials
	// ECInvalidName : invalid name
	ECInvalidName
	// ECUpdateUser : cannot update user profile
	ECUpdateUser
	// ECChangePasswordSame : new password and old password are the same
	ECChangePasswordSame
	// ECChangePasswordOldPwdNotSame : old password doesn't match
	ECChangePasswordOldPwdNotSame
	// ECForgotPasswordCode : invalid forgot password code
	ECForgotPasswordCode
	// ECPOSServerNotFound : POS server not found
	ECPOSServerNotFound
	// ECPOSTokenInvalid : POS's token is invalid
	ECPOSTokenInvalid
	// ECPOSTokenParamsNotFound : POS's token params not found
	ECPOSTokenParamsNotFound
	// ECInvalidVerificationCode : invalid verification code
	ECInvalidVerificationCode
	// ECIdentityDocumentAlreadyExists : identity document already exists
	ECIdentityDocumentAlreadyExists
	// ECIdentityDocumentNotFound : identity document not found
	ECIdentityDocumentNotFound
	// ECIdentityDocumentRejected24h : identity document rejected
	ECIdentityDocumentRejected24h
	// ECUserVerificationAlreadyExists : user verification already exists
	ECUserVerificationAlreadyExists
	// ECUserVerificationNotFound : user verification not found
	ECUserVerificationNotFound
	// ECUserVerificationStatusInvalid : user verification status invalid
	ECUserVerificationStatusInvalid
	// ECInvalidPhone : invalid phone
	ECInvalidPhone
	// ECPhoneAlreadyExists : phone already exists
	ECPhoneAlreadyExists
	// ECNotAllowUpdateEmail : not allow update email
	ECNotAllowUpdateEmail
	// ECVersionKeyInvalid : version key invalid
	ECVersionKeyInvalid
)

// server errors
const (
	// ECInvalidMessage : invalid message
	ECInvalidMessage ErrorCode = iota + 500
	// ECInvalidArgument : invalid argument
	ECInvalidArgument
	// ECInternalServerError : internal server error
	ECInternalServerError
	// ECInvalidLimit : invalid limit
	ECInvalidLimit
	// ECInvalidPage : invalid page
	ECInvalidPage
	// ECSystemError : system error
	ECSystemError
	// ECPermissionDenied : permission denied
	ECPermissionDenied

	// ECInvalidAdminKey : admin key is invalid
	ECInvalidAdminKey

	// ECInvalidJSON : JSON is invalid
	ECInvalidJSON
	// ECFileNotFound : file is not found
	ECFileNotFound
	// ECInvalidMIME : invalid MIME type
	ECInvalidMIME
	// ECUploadFileFail : upload file is fail
	ECUploadFileFail
	// ECMarshal : marshal fail
	ECMarshal
	// ECUnmarshal : unmarshal fail
	ECUnmarshal

	// ECInvalidDateTime : DateTime is invalid
	ECInvalidDateTime
	// ECHashPasswordFail : hash password fail
	ECHashPasswordFail
)

// third - party errors
const (
	// ECMongoConnection : mongodb connection error
	ECMongoConnection ErrorCode = iota + 1000
	// ECMongoCreate : mongodb create model error
	ECMongoCreate
	// ECMongoRead : mongodb read model error
	ECMongoRead
	// ECMongoUpdate : mongodb update model error
	ECMongoUpdate
	// ECMongoDelete : mongodb delete model error
	ECMongoDelete

	// ECMySQLConnection : mysql connection error
	ECMySQLConnection
	// ECMySQLDBEmpty : mysql db is empty
	ECMySQLDBEmpty
	// ECMySQLDBAutoMigrate : mysql db auto migrate error
	ECMySQLDBAutoMigrate
	// ECMySQLCreate : mysql create model error
	ECMySQLCreate
	// ECMySQLRead : mysql read model error
	ECMySQLRead
	// ECMySQLUpdate : mysql update model error
	ECMySQLUpdate
	// ECMySQLDelete : mysql delete model error
	ECMySQLDelete
	// ECMySQLAfterSave : mysql after_save hook model error
	ECMySQLAfterSave

	// ECSolrConnection : solr connection error
	ECSolrConnection
	// ECSolrRead : solr read error
	ECSolrRead
	// ECSolrCreate : solr create error
	ECSolrCreate
	// ECSolrUpdate : solr update error
	ECSolrUpdate
	// ECSolrDelete : solr delete error
	ECSolrDelete
	// ECSolrCommit : solr commit error
	ECSolrCommit

	// ECRedisConnection : redis connection error
	ECRedisConnection
	// ECRedisGet : redis get error
	ECRedisGet
	// ECRedisSet : redis set error
	ECRedisSet
	// ECRedisRemove : redis remove error
	ECRedisRemove
)
