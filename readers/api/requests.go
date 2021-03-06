package api

type apiReq interface {
	validate() error
}

type listMessagesReq struct {
	chanID uint64
	offset uint64
	limit  uint64
}

func (req listMessagesReq) validate() error {
	if req.limit < 1 {
		return errInvalidRequest
	}

	return nil
}
