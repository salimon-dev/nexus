package db

import (
	"salimon/nexus/types"

	"gorm.io/gorm"
)

func GetUserPublicObject(user *types.User) types.PublicUser {
	return types.PublicUser{
		Id:           user.Id.String(),
		Username:     user.Username,
		Email:        user.Email,
		Credit:       user.Credit,
		Usage:        user.Usage,
		Role:         types.UserRoleToString(user.Role),
		Status:       types.UserStatusToString(user.Status),
		RegisteredAt: user.RegisteredAt,
		UpdatedAt:    user.UpdatedAt,
	}
}

func UsersModel() *gorm.DB {
	return DB.Model(types.User{})
}

func FindUser(query interface{}, args ...interface{}) (*types.User, error) {
	var user types.User

	result := DB.Model(types.User{}).Where(query, args...).Find(&user)

	if result.RowsAffected == 0 {
		return nil, result.Error
	} else {
		return &user, nil
	}
}

func InsertUser(user *types.User) error {
	result := DB.Model(types.User{}).Create(user)
	return result.Error
}
func UpdateUser(user *types.User) error {
	result := DB.Model(types.User{}).Where("id = ?", user.Id).Updates(user)
	return result.Error
}
