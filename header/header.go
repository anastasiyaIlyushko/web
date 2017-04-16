package header // import "gopkg.in/webnice/web.v1/header"

const (
	// Accept header
	Accept = `Accept`

	// AcceptEncoding header
	AcceptEncoding = `Accept-Encoding`

	// AcceptLanguage Accept-Language header
	AcceptLanguage = `Accept-Language`

	// Authorization header
	Authorization = `Authorization`

	// ContentDisposition header
	ContentDisposition = `Content-Disposition`

	// ContentEncoding header
	ContentEncoding = `Content-Encoding`

	// ContentLength header
	ContentLength = `Content-Length`

	// ContentType header
	ContentType = `Content-Type`

	// IfModifiedSince header
	IfModifiedSince = `If-Modified-Since`

	// LastModified header
	LastModified = `Last-Modified`

	// Location header
	Location = `Location`

	// Upgrade header
	Upgrade = `Upgrade`

	// Vary header
	Vary = `Vary`

	// WWWAuthenticate header
	WWWAuthenticate = `WWW-Authenticate`

	// XForwardedFor header
	XForwardedFor = `X-Forwarded-For`

	// XForwardedProto header
	XForwardedProto = `X-Forwarded-Proto`

	// XRealIP header
	XRealIP = `X-Real-IP`

	// Origin header
	Origin = `Origin`

	// AccessControlRequestMethod header
	AccessControlRequestMethod = `Access-Control-Request-Method`

	// AccessControlRequestHeaders header
	AccessControlRequestHeaders = `Access-Control-Request-Headers`

	// AccessControlAllowOrigin header
	AccessControlAllowOrigin = `Access-Control-Allow-Origin`

	// AccessControlAllowMethods header
	AccessControlAllowMethods = `Access-Control-Allow-Methods`

	// AccessControlAllowHeaders header
	AccessControlAllowHeaders = `Access-Control-Allow-Headers`

	// AccessControlAllowCredentials header
	AccessControlAllowCredentials = `Access-Control-Allow-Credentials`

	// AccessControlExposeHeaders header
	AccessControlExposeHeaders = `Access-Control-Expose-Headers`

	// AccessControlMaxAge header
	AccessControlMaxAge = `Access-Control-Max-Age`

	// MsEchoRequest Ms-Echo-Request header
	MsEchoRequest = `Ms-Echo-Request`

	// RetryAfter Retry-After header
	RetryAfter = `Retry-After`

	// Referer header
	Referer = `Referer`

	// UserAgent User-Agent header
	UserAgent = `User-Agent`

	// Pragma header
	Pragma = `Pragma`

	// CacheControl Cache-Control header
	CacheControl = `Cache-Control`

	// Expires header
	Expires = `Expires`

	// Кастомные заголовки

	// XUserAgent X-User-Agent header. Аналог User-Agent
	XUserAgent = `X-User-Agent`

	// XConnectionType X-Connection-Type header. Тип подключения, значения: ["3g", "wifi", "cable", "broadband", ...]
	XConnectionType = `X-Connection-Type`

	// XInternetServiceProvider X-Internet-Service-Provider header. Расширенная сегментация типа подключения по ISP (интернет-провайдер)
	XInternetServiceProvider = `X-Internet-Service-Provider`

	// XCarrier X-Carrier header. Расширенный тип мобильного трафика, значения: ["tele2", "mts", ...]
	XCarrier = `X-Carrier`

	// XTimeZone X-Time-Zone header. Значения: [0, 3, -1, 12, ...]
	XTimeZone = `X-Time-Zone`

	// XDayOfWeek X-Day-Of-Week header. День недели, значения: [0,1,2,3,4,5,6] 0=пнд, -1=не известен
	XDayOfWeek = `X-Day-Of-Week`

	// XGeoRegion X-Geo-Region header. Географический регион
	XGeoRegion = `X-Geo-Region`

	// XGeoCountry X-Geo-Country header. Страна
	XGeoCountry = `X-Geo-Country`

	// XGeoCity X-Geo-City header. Город
	XGeoCity = `X-Geo-City`

	// XDeviceManufacturer X-Device-Manufacturer header. Производитель/Марка устройства
	XDeviceManufacturer = `X-Device-Manufacturer`

	// XDeviceModel X-Device-Model header. Модель устройства
	XDeviceModel = `X-Device-Model`

	// XScreenSize X-Screen-Size header. Разрешение экрана
	XScreenSize = `X-Screen-Size`

	// XProxyType X-ProxyType header. Тип прокси, значения: ["vpn", "tor", "dch", "pub", "web"]
	XProxyType = `X-ProxyType`

	// XRedefinitionDeny X-Redefinition-Deny header. Запрет переопределения значений маркеров, значения: ["all", "query", "path", "entrypoint"]
	XRedefinitionDeny = `X-Redefinition-Deny`
)
