package request

type GetStatRequest struct {
	From int64 `query:"tsFrom" validate:"required"`
	To   int64 `query:"tsTo" validate:"required"`
}
