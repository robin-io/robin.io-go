package robin

import (
	"fmt"
	"mime/multipart"
	"os"
	"testing"
)

func TestRobin_CreateUserToken(t *testing.T) {
	robin := Robin{
		Secret: "NT-qBsdCDfPFYQkAKcfxMeNgSXvYTmqakOBVYRr",
		Tls:    true,
	}

	token, err := robin.CreateUserToken(UserToken{MetaData: map[string]interface{}{
		"name": "elvis",
	}})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(token)
}

func TestRobin_GetUserToken(t *testing.T) {
	robin := Robin{
		Secret: "NT-qBsdCDfPFYQkAKcfxMeNgSXvYTmqakOBVYRr",
		Tls:    true,
	}

	res, err := robin.GetUserToken("YFXOKVyKBGvHxuBaqKgDWOhE")

	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}

func TestRobin_SyncUserToken(t *testing.T) {
	robin := Robin{
		Secret: "NT-QuNtKolpzoWLahimkIjGAllEcJwGrymaVxQX",
		Tls:    true,
	}

	res, err := robin.SyncUserToken(UserToken{
		UserToken: "YFXOKVyKBGvHxuBaqKgDWOhE",
		MetaData: map[string]interface{}{
			"email": "elvis@acumen.com.ng",
		},
	})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}

// this will fail
func TestRobin_UpdateDisplayPhoto(t *testing.T) {
	robin := Robin{
		Secret: "NT-XmIzEmWUlsrQYypZOFRlogDFvQUsaEuxMfZf",
		Tls:    true,
	}

	path, err := os.Getwd()

	if err != nil {
		t.Error(err)
	}

	path += "/test.png"

	file, err := os.Open(path)

	if err != nil {
		t.Error(err)
	}

	file1, err := getFileHeader(file)

	if err != nil {
		t.Error(err)
	}

	res, err := robin.UpdateDisplayPhoto("FefXITDgAeTVrghcOHiimDVB", file1)

	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}

func getFileHeader(file *os.File) (*multipart.FileHeader, error) {
	// get file size
	fileStat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	// create *multipart.FileHeader
	return &multipart.FileHeader{
		Filename: fileStat.Name(),
		Size:     fileStat.Size(),
	}, nil
}
