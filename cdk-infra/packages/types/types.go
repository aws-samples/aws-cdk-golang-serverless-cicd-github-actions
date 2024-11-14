package types

type Config struct {
	LambdaVariables LambdaVariables   `yaml:"LambdaVariables"`
	Common_Tags     map[string]string `yaml:"Common_Tags"`
}

type LambdaEnvVar struct {
	ENV   string `yaml:"ENV"`
	LOGLEVEL string `yaml:"LOG_LEVEL"`
	APPNAME  string `yaml:"APP_NAME"`
}
type LambdaVariables struct {
	LambdaEnvVar  map[string]string `yaml:"LambdaEnvVar"`
	Repo          string            `yaml:"Repo"`
	Name          string            `yaml:"Name"`
	AccountNumber string            `yaml:"Account_Number"`
	Region        string            `yaml:"Region"`
	Timeout       int32             `yaml:"Timeout"`
}

type Tags struct {
	ApplicationName    string `yaml:"ApplicationName"`
	ManagedBy          string `yaml:"ManagedBy"`
	Environment        string `yaml:"Environment"`
}
