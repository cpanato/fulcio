// Copyright 2023 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package codefresh

import (
	"context"
	"fmt"

	"github.com/sigstore/fulcio/pkg/config"
	"github.com/sigstore/fulcio/pkg/identity"
	"github.com/sigstore/fulcio/pkg/identity/base"
)

// Deprecated: Use ciprovider.ciProviderIssuer instead
type codefreshIssuer struct {
	identity.Issuer
}

// Deprecated: Use ciprovider.Issuer instead
func Issuer(issuerURL string) identity.Issuer {
	return &codefreshIssuer{base.Issuer(issuerURL)}
}

func (e *codefreshIssuer) Authenticate(ctx context.Context, token string, opts ...config.InsecureOIDCConfigOption) (identity.Principal, error) {
	idtoken, err := identity.Authorize(ctx, token, opts...)
	if err != nil {
		return nil, fmt.Errorf("authorizing codefresh issuer: %w", err)
	}
	return WorkflowPrincipalFromIDToken(ctx, idtoken)
}
