package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWT 签名结构
type JWT struct {
	config JwtConfig
}

// 一些常量
var (
	TokenExpired     error = errors.New("授权已过期")
	TokenNotValidYet error = errors.New("令牌未激活")
	TokenMalformed   error = errors.New("令牌非法")
	TokenInvalid     error = errors.New("令牌解析异常")
)

// JwtConfig 配置
type JwtConfig struct {
	TimeOut time.Duration // 超时时间
	Issuer  string        // 签证签发人
	SignKey string
}

// Member 载荷
type Member struct {
	UserID uint64 // 用户id
}

// CustomClaims 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	UserInfo Member `json:"userInfo"`
	jwt.StandardClaims
}

// NewJWT 新建一个jwt实例
func NewJWT(config JwtConfig) *JWT {
	return &JWT{
		config: config,
	}
}

// getSignKey 获取signKey
func (j *JWT) getSignKey() []byte {
	return []byte(j.config.SignKey)
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.getSignKey())
}

// CreateUserToken 生成含有用户信息的token
func (j *JWT) CreateUserToken(u *Member) (string, error) {
	jwtConfig := j.config
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		UserInfo: *u,
		StandardClaims: jwt.StandardClaims{
			// 设置时效
			ExpiresAt: time.Now().Add(jwtConfig.TimeOut * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    jwtConfig.Issuer,
		},
	})
	return claims.SignedString(j.getSignKey())
}

// ParseToken 解析Token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.getSignKey(), nil
	})

	if err != nil {
		// jwt定义了一系列的常量，枚举在解析和验证一个JWT（JSON Web Token）时可能遇到的错误情况。每个常量都对应一个特定的验证错误类型。这些常量通常用于JWT库中，以便在验证过程中提供详细的错误信息。下面是每个常量的解释：
		// 1. `ValidationErrorMalformed`：表示JWT格式不正确。这可能是因为它不是一个有效的JSON格式，或者它的结构不符合JWT的标准结构（例如，缺少头部、载荷或签名部分）。
		// 2. `ValidationErrorUnverifiable`：表示JWT无法验证，因为签名问题。这可能是因为提供的密钥不正确，或者签名算法不被支持。
		// 3. `ValidationErrorSignatureInvalid`：表示签名验证失败。这意味着JWT的签名与头部和载荷不匹配，可能是因为数据被篡改或者签名密钥不正确。
		// 以下是与JWT标准声明（Standard Claims）相关的验证错误：
		// 4. `ValidationErrorAudience`：表示受众（AUD）验证失败。JWT的受众声明可能不匹配预期的受众值。
		// 5. `ValidationErrorExpired`：表示过期时间（EXP）验证失败。JWT的过期时间已经过去，因此令牌不再有效。
		// 6. `ValidationErrorIssuedAt`：表示签发时间（IAT）验证失败。如果签发时间是未来的时间点，那么令牌可能被认为是无效的。
		// 7. `ValidationErrorIssuer`：表示发行人（ISS）验证失败。JWT的发行人声明可能不匹配预期的发行人值。
		// 8. `ValidationErrorNotValidYet`：表示不在有效期内（NBF）验证失败。JWT的不在有效期内声明指定了一个时间，在这个时间之前令牌不应被接受。
		// 9. `ValidationErrorId`：表示JWT ID（JTI）验证失败。如果提供了JTI声明，它可能需要匹配特定的值或满足某些条件。
		// 10. `ValidationErrorClaimsInvalid`：表示一般性的声明验证错误。这是一个通用错误，可能是因为某些自定义声明没有通过验证或者令牌中的某些声明不符合预期的格式或值。
		// 这些常量通常与位运算结合使用，以便可以同时表示多个错误。例如，如果一个JWT同时有过期和签名无效的问题，那么相应的错误代码可能是`ValidationErrorExpired | ValidationErrorSignatureInvalid`。这样，验证函数可以返回一个包含所有相关错误的单一值，调用者可以通过位运算来检查特定的错误是否发生。
		var ve *jwt.ValidationError
		if errors.As(err, &ve) {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				// JWT格式不正确。这可能是因为它不是一个有效的JSON格式，或者它的结构不符合JWT的标准结构（例如，缺少头部、载荷或签名部分）。
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// 过期时间（EXP）验证失败。JWT的过期时间已经过去，因此令牌不再有效。
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				// 不在有效期内（NBF）验证失败。JWT的不在有效期内声明指定了一个时间，在这个时间之前令牌不应被接受。
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	// 解析成功
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	// 返回失败
	return nil, TokenInvalid
}

// RefreshToken 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.getSignKey(), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
