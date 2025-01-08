package helpers

import (
	"salimon/nexus/types"
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

func TestJWTString(t *testing.T) {
	sub := "someid"
	token, err := generateJWTString(jwt.MapClaims{"sub": sub})

	if err != nil {
		t.Fatal(err)
	}

	rSub, err := VerifyJWT(token)

	if err != nil {
		t.Fatal(err)
	}

	if rSub == nil {
		t.Fatal("sub is expected to be non-nil")
	}

	if *rSub != sub {
		t.Fatalf("verified sub is expected to be %s, actual = %s", sub, *rSub)
	}
}

func TestJWTTokens(t *testing.T) {
	user := types.User{
		Id: "someid",
	}

	accessToken, refreshToken, err := GenerateJWT(&user)
	if err != nil {
		t.Fatal(err)
	}
	if accessToken == nil {
		t.Fatal("access token is nil")
	}
	if refreshToken == nil {
		t.Fatal("refresh token is nil")
	}

	accessSub, err := VerifyJWT(*accessToken)
	if err != nil {
		t.Fatal(err)
	}
	if accessSub == nil {
		t.Fatal("access token sub is nil")
	}

	if *accessSub != user.Id {
		t.Fatalf("expected access token sub to be %s but it's %s", user.Id, *accessSub)
	}

	refreshSub, err := VerifyJWT(*refreshToken)
	if err != nil {
		t.Fatal(err)
	}
	if refreshSub == nil {
		t.Fatal("refresh token sub is nil")
	}

	if *refreshSub != user.Id {
		t.Fatalf("expected refresh token sub to be %s but it's %s", user.Id, *refreshSub)
	}

}
