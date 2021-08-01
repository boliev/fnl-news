package parser

import log "github.com/sirupsen/logrus"

// GetParsers returns parsers list
func GetParsers(config map[string]Config) []Parser {
	var parsers []Parser
	if sportboxConfig, ok := config["sportbox"]; ok {
		sportbox := newSportboxParser(sportboxConfig)
		parsers = append(parsers, sportbox)
	} else {
		log.Warnf("Unable to find config for sportbox parser")
	}

	if onefnlConfig, ok := config["onefnl"]; ok {
		onefnl := newOnefnlParser(onefnlConfig)
		parsers = append(parsers, onefnl)
	} else {
		log.Warnf("Unable to find config for onefnl parser")
	}

	if sportsruConfig, ok := config["sportsru"]; ok {
		sportsru := newSportsruParser(sportsruConfig)
		parsers = append(parsers, sportsru)
	} else {
		log.Warnf("Unable to find config for sportsru parser")
	}

	if kulichkiConfig, ok := config["kulichki"]; ok {
		kulichki := newKulichkiParser(kulichkiConfig)
		parsers = append(parsers, kulichki)
	} else {
		log.Warnf("Unable to find config for sportsru parser")
	}

	return parsers
}
