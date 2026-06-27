
// Claims are simply the data stored inside a JWT.

// Think of a JWT as an envelope.

// Header → tells you how it was signed.
// Payload (Claims) → contains information about the user.
// Signature → proves the token hasn't been tampered with.
// JWT
// │
// ├── Header
// ├── Claims (Payload)
// └── Signature
// Example

// Suppose user 24 logs in.

// The server creates a JWT whose payload might look like:

// {
//   "user_id": 24,
//   "exp": 1782446400,
//   "iat": 1782360000
// }

// These key-value pairs are the claims.

// Registered Claims

// The JWT specification defines some standard claim names.

// Claim	Meaning
// sub	Subject (usually user id)
// exp	Expiration time
// iat	Issued at
// iss	Issuer
// aud	Audience
// nbf	Not valid before




// Custom Claims

// You can also add your own information.

// For example:

// {
//   "user_id": 24,
//   "role": "admin",
//   "company": "OpenAI",
//   "exp": 1782446400
// }

// These are custom claims.


package auth

import "context"

type contextKey string

const ClaimsContextKey contextKey = "claims"

func WithClaims(
	ctx context.Context,
	claims *Claims,
) context.Context {

	return context.WithValue(
		ctx,
		ClaimsContextKey,
		claims,
	)
}

func ForContext(
	ctx context.Context,
) *Claims {

	claims, ok := ctx.Value(
		ClaimsContextKey,
	).(*Claims)

	if !ok {
		return nil
	}

	return claims
}
