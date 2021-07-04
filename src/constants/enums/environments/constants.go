package environments

type Environment struct {
	name         string
	isProduction bool
	isDocker     bool
	isDev        bool
}

func (env *Environment) String() string {
	return env.name
}

func (env *Environment) IsProduction() bool {
	return env.isProduction
}

func (env *Environment) IsDocker() bool {
	return env.isDocker
}

func (env *Environment) IsDev() bool {
	return env.isDev
}

var (
	Dev = Environment{
		name:  "dev",
		isDev: true,
	}

	Prod = Environment{
		name: "prod",
		isDocker:     true,
		isProduction: true,
	}
)

func GetEnvironment(env string) *Environment {
	switch env {
	case "dev":
		return &Dev
	case "prod":
		return &Prod
	default:
		return &Dev
	}
}
