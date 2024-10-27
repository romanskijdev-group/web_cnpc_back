package restauthcore

import (
	usersdb "cnpc_backend/core/module/user/users/db"
	"cnpc_backend/core/typescore"
	"context"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type CheckerRestAuth struct {
	TokenRestAuth *TokenRestAuth
	ConfigG       *typescore.Config
	UsersDB       usersdb.UsersProviderControlsDBI
}

func (m *CheckerRestAuth) ControlAuthRest(r *http.Request, params *typescore.ControlAuthRestParams) (*http.Request, *typescore.UsersProviderControl, *typescore.WEvent) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if params.UserAuthorizationChecked {
		return m.authRequestProcess(ctx, r, params)
	}

	return r, nil, nil
}

func (m *CheckerRestAuth) authRequestProcess(ctx context.Context, r *http.Request, params *typescore.ControlAuthRestParams) (*http.Request, *typescore.UsersProviderControl, *typescore.WEvent) {
	authHeader := r.Header.Get("Authorization")
	authData, errW := m.parseAuthorizationHeader(authHeader)
	if errW != nil {
		log.Println("💔 error parse auth header", errW)
		return nil, nil, errW
	}

	userSystemID, userRoleToken, expTime, errW := m.TokenRestAuth.GetValuesFromToken(authData, m.ConfigG.Secure.JWTSecret)
	if errW != nil {
		log.Println("💔 error get values from token", errW)
		return nil, nil, errW
	}

	userObj, errW := m.UsersDB.GetUserDB(ctx, &typescore.UsersProviderControl{
		SystemID: &userSystemID,
	})
	if errW != nil {
		log.Println("💔 error get user from db", errW)
		return nil, nil, errW
	}
	if userObj == nil {
		return nil, nil, &typescore.WEvent{
			Err:  errors.New("user not found"),
			Text: "invalid_auth_user_token",
		}
	}

	// Проверка, заблокирован ли пользователь
	if userObj.IsBlocked != nil {
		if *userObj.IsBlocked {
			return nil, nil, &typescore.WEvent{
				Err:  errors.New("user_blocked"),
				Text: "invalid_auth_user_token",
			}
		}
	}

	if params.RoleCheck {
		if errW = m.checkUserRole(params.EnabledRoles, userRoleToken, userObj); errW != nil {
			log.Println("💔 error check user role", errW)
			return nil, nil, errW
		}
	}
	if errW = m.checkTokenExpiration(expTime); errW != nil {
		return nil, nil, errW
	}

	return r, userObj, nil
}

func (m *CheckerRestAuth) CheckSocketAuth(ctx context.Context, authHeader *string) (*typescore.UsersProviderControl, *typescore.WEvent) {
	if authHeader == nil {
		return nil, &typescore.WEvent{
			Err:  errors.New("invalid auth header"),
			Text: "invalid_auth_header",
		}
	}

	authData, errW := m.parseAuthorizationHeader(*authHeader)
	if errW != nil {
		log.Println("💔 error parse auth header", errW)
		return nil, errW
	}

	userSystemID, _, expTime, errW := m.TokenRestAuth.GetValuesFromToken(authData, m.ConfigG.Secure.JWTSecret)
	if errW != nil {
		log.Println("💔 error get values from token", errW)
		return nil, errW
	}

	userObj, errW := m.UsersDB.GetUserDB(ctx, &typescore.UsersProviderControl{
		SystemID: &userSystemID,
	})
	if errW != nil {
		log.Println("💔 error get user from db", errW)
		return nil, errW
	}
	if userObj == nil {
		return nil, &typescore.WEvent{
			Err:  errors.New("user not found"),
			Text: "invalid_auth_user_token",
		}
	}

	// Проверка, заблокирован ли пользователь
	if userObj.IsBlocked != nil {
		if *userObj.IsBlocked {
			return nil, &typescore.WEvent{
				Err:  errors.New("user_blocked"),
				Text: "invalid_auth_user_token",
			}
		}
	}

	if errW = m.checkTokenExpiration(expTime); errW != nil {
		return nil, errW
	}

	return userObj, nil
}

func (m *CheckerRestAuth) checkUserRole(enabledRoles []*string, userRole string, userObj *typescore.UsersProviderControl) *typescore.WEvent {
	if len(enabledRoles) > 0 {

		if userObj.Role == nil {
			return &typescore.WEvent{
				Err:  errors.New("user role not found"),
				Text: "invalid_auth_user_token",
			}
		}
		userRoleFact := string(*userObj.Role)
		if userRoleFact != userRole {
			return &typescore.WEvent{
				Err:  errors.New("user role not allowed"),
				Text: "invalid_auth_user_token",
			}
		}
		for _, role := range enabledRoles {
			if *role == userRole {
				return nil
			}
		}

		return &typescore.WEvent{
			Err:  errors.New("user role not allowed"),
			Text: "invalid_auth_user_token",
		}
	}
	return nil
}

func (m *CheckerRestAuth) checkTokenExpiration(expTime int64) *typescore.WEvent {
	if expTime < time.Now().UTC().Unix() {
		return &typescore.WEvent{
			Err:  errors.New("token has expired"),
			Text: "invalid_auth_token_expired",
		}
	}
	return nil
}

func (m *CheckerRestAuth) parseAuthorizationHeader(authHeader string) (authData string, errW *typescore.WEvent) {
	splitAuthHeader := strings.Split(authHeader, "Bearer ")
	if len(splitAuthHeader) != 2 {
		return authData, &typescore.WEvent{
			Err:  errors.New("not valid user token"),
			Text: "invalid_auth_user_token",
		}
	}

	authDataString, err := url.QueryUnescape(splitAuthHeader[1])
	if err != nil {
		return authData, errW
	}

	return authDataString, nil
}
