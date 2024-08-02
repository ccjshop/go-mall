package repo

import "gorm.io/gorm"

type RegisterFuncType func(db *gorm.DB) error

var (
	initFieldFuncTypes []RegisterFuncType
)

// 注册全字段更新初始化函数回调
func registerInitField(funcType RegisterFuncType) {
	initFieldFuncTypes = append(initFieldFuncTypes, funcType)
}

// InitField 全字段更新，初始化那些字段不更新，那些字段需要更新
func InitField(db *gorm.DB) error {
	for _, funcType := range initFieldFuncTypes {
		if err := funcType(db); err != nil {
			return err
		}
	}
	return nil
}
