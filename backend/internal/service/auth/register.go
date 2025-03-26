package auth

import (
	"context"
	"time"

	converter "github.com/Krab1o/meebin/internal/converter/service/user"
	"github.com/Krab1o/meebin/internal/model"
	rmodel "github.com/Krab1o/meebin/internal/model/user/r_model"
	smodel "github.com/Krab1o/meebin/internal/model/user/s_model"
	"github.com/Krab1o/meebin/internal/service"
	authHelper "github.com/Krab1o/meebin/internal/service/auth/helper"
	"github.com/Krab1o/meebin/internal/shared"
	"golang.org/x/crypto/bcrypt"
)

func registrationRole(email string) []model.Role {
	if email == service.AdminEmail {
		return []model.Role{service.RoleAdminName}
	} else {
		return []model.Role{service.RoleUserName}
	}
}

// TODO: add error messages
func (s *serv) Register(ctx context.Context, user *smodel.User) (*smodel.Tokens, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword(
		[]byte(user.Creds.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, service.NewInternalError(err)
	}
	user.Creds.Password = string(hashedBytes)

	user.Stats = &smodel.Stats{
		UtilizeCount: service.StartUtilizeCount,
		ReportCount:  service.StartReportCount,
		Rating:       service.StartRating,
	}

	user.Roles = registrationRole(user.Creds.Email)

	repoUser := converter.UserServiceToRepo(user)

	// TODO: add multiple roles
	roleId, err := s.roleRepository.ListRolesByTitles(ctx, user.Roles)
	if err != nil {
		return nil, service.ErrorDBToService(err)
	}

	var userId uint64
	err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		userId, err = s.userRepository.AddUser(ctx, repoUser, roleId)
		if err != nil {
			// TODO: fix wrong error
			// log.Println("===")
			// fmt.Println(errors.Is(err, repository.ErrDuplicate))
			// log.Println("===")
			return service.ErrorDBToService(err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Creating database row with session
	timeNow := time.Now()
	refreshExpirationTime := timeNow.Add(time.Duration(s.jwtConfig.RefreshTimeout()) * time.Hour)
	repoSession := &rmodel.Session{
		UserId:         userId,
		ExpirationTime: refreshExpirationTime,
	}
	sessionId, err := s.sessionRepository.AddSession(ctx, repoSession)
	if err != nil {
		return nil, service.ErrorDBToService(err)
	}

	// Write database data to refresh token
	refreshToken, err := authHelper.GenerateRefreshToken(
		shared.CustomRefreshFields{
			SessionID: sessionId,
		},
		refreshExpirationTime,
		timeNow,
		s.jwtConfig.Secret(),
	)
	if err != nil {
		return nil, service.NewInternalError(err)
	}

	// Generate access token
	accessToken, err := authHelper.GenerateAccessToken(
		shared.CustomAccessFields{
			UserID:    userId,
			SessionID: sessionId,
			Roles:     user.Roles,
		},
		timeNow,
		s.jwtConfig.Secret(),
		s.jwtConfig.AccessTimeout(),
	)
	if err != nil {
		return nil, service.NewInternalError(err)
	}

	return &smodel.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
