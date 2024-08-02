package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

// PasswordEncoder 接口定义了密码编码器必须实现的方法。
type PasswordEncoder interface {
	Encode(rawPassword string) (string, error)        // 加密原始密码
	Matches(rawPassword, encodedPassword string) bool // 检查原始密码和加密密码是否匹配
	UpgradeEncoding(encodedPassword string) bool      // 检查加密密码是否需要升级
}

// BcryptPasswordEncoder 是 PasswordEncoder 接口的一个实现，使用 bcrypt 算法。
type BcryptPasswordEncoder struct{}

// Encode 使用 bcrypt 对密码进行加密。
// rawPassword 是用户提供的原始密码。
// 返回加密后的密码，或者在加密过程中出现错误时返回错误。
func (b *BcryptPasswordEncoder) Encode(rawPassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// Matches 比较原始密码和加密后的密码是否匹配。
// rawPassword 是用户提供的原始密码。
// encodedPassword 是存储的加密密码。
// 返回一个布尔值，表示密码是否匹配。
func (b *BcryptPasswordEncoder) Matches(rawPassword, encodedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encodedPassword), []byte(rawPassword))
	return err == nil
}

// UpgradeEncoding 检查加密密码是否需要升级。
// encodedPassword 是存储的加密密码。
// 返回一个布尔值，表示密码编码是否需要升级。
// 在这个简单的实现中，我们总是返回 false，表示不需要升级。
func (b *BcryptPasswordEncoder) UpgradeEncoding(encodedPassword string) bool {
	// 这里可以根据需要实现密码编码的升级逻辑。
	// 例如，可以检查密码散列的版本或算法强度。
	// 目前，我们假设所有密码都是最新的，不需要升级。
	return false
}
