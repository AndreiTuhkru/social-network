package model

import "time"

// Data structures and domain model

var UserID int

type User struct {
	Id 			int
	Username 	string
	Email 		string
	Password 	string
	FirstName 	string
	LastName 	string
	DOB 		string
	AvatarURL 	string
	About 		string
	CreatedAt 	string
	UpdatedAt 	string
}

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegistrationData struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	DOB       string `json:"dob"`
	AvatarURL string `json:"avatar_url,omitempty"`
	About     string `json:"about,omitempty"`
}

type AuthResponse struct {
	IsAuthenticated bool `json:"is_authenticated"`
}

type Session struct {
	Id           int       `json:"id"`
	SessionToken string    `json:"session_token"`
	UserID       int       `json:"user_id"`
	ExpiresAt    time.Time `json:"expires_at"`
}

type Post struct {
	Id             int       `json:"id"`
	UserID         int       `json:"user_id"`
	Title          string    `json:"title"`
	Content        string    `json:"content,omitempty"`
	ImageURL       string    `json:"image_url,omitempty"`
	PrivacySetting string    `json:"privacy_setting"`
	CreatedAt      time.Time `json:"created_at"`
}

type CreatePostRequest struct {
	Title          string `json:"title"`
	Content        string `json:"content,omitempty"`
	ImageURL       string `json:"image_url,omitempty"`
	PrivacySetting string `json:"privacy_setting"`
}

type UpdatePostRequest struct {
	Id             int    `json:"id"`
	Title          string `json:"title"`
	Content        string `json:"content,omitempty"`
	ImageURL       string `json:"image_url,omitempty"`
	PrivacySetting string `json:"privacy_setting"`
}

type Comment struct {
	Id        int       `json:"id"`
	PostID    int       `json:"post_id"`
	UserID    int       `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateCommentRequest struct {
	Id        int       `json:"id"`
	PostID    int       `json:"post_id"`
	UserID    int       `json:"user_id"`
	Content   string    `json:"content"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Group struct {
	Id          int       `json:"id"`
	CreatorId   int       `json:"creator_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type GroupMember struct {
	GroupId  int       `json:"group_id"`
	UserId   int       `json:"user_id"`
	JoinedAt time.Time `json:"joined_at"`
}

type Friend struct {
	Id           int       `json:"id"`
	UserId1      int       `json:"user_id_1"`
	UserId2      int       `json:"user_id_2"`
	Status       string    `json:"status"`
	ActionUserId int       `json:"action_user_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Notification struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Type      string    `json:"type"`
	Message   string    `json:"message"`
	IsRead    bool      `json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
}

type GroupInvitation struct {
	Id           int       `json:"id"`
	GroupId      int       `json:"group_id"`
	JoinUserId   int       `json:"join_user_id"`
	InviteUserId int       `json:"invite_user_id,omitempty"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
}

type Event struct {
	Id          int       `json:"id"`
	CreatorId   int       `json:"creator_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	CreatedAt   time.Time `json:"created_at"`
}
