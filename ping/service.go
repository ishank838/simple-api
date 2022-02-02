package ping

type pingService struct {
	store pingStore
}

func (svc *pingService) PingDb() (string, error) {
	message, err := svc.store.pindDB()

	if err != nil {
		return "", err
	}
	return message, nil
}
