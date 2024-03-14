package auth0client

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

type CustomClaimOrg struct {
	OrgId string `json:"org_id"`
}

func (c *CustomClaimOrg) Validate(ctx context.Context) error {
	if c.OrgId == "" {
		return errors.New("org_id is required")
	}
	return nil
}

var jwtValidator *validator.Validator
