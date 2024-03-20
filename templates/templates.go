package templates

type Lang string

const (
	LangEN Lang = "en"
	LangFR Lang = "fr"
	LangES Lang = "es"
	LangPT Lang = "pt"
	LangDE Lang = "de"
	LangIT Lang = "it"
)

type Content struct {
	Subject string `json:"Subject" bson:"Subject" yaml:"Subject"`
	Body    string `json:"Body" bson:"Body" yaml:"Body"`
}
