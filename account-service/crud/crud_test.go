package crud_test

import (
	"account-service/crud"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMain(m *testing.M) {
	setDevEnv()
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	os.Chdir(dir)

	crud.Init()

	os.Exit(m.Run())
}

func TestCrudCreateDeleteUser(t *testing.T) {
	u_ref := crud.User{
		Email:       "testmail@mail.hui",
		Credentials: "abobus",
	}

	u_created, err := crud.CreateUser(u_ref)
	if err != nil {
		t.Errorf("[TestCrudCreateDeleteUser]: %v", err)
	}

	u_returned, err := crud.DeleteUserByID(u_created)
	if err != nil {
		t.Errorf("[TestCrudCreateDeleteUser]: %v", err)
	}

	if !cmp.Equal(u_created, u_returned) {
		t.Errorf(
			"[TestCrudCreateDeleteUser]: expected results mismatched.\n"+
				"Expected: %v\nGot: %v\n",
			u_created,
			u_returned,
		)
	}
}

// TODO remove test redis keys
func TestUserRegistration(t *testing.T) {
	u_ref := crud.User{
		Email:       "testmail@mail.hui",
		Credentials: "abobus",
	}

	s, err := crud.UserRegistration(u_ref)
	defer crud.DeleteUserByID(crud.User{UserID: s.User.UserID})
	t.Log(s)
	if err != nil {
		t.Error(err)
	}

	// Secret withdrawal test
	s2 := crud.Session{Token: s.Token}
	err = s2.GetSecret()
	if err != nil {
		t.Error(err)
	}

	if s.Secret != s2.Secret {
		t.Error("authentication failed")
	}
}

func TestCheckRequest(t *testing.T) {
	u_ref := crud.User{
		Email:       "testmail@mail.hui",
		Credentials: "abobus",
	}

	s, err := crud.UserRegistration(u_ref)
	defer crud.DeleteUserByID(crud.User{UserID: s.User.UserID})
	t.Log(s)
	if err != nil {
		t.Error(err)
	}
	s.GetSecret()
	if err != nil {
		t.Error(err)
	}

	postBody, err := json.Marshal(map[string]string{
		"name":    "Biba",
		"surname": "Boba",
	})
	if err != nil {
		t.Error(err)
	}

	var req = crud.AuthRequest{
		Base64Body: base64.StdEncoding.EncodeToString(postBody),
		Session:    s,
	}

	strToHash := fmt.Sprintf("%s:%s", req.Session.Secret, req.Base64Body)
	hashed := sha256.Sum256([]byte(strToHash))
	req.Base64Signature = base64.StdEncoding.EncodeToString(hashed[:])

	if !crud.AuthenticateRequest(req).Authenticated {
		t.Error("check payload failure")
	}
}

func setDevEnv() {
	// Postgres
	err := os.Setenv("DB_ADDR", "localhost:8081")
	if err != nil {
		panic(err)
	}
	err = os.Setenv("DB_DATABASE", "postgres")
	if err != nil {
		panic(err)
	}
	err = os.Setenv("DB_USER", "postgres")
	if err != nil {
		panic(err)
	}
	err = os.Setenv("DB_PASSWORD", "postgres")
	if err != nil {
		panic(err)
	}

	//Redis
	err = os.Setenv("REDIS_ADDR", "localhost:8090")
	if err != nil {
		panic(err)
	}
	err = os.Setenv("REDIS_PASSWORD", "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81")
	if err != nil {
		panic(err)
	}
	err = os.Setenv("REDIS_DB", "0")
	if err != nil {
		panic(err)
	}
}
