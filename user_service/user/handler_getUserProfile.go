package user

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/lucas-j-k/kube-go-microservices/user-service/httpTools"
)

// GetUserProfile returns the profile info for the current logged in user
func GetUserProfile(service *UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		userId, _ := ctx.Value("userId").(int64)

		profileResult, err := service.FindUserProfile(GetUserProfilePayload{
			Id: userId,
		})

		if err != nil {
			render.Render(w, r, httpTools.ErrInternalServer())
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, profileResult)
	}
}
