package url

type Gateway interface {
	FindByReference(string) URL
}

type MockGateway struct{}

func (g MockGateway) FindByReference(reference string) URL {
	return URL{}
}

type MysqlGateway struct{}

func (g MysqlGateway) FindByReference(reference string) URL {
	return URL{}
}
