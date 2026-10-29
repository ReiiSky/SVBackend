package dto_test

import (
	"testing"

	"github.com/ReiiSky/SharingVisionBackend/application/dto"
	"github.com/go-playground/validator/v10"
)

func TestCreateArticleValue(t *testing.T) {
	input := dto.CreateArticleInput{
		Title:    "What is Lorem Ipsum?",
		Content:  "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since 1966, when designers at Letraset and James Mosley, the librarian at St Bride Printing Library in London, took a 1914 Cicero translation and scrambled it to make dummy text for Letraset's Body Type sheets. It has survived not only many decades, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised thanks to these sheets and more recently with desktop publishing software like Aldus PageMaker and Microsoft Word including versions of Lorem Ipsum.",
		Category: "Lorem",
		Status:   "draft",
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if got := dto.Validate(validate, input); got != nil {
		t.Errorf("expected valid input, got error: %v (field: %s)", got.Detail(), got.Msg())
	}
}
