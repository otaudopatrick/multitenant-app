package request

type TenantRequest struct {
	Name      string `json:"name"`
	Subdomain string `json:"subdomain"`
}

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserWithTenantRequest struct {
	Tenant TenantRequest `json:"tenant"`
	User   UserRequest   `json:"user"`
}
