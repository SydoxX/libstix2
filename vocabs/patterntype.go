package vocabs

type PatternType string

const (
	PatternStix     PatternType = "stix"
	PatternPcre     PatternType = "pcre"
	PatternSigma    PatternType = "sigma"
	PatternSnort    PatternType = "snort"
	PatternSuricata PatternType = "suricata"
	PatternYara     PatternType = "yara"
)
