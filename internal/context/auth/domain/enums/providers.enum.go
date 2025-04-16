package enums

type UserProvider string

const (
	Google UserProvider = "google"
	Email  UserProvider = "email"
	Github UserProvider = "github"
)

func (up UserProvider) String() string {
	return string(up)
}

func GetAllProviders() []UserProvider {
	return []UserProvider{Google, Email, Github}
}

func GetAllProvidersStrings() []string {
	return []string{string(Google), string(Email), string(Github)}
}
