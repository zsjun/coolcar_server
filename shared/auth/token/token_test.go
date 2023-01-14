package token

import (
	"fmt"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const publicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAu1SU1LfVLPHCozMxH2Mo
4lgOEePzNm0tRgeLezV6ffAt0gunVTLw7onLRnrq0/IzW7yWR7QkrmBL7jTKEn5u
+qKhbwKfBstIs+bMY2Zkp18gnTxKLxoS2tFczGkPLPgizskuemMghRniWaoLcyeh
kd3qqGElvW/VDL5AaWTg0nLVkjRo9z+40RQzuVaE8AkAFmxZzow3x+VJYKdjykkJ
0iT9wCS0DRTXu269V264Vf/3jvredZiKRkgwlL9xNAwxXFg0x/XFw005UWVRIkdg
cKWTjpBP2dPwVZ4WWC+9aGVd+Gyn1o0CLelf4rEjGoXbAAEgAqeGUxrcIlbjXfbc
mwIDAQAB
-----END PUBLIC KEY-----`

func TestVerify(t *testing.T) {
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	if err != nil {
		t.Fatalf("cannot parse public key: %v", err)
	}
	v := &JWTTokenVerifier{
		PublicKey: pubKey,
	}

	cases := []struct {
		name    string
		tkn     string
		now     time.Time
		want    string
		wantErr bool
	}{
		{
			name: "valid_token",
			tkn:  "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyMzk3MjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiNjNjMTFiMzRiYTI2NGViMGVhOWVjMTRhIn0.O7d8glbn-4AheMCBY7T7QNLJQV8rBmfadE5xRk53fdrHO084aA1FxUB4uVugUJBXaxN_IKmio1ZodORSuZlSDx1C24i0c8PRI9ciyOmmR5abXYTyovpwCqe4WQvagNwIfqYeBgWCg-RBnvmXNAE5gis90xSKrBnB4FfvvjmrdhEcuD7bM0khNeNRVkcSM3_BDtuqCshWclz9i4-L39cvQTQ8bH7VMN_qKNI3e_k0D52HgPeZ_zwEax-18gX1UCbwJ8R0HXbk71vBlzJB0fjyZFgsbA9b8Ukr39KWyHTCOINEM8IK_VjwWiOXC9tPu-cIiphXYDq_nLPKTTbLULtr2g",
			now:  time.Unix(1516239122, 0),
			want: "63c11b34ba264eb0ea9ec14a",
		},
		{
			name:    "token_expired",
			tkn:     "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjcyMDE2NzQ2NzY5ODAsImlhdCI6MTY3MzY3Njk4MCwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiNjNjMTFiMzRiYTI2NGViMGVhOWVjMTRhIn0.hH98qq2Pam1nWD0-07P-ZwMvXY9MOrVAO6--vkHvtr2_HD_VQccndJR28eG4z3carU-2SyUM47cwPtTSDj8sqE2NPO_4pzEaOCUKyjOf4Irk--eEqR213DtCh1qBrjk2O7QVm2Ij5ak8cGMTaBO-veWdNhR7iA6SuPSl0aUgAJ3WD-oMAKELHHLXPV5ryw3Jm1SS2aDUbrlXiCOdtcpXRepmJ8KM9go2aqh6VyfXni8es7Rooz89JmjfS-h6QAQZvX4da6nzRb0ORVoPseH0p8RoZFFDX6MOfsVzLzAxT02a5kBxd5iqw71KCt0ltIoEpKto5F9LvJ2mRCug-J8w75ig",
			now:     time.Unix(1517239122, 0),
			wantErr: true,
		},
		{
			name:    "bad_token",
			tkn:     "bad_token",
			now:     time.Unix(1517239122, 0),
			wantErr: true,
		},
		{
			name:    "wrong_signature",
			tkn:     "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiNWY3YzMxNjhlMjI4M2FhNzIyZTM1MWE0In0.jPVRIZXsNz08OCudP4cC8KGzVEIWC42TOMHpc6cN-_3yUgbPcrhuJL6C27fzoxt0j8J3L0z6nv0ni_7aa13fzYjo1Y_b4Axxz4sI5bz-b9O1BziFU1NC9t3IJbwFsF2Svz2OpG3aY388rTZ4orHShfRbrzGnzK8NbNXIZ7CcCvEznHiJEmSgqSZSYeZVjjid2p2l_T_eTQxJTkHi9LE-3g_AfLKLXXmqLlXYpurTGMWEBkJq51uNs6MnESi4pEwbLviTmZTTtC6qAhkVmeJh7QUZA8BPKoxSbNEYQxYYQK1aiRGyrrsdsONsK1etXW6JG2F4x0wiNjTKMvQSAsq7GnWvkoQ",
			now:     time.Unix(1516239122, 0),
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			jwt.TimeFunc = func() time.Time {
				return c.now
			}
			accountID, err := v.Verify(c.tkn)
			if c.want != "" && accountID == c.want  {
				fmt.Println("999",accountID,c.want)
			}
			
			if !c.wantErr && err != nil {
				t.Errorf("verification failed: %v", err)
			}

			if c.wantErr && err == nil {
				t.Errorf("want error; got no error")
			}

			if accountID != c.want {
				t.Errorf("wrong account id. want: %q, got: %q", c.want, accountID)
			}
		})
	}
}
