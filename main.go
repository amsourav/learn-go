package main

import (
	"fmt"
	"time"
)

//go:generate stringer -type=MoodState
type MoodState int

const (
	MoodStateNeutral MoodState = iota
	MoodStateHappy
	MoodStateSad
	MoodStateAngry
)

type AuditableContent struct {
	TimeCreated  time.Time
	TimeModified time.Time
	CreatedBy    string
	ModifiedBy   string
}

type Post struct {
	AuditableContent
	Caption      string
	MessageBody  string
	URL          string
	ImageURI     string
	ThumbnailURI string
	Keywords     []string
	Linkers      []string
	AuthorMood   MoodState
}

var Moods map[string]MoodState

func NewPost(username string, mood MoodState, caption string, messageBody string,
	url string, imageURI string, thumbnailURI string, keywords []string) *Post {
	auditableContent := AuditableContent{CreatedBy: username, TimeCreated: time.Now()}
	return &Post{AuditableContent: auditableContent, Caption: caption, MessageBody: messageBody,
		URL: url, ImageURI: imageURI, ThumbnailURI: thumbnailURI, Keywords: keywords}
}

func init() {
	Moods = map[string]MoodState{"neutral": MoodStateNeutral, "happy": MoodStateHappy, "sad": MoodStateSad, "angry": MoodStateAngry}
}

func main() {
	fmt.Printf("Contents of Moods %v", Moods)
}
