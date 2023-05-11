package jwe

import (
	"d2din.com/internal/pkg/jwe"
	"d2din.com/internal/pkg/log"
	"d2din.com/internal/pkg/util"
	model "d2din.com/internal/v1/structure/jwe"
	"time"
)

func (s service) Created(input *model.JWE) (output *model.Token, err error) {
	// TODO implement me
	other := map[string]interface{}{
		"company_id": input.CompanyID,
		"account_id": input.AccountID,
		"name":       input.Name,
	}

	publicKey := "-----BEGIN PUBLIC KEY-----\nMIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAuNH+0eswuLx0SPBZVfpA\nm7B3rxp7eb/pOzyPM5n2d+/beZgwGfnkUsmDM35xGk862JY4z4lz96LAvHQ5Vq5l\nD3iY7dSxnsH2q9wRCkZVgo/wkes655Rx6nG/qRdpVFqT67Z0jO21xYCGtOL6BpOr\n/siSUCCh/ABpLUcXRUd5Ra5nWds8T9+Iash0TMgAVohg3SAFCHw9qyUT0DYQKvNj\npx391mZVGgviMr3/o9koC4J/3DgAUg+8QSGOGU8fkXsMfyyF9c/9xvj5vulA4o1S\nFiCYw+cmBKEYT5r0Kype0jVlCUtE5kSVnjWtCCGLHps62bgDFNer83mWcd94mSSf\nESxhaoFhISlJ8QApcNP+etljGsnN5+/Og30c1rBVykK/gbiTTXjuV41JeiGcAFZn\ndkKAsjoOXOo77chQ0Lzj1NrKhg+xZ+KLa1s6pkxnHvNbyQ/KLy6dengnmxV794TZ\nF6neQqf+8YjB5r1a0aMS+sCMiKvv4xitgP7i/IRhlAqUJVcvVz/0vvv9G6CciSVD\nY61X/uNiveTgsFh4e/zl6xwQsqu69JZ0g4JOkyVS6GZZG5veBR0A2+aWY5pZboUM\njtXvDw9k+47VgXmBbLWyT6NVDZV2gPgNLOimSeqvMmk2XORjUqmcqj0GxhGxWvda\nmibtWWEY7iJkvN1U/PG5PoMCAwEAAQ==\n-----END PUBLIC KEY-----"

	accessExpiration := util.NowToUTC().Add(time.Minute * 5).Unix()
	j := &jwe.JWT{
		PublicKey:     publicKey,
		Other:         other,
		ExpirationKey: accessExpiration,
	}

	j, err = j.Created()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	accessToken := j.Token
	refreshTokenExpiration := util.NowToUTC().Add(time.Hour * 8).Unix()
	j.ExpirationKey = refreshTokenExpiration
	j, err = j.Created()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	refreshToken := j.Token
	output = &model.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return output, nil
}
