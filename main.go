package main

import (
	"crypto"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/docker/distribution/registry/auth"
	"github.com/docker/distribution/registry/auth/token"

	"github.com/docker/libtrust"
)

type TokenIssuer struct {
	Issuer     string
	SigningKey libtrust.PrivateKey
	Expiration time.Duration
}

func main() {
	var (
		issuer   = &TokenIssuer{}
		pkFile   string
		service  string
		scope    string
		username string
		err      error
	)

	flag.StringVar(&issuer.Issuer, "issuer", "distribution-token-server", "Issuer string for token")
	flag.StringVar(&pkFile, "key", "", "Private key file")
	flag.StringVar(&service, "service", "", "service")
	flag.StringVar(&scope, "scope", "", "scope")
	flag.StringVar(&username, "username", "", "username")

	flag.Parse()

	issuer.SigningKey, err = libtrust.LoadKeyFile(pkFile)
	if err != nil {
		fmt.Printf("Error loading key file %s: %v\n", pkFile, err)
	}

	token, err := getToken(issuer, service, scope, username)

	fmt.Print(token)
}

func getToken(issuer *TokenIssuer, service string, scope string, username string) (string, error) {

	requestedAccessList := resolveScopeSpecifiers(scope)

	token, err := issuer.CreateJWT(username, service, requestedAccessList)
	if err != nil {
		return "", err
	}

	return token, nil
}

func resolveScopeSpecifiers(scope string) []auth.Access {
	requestedAccessSet := make(map[auth.Access]struct{}, 2)

	if scope != "" {
		parts := strings.SplitN(scope, ":", 3)

		if len(parts) != 3 {
			fmt.Printf("ignoring unsupported scope format %s", scope)
		}

		resourceType, resourceName, actions := parts[0], parts[1], parts[2]

		for _, action := range strings.Split(actions, ",") {
			requestedAccess := auth.Access{
				Resource: auth.Resource{
					Type: resourceType,
					Name: resourceName,
				},
				Action: action,
			}

			// Add this access to the requested access set.
			requestedAccessSet[requestedAccess] = struct{}{}
		}
	}

	requestedAccessList := make([]auth.Access, 0, len(requestedAccessSet))
	for requestedAccess := range requestedAccessSet {
		requestedAccessList = append(requestedAccessList, requestedAccess)
	}

	return requestedAccessList
}

// CreateJWT creates and signs a JSON Web Token for the given subject and
// audience with the granted access.
func (issuer *TokenIssuer) CreateJWT(subject string, audience string, grantedAccessList []auth.Access) (string, error) {
	// Make a set of access entries to put in the token's claimset.
	resourceActionSets := make(map[auth.Resource]map[string]struct{}, len(grantedAccessList))
	for _, access := range grantedAccessList {
		actionSet, exists := resourceActionSets[access.Resource]
		if !exists {
			actionSet = map[string]struct{}{}
			resourceActionSets[access.Resource] = actionSet
		}
		actionSet[access.Action] = struct{}{}
	}

	accessEntries := make([]*token.ResourceActions, 0, len(resourceActionSets))
	for resource, actionSet := range resourceActionSets {
		actions := make([]string, 0, len(actionSet))
		for action := range actionSet {
			actions = append(actions, action)
		}

		accessEntries = append(accessEntries, &token.ResourceActions{
			Type:    resource.Type,
			Name:    resource.Name,
			Actions: actions,
		})
	}

	randomBytes := make([]byte, 15)
	_, err := io.ReadFull(rand.Reader, randomBytes)
	if err != nil {
		return "", err
	}
	randomID := base64.URLEncoding.EncodeToString(randomBytes)

	now := time.Now()

	signingHash := crypto.SHA256
	var alg string
	switch issuer.SigningKey.KeyType() {
	case "RSA":
		alg = "RS256"
	case "EC":
		alg = "ES256"
	default:
		panic(fmt.Errorf("unsupported signing key type %q", issuer.SigningKey.KeyType()))
	}

	joseHeader := token.Header{
		Type:       "JWT",
		SigningAlg: alg,
	}

	if x5c := issuer.SigningKey.GetExtendedField("x5c"); x5c != nil {
		joseHeader.X5c = x5c.([]string)
	} else {
		var jwkMessage json.RawMessage
		jwkMessage, err = issuer.SigningKey.PublicKey().MarshalJSON()
		if err != nil {
			return "", err
		}
		joseHeader.RawJWK = &jwkMessage
	}

	exp := issuer.Expiration
	if exp == 0 {
		exp = 5 * time.Minute
	}

	claimSet := token.ClaimSet{
		Issuer:     issuer.Issuer,
		Subject:    subject,
		Audience:   audience,
		Expiration: now.Add(exp).Unix(),
		NotBefore:  now.Unix(),
		IssuedAt:   now.Unix(),
		JWTID:      randomID,

		Access: accessEntries,
	}

	var (
		joseHeaderBytes []byte
		claimSetBytes   []byte
	)

	if joseHeaderBytes, err = json.Marshal(joseHeader); err != nil {
		return "", fmt.Errorf("unable to encode jose header: %s", err)
	}
	if claimSetBytes, err = json.Marshal(claimSet); err != nil {
		return "", fmt.Errorf("unable to encode claim set: %s", err)
	}

	encodedJoseHeader := joseBase64Encode(joseHeaderBytes)
	encodedClaimSet := joseBase64Encode(claimSetBytes)
	encodingToSign := fmt.Sprintf("%s.%s", encodedJoseHeader, encodedClaimSet)

	var signatureBytes []byte
	if signatureBytes, _, err = issuer.SigningKey.Sign(strings.NewReader(encodingToSign), signingHash); err != nil {
		return "", fmt.Errorf("unable to sign jwt payload: %s", err)
	}

	signature := joseBase64Encode(signatureBytes)

	return fmt.Sprintf("%s.%s", encodingToSign, signature), nil
}

func joseBase64Encode(data []byte) string {
	return strings.TrimRight(base64.URLEncoding.EncodeToString(data), "=")
}
