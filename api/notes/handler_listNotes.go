package notes

import (
	"net/http"

	"github.com/lucas-j-k/kube-go-api/httpTools"

	"github.com/go-chi/render"
)

// ListNotesForUser returns all notes for the current logged in user
func ListNotesForUser(service *NoteService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		userId, _ := ctx.Value("userId").(int64)

		noteRows, err := service.ListNotesForUser(ListNotesByUserPayload{
			UserId: userId,
		})

		if err != nil {
			render.Render(w, r, httpTools.ErrInternalServer())
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, noteRows)
	}
}
