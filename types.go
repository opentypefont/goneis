package goneis

import (
	"errors"
	"github.com/valyala/fastjson"
	"net/http"
)

const baseURL = "https://open.neis.go.kr/hub"

var (
	ErrMissingValues   = errors.New("required values are missing")
	ErrInvalidKey      = errors.New("key is invalid")
	ErrServiceNotFound = errors.New("cannot find requested service")
	ErrInvalidLVType   = errors.New("location value type is invalid")
	ErrTooManyData     = errors.New("too many data request(no more than 1000)")
	ErrTooManyRequest  = errors.New("too many request")
	ErrServerError     = errors.New("server error occurred")
	ErrDBConnection    = errors.New("database connection error occurred")
	ErrSQLQuery        = errors.New("sql query error occurred")
	ErrKeyRestricted   = errors.New("key is restricted by administrator")
	ErrMissingData     = errors.New("cannot find data")
)

type HTTPClient struct {
	key    string
	http   *http.Client
	parser *fastjson.Parser
}

type NeisClient struct {
	http *HTTPClient
}

type Parameters map[string]string

type Error struct {
	Code    string
	Message string
}

type School struct {
	OfficeCode       string
	OfficeName       string
	Code             string
	Name             string
	EnglishName      string
	Type             string
	ZipCode          string
	Address          string
	DetailedAddress  string
	TelNumber        string
	Homepage         string
	Gender           string
	FaxNumber        string
	HighSchoolType   string
	SPHighSchoolType string
	FoundationDate   string
	FoundationDay    string
}
