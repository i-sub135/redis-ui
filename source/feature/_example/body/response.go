package body

// {{FEATURE_CAMEL}}Data example data structure
type {{FEATURE_CAMEL}}Data struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// {{FEATURE_CAMEL}}Response example response structure
type {{FEATURE_CAMEL}}Response struct {
	Status     string      `json:"status"`
	Message    *string     `json:"message,omitempty"`
	Data       {{FEATURE_CAMEL}}Data `json:"data"`
	Time       string      `json:"timestamp"`
	AppVersion string      `json:"app_version"`
}