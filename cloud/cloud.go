package cloud

type CloudDB struct {
	url string
}

func NewCloudDB(url string) *CloudDB {
	return &CloudDB{
		url: url,
	}
}

// Write запись файла
func (db *CloudDB) Write(content []byte) {

}

// Read чтение файла
func (db *CloudDB) Read() ([]byte, error) {

	return []byte{}, nil
}
