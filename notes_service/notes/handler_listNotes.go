package note

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/lucas-j-k/kube-go-microservices/notes-service/httpTools"
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
			fmt.Printf("ERR %v \n\n", err)
			render.Render(w, r, httpTools.ErrInternalServer())
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, noteRows)
	}
}
