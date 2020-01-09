package main

type Quiz struct {
	ID        int             `json:"id"`
	Name      string          `json:"name"`
	Image     string          `json:"image"`
	Questions []*QuizQuestion `json:"questions"`
}

type QuizQuestion struct {
	Answers          []string `json:"answers"`
	Correct          string   `json:"correct"`
	Question         string   `json:"question"`
	Image            string   `json:"image"`
	Time             int      `json:"time"`
	PointsMultiplier int      `json:"points-multiplier"`
}

type QuizAnswer struct {
	Label string `json:"label"`
	Image string `json:"image"`
}

type Data struct {
	Card   DataCard   `json:"card"`
	Kahoot DataKahoot `json:"kahoot"`
}

type DataCard struct {
	Access              DataCardAccess           `json:"access"`
	Combined            bool                     `json:"combined"`
	CompatibilityLevel  int                      `json:"compatibility_level"`
	Cover               string                   `json:"cover"`
	CoverMetadata       DataCardCoverMetadata    `json:"coverMetadata"`
	Created             float64                  `json:"created"`
	Creator             string                   `json:"creator"`
	CreatorUsername     string                   `json:"creator_username"`
	Description         string                   `json:"description"`
	Draft               bool                     `json:"draft"`
	DraftExists         bool                     `json:"draftExists"`
	DuplicationDisabled bool                     `json:"duplication_disabled"`
	Featured            bool                     `json:"featured"`
	LastEdit            DataCardLastEdit         `json:"last_edit"`
	Locked              bool                     `json:"locked"`
	Modified            float64                  `json:"modified"`
	NumberOfPlayers     int                      `json:"number_of_players"`
	NumberOfPlays       int                      `json:"number_of_plays"`
	NumberOfQuestions   int                      `json:"number_of_questions"`
	QuestionTypes       []string                 `json:"question_types"`
	SampleQuestions     []DataCardSampleQuestion `json:"sample_questions"`
	Slug                string                   `json:"slug"`
	Sponsored           bool                     `json:"sponsored"`
	Title               string                   `json:"title"`
	TotalFavourites     int                      `json:"total_favourites"`
	Type                string                   `json:"type"`
	UUID                string                   `json:"uuid"`
	Visibility          int                      `json:"visibility"`
	WriteProtection     bool                     `json:"writeProtection"`
	YoungFeatured       bool                     `json:"young_featured"`
}

type DataCardAccess struct {
	Features []interface{} `json:"features"`
}

type DataCardCoverMetadata struct {
	ID string `json:"id"`
}

type DataCardLastEdit struct {
	EditTimestamp  float64 `json:"editTimestamp"`
	EditorUserID   string  `json:"editorUserId"`
	EditorUsername string  `json:"editorUsername"`
}

type DataCardSampleQuestion struct {
	Image         string                              `json:"image"`
	ImageMetadata DataCardSampleQuestionImageMetadata `json:"imageMetadata"`
	Title         string                              `json:"title"`
	Type          string                              `json:"type"`
}

type DataCardSampleQuestionImageMetadata struct {
	Effects []interface{} `json:"effects"`
	ID      string        `json:"id"`
}

type DataKahoot struct {
	Audience            string                  `json:"audience"`
	CompatibilityLevel  int                     `json:"compatibilityLevel"`
	Cover               string                  `json:"cover"`
	CoverMetadata       DataKahootCoverMetadata `json:"coverMetadata"`
	Created             float64                 `json:"created"`
	Creator             string                  `json:"creator"`
	CreatorPrimaryUsage string                  `json:"creator_primary_usage"`
	CreatorUsername     string                  `json:"creator_username"`
	Description         string                  `json:"description"`
	FolderID            string                  `json:"folderId"`
	Language            string                  `json:"language"`
	Metadata            DataKahootMetadata      `json:"metadata"`
	Modified            float64                 `json:"modified"`
	Parent              DataKahootParent        `json:"parent"`
	Questions           []DataKahootQuestion    `json:"questions"`
	QuizType            string                  `json:"quizType"`
	Resources           string                  `json:"resources"`
	Slug                string                  `json:"slug"`
	Title               string                  `json:"title"`
	Type                string                  `json:"type"`
	UUID                string                  `json:"uuid"`
	Visibility          int                     `json:"visibility"`
}

type DataKahootCoverMetadata struct {
	ID string `json:"id"`
}

type DataKahootMetadata struct {
	Access                DataKahootMetadataAccess   `json:"access"`
	DuplicationProtection bool                       `json:"duplicationProtection"`
	LastEdit              DataKahootMetadataLastEdit `json:"lastEdit"`
}

type DataKahootMetadataAccess struct {
	Features []interface{} `json:"features"`
}

type DataKahootMetadataLastEdit struct {
	EditTimestamp  float64 `json:"editTimestamp"`
	EditorUserID   string  `json:"editorUserId"`
	EditorUsername string  `json:"editorUsername"`
}

type DataKahootParent struct {
	CreatorUsername string `json:"creator_username"`
	ID              string `json:"id"`
}

type DataKahootQuestion struct {
	Choices          []DataKahootQuestionChoice      `json:"choices"`
	Image            string                          `json:"image"`
	ImageMetadata    DataKahootQuestionImageMetadata `json:"imageMetadata"`
	NumberOfAnswers  int                             `json:"numberOfAnswers"`
	Points           bool                            `json:"points"`
	PointsMultiplier int                             `json:"pointsMultiplier"`
	Question         string                          `json:"question"`
	QuestionFormat   int                             `json:"questionFormat"`
	Resources        string                          `json:"resources"`
	Time             int                             `json:"time"`
	Type             string                          `json:"type"`
	Video            DataKahootQuestionVideo         `json:"video"`
}

type DataKahootQuestionChoice struct {
	Answer  string `json:"answer"`
	Correct bool   `json:"correct"`
}

type DataKahootQuestionImageMetadata struct {
	ContentType string        `json:"contentType"`
	Effects     []interface{} `json:"effects"`
	ExternalRef string        `json:"externalRef"`
	Height      int           `json:"height"`
	ID          string        `json:"id"`
	Origin      string        `json:"origin"`
	Width       int           `json:"width"`
}

type DataKahootQuestionVideo struct {
	EndTime   float64 `json:"endTime"`
	FullURL   string  `json:"fullUrl"`
	ID        string  `json:"id"`
	Service   string  `json:"service"`
	StartTime float64 `json:"startTime"`
}
