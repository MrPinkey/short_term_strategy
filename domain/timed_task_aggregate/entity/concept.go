package entity

type ConceptData struct {
	Ssbk []struct {
		Secucode         string `json:"SECUCODE"`
		SecurityCode     string `json:"SECURITY_CODE"`
		SecurityNameAbbr string `json:"SECURITY_NAME_ABBR"`
		BoardCode        string `json:"BOARD_CODE"`
		BoardName        string `json:"BOARD_NAME"`
		IsPrecise        string `json:"IS_PRECISE"`
		BoardRank        int    `json:"BOARD_RANK"`
	} `json:"ssbk"`
	Hxtc []struct {
		Secucode         string `json:"SECUCODE"`
		SecurityCode     string `json:"SECURITY_CODE"`
		SecurityNameAbbr string `json:"SECURITY_NAME_ABBR"`
		Keyword          string `json:"KEYWORD"`
		Mainpoint        int    `json:"MAINPOINT"`
		MainpointContent string `json:"MAINPOINT_CONTENT"`
		KeyClassif       string `json:"KEY_CLASSIF"`
		KeyClassifCode   string `json:"KEY_CLASSIF_CODE"`
		IsPoint          string `json:"IS_POINT"`
	} `json:"hxtc"`
}
