package converter

import "github.com/Krab1o/meebin/internal/model"

func ConvertRoles(roles []model.Role) []model.Role {
	if roles == nil {
		return nil
	}
	copiedRoles := make([]model.Role, len(roles))
	copy(copiedRoles, roles)

	return copiedRoles
}
