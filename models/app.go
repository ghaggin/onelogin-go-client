package models

type App struct {
	ID                 int                  `json:"id,omitempty"`
	ConnectorID        int                  `json:"connector_id"`
	Name               string               `json:"name"`
	Description        string               `json:"description,omitempty"`
	Notes              string               `json:"notes,omitempty"`
	PolicyID           int                  `json:"policy_id,omitempty"`
	BrandID            int                  `json:"brand_id,omitempty"`
	IconURL            string               `json:"icon_url,omitempty"`
	Visible            bool                 `json:"visible,omitempty"`
	AuthMethod         int                  `json:"auth_method,omitempty"`
	TabID              int                  `json:"tab_id,omitempty"`
	CreatedAt          string               `json:"created_at,omitempty"`
	UpdatedAt          string               `json:"updated_at,omitempty"`
	RoleIDs            []int                `json:"role_ids,omitempty"`
	AllowAssumedSignin bool                 `json:"allow_assumed_signin,omitempty"`
	Provisioning       *Provisioning        `json:"provisioning,omitempty"`
	SSO                *SSO                 `json:"sso,omitempty"`
	Configuration      *Configuration       `json:"configuration,omitempty"`
	Parameters         map[string]Parameter `json:"parameters,omitempty"`
	EnforcementPoint   *EnforcementPoint    `json:"enforcement_point,omitempty"`
}

type Provisioning struct {
	Enabled bool `json:"enabled,omitempty"`
}

type SSO struct {
	ClientID     string       `json:"client_id,omitempty"`
	ClientSecret string       `json:"client_secret,omitempty"`
	MetadataURL  string       `json:"metadata_url,omitempty"`
	AcsURL       string       `json:"acs_url,omitempty"`
	SlsURL       string       `json:"sls_url,omitempty"`
	Issuer       string       `json:"issuer,omitempty"`
	Certificate  *Certificate `json:"certificate,omitempty"`
}

type Certificate struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type Configuration struct {
	RedirectURI                   string `json:"redirect_uri,omitempty"`
	RefreshTokenExpirationMinutes int    `json:"refresh_token_expiration_minutes,omitempty"`
	LoginURL                      string `json:"login_url,omitempty"`
	OidcApplicationType           int    `json:"oidc_application_type,omitempty"`
	TokenEndpointAuthMethod       int    `json:"token_endpoint_auth_method,omitempty"`
	AccessTokenExpirationMinutes  int    `json:"access_token_expiration_minutes,omitempty"`
	ProviderArn                   string `json:"provider_arn,omitempty"`
	IdpList                       string `json:"idp_list,omitempty"`
	SignatureAlgorithm            string `json:"signature_algorithm,omitempty"`
	LogoutURL                     string `json:"logout_url,omitempty"`
	PostLogoutRedirectURI         string `json:"post_logout_redirect_uri,omitempty"`
	Audience                      string `json:"audience,omitempty"`
	ConsumerURL                   string `json:"consumer_url,omitempty"`
	Login                         string `json:"login,omitempty"`
	Recipient                     string `json:"recipient,omitempty"`
	Validator                     string `json:"validator,omitempty"`
	RelayState                    string `json:"relaystate,omitempty"`
	Relay                         string `json:"relay,omitempty"`
	SAMLNotValidOnOrAafter        string `json:"saml_notonorafter,omitempty"`
	GenerateAttributeValueTags    string `json:"generate_attribute_value_tags,omitempty"`
	SAMLInitiaterID               string `json:"saml_initiater_id,omitempty"`
	SAMLNotValidBefore            string `json:"saml_notbefore,omitempty"`
	SAMLIssuerType                string `json:"saml_issuer_type,omitempty"`
	SAMLSignElement               string `json:"saml_sign_element,omitempty"`
	EncryptAssertion              string `json:"encrypt_assertion,omitempty"`
	SAMLSessionNotValidOnOrAfter  string `json:"saml_sessionnotonorafter,omitempty"`
	SAMLEncryptionMethodID        string `json:"saml_encryption_method_id,omitempty"`
	SAMLNameIDFormatID            string `json:"saml_nameid_format_id,omitempty"`
}

type Parameter struct {
	ID                        int    `json:"id,omitempty"`
	Label                     string `json:"label,omitempty"`
	UserAttributeMappings     string `json:"user_attribute_mappings,omitempty"`
	UserAttributeMacros       string `json:"user_attribute_macros,omitempty"`
	AttributesTransformations string `json:"attributes_transformations,omitempty"`
	Values                    string `json:"values,omitempty"`
	ProvisionedEntitlements   bool   `json:"provisioned_entitlements,omitempty"`
	SkipIfBlank               bool   `json:"skip_if_blank,omitempty"`
	DefaultValues             string `json:"default_values"`
	IncludeInSamlAssertion    bool   `json:"include_in_saml_assertion,omitempty"`
}

type EnforcementPoint struct {
	RequireSitewideAuthentication bool        `json:"require_sitewide_authentication"`
	Conditions                    *Conditions `json:"conditions,omitempty"`
	SessionExpiryFixed            *Duration   `json:"session_expiry_fixed"`
	SessionExpiryInactivity       *Duration   `json:"session_expiry_inactivity"`
	Permissions                   string      `json:"permissions"`
	Token                         string      `json:"token,omitempty"`
	Target                        string      `json:"target"`
	Resources                     []*Resource `json:"resources"`
	ContextRoot                   string      `json:"context_root"`
	UseTargetHostHeader           bool        `json:"use_target_host_header"`
	Vhost                         string      `json:"vhost"`
	LandingPage                   string      `json:"landing_page"`
	CaseSensitive                 bool        `json:"case_sensitive"`
}

type Conditions struct {
	Type  string   `json:"type"`
	Roles []string `json:"roles"`
}

type Duration struct {
	Value int `json:"value"`
	Unit  int `json:"unit"`
}

type Resource struct {
	Path        string  `json:"path"`
	RequireAuth string  `json:"require_authentication"`
	Permissions string  `json:"permissions"`
	Conditions  *string `json:"conditions,omitempty"`
	IsPathRegex *bool   `json:"is_path_regex,omitempty"`
	ResourceID  int     `json:"resource_id,omitempty"`
}
