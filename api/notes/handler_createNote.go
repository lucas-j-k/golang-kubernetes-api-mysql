package notes

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/lucas-j-k/kube-go-api/httpTools"
)

// Signup inserts a new user record into the DB
func CreateNote(service *NoteService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		userId, _ := ctx.Value("userId").(int64)

		reqBody := &MutateNoteRequestBody{}
		if err := render.Bind(r, reqBody); err != nil {
			render.Render(w, r, httpTools.ErrInvalidRequest())
			return
		}

		insertedId, err := service.InsertNote(InsertNotePayload{
			UserID: userId,
			Label:  reqBody.Label,
			Body:   reqBody.Body,
		})

		if err != nil {
			render.Render(w, r, httpTools.ErrInternalServer())
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, insertedId)
	}
}
