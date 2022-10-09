package i18n

var config = new(ymlConfig)

type ymlConfig struct {
	I18n Yml
}

func (cfg *ymlConfig) initObj(obj Yml) {
	cfg.I18n = obj
	cfg.fillNull()
}

func (cfg *ymlConfig) fillNull() {
	if cfg.I18n.Duration == 0 {
		cfg.I18n.Duration = 360
	}
	if cfg.I18n.ServerName == "" {
		cfg.I18n.ServerName = "x-lang"
	}
	if cfg.I18n.CheckUri == "" {
		cfg.I18n.CheckUri = "/lang/string/app/version"
	}
	if cfg.I18n.QueryUri == "" {
		cfg.I18n.QueryUri = "/lang/string/list"
	}
}
