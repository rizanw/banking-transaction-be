package module

import (
	hUtils "tx-bank/internal/handler/http/utils"
	ucUtils "tx-bank/internal/usecase/utils"
)

type handler struct {
	ucUtils ucUtils.UseCase
}

func New(ucUtils ucUtils.UseCase) hUtils.Handler {
	return &handler{
		ucUtils: ucUtils,
	}
}
