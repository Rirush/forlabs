package forlabs

import (
	"encoding/json"
)

type Day struct {
	Name string `json:"name"`
	Active bool `json:"active"`
}

type Position struct {
	Start string `json:"start"`
	End string `json:"end"`
}

// encapsulated into grid
type Grid struct {
	Days struct {
		Mon Day `json:"mon"`
		Tue Day	`json:"tue"`
		Wed Day `json:"wed"`
		Thu Day `json:"thu"`
		Fri Day `json:"fri"`
		Sat Day `json:"sat"`
		Sun Day `json:"sun"`
	} `json:"days"`
	Positions []Position `json:"positions"`
	LessonName string `json:"lesson_name"`
	Type int `json:"type"`
	Week int `json:"upperweek"`
}

// /lm-vendor/repositories/sched/get_grid
func (c *Client) GetGrid() (g *Grid, err error) {
	resp, err := c.Post(Endpoint+"/lm-vendor/repositories/sched/get_grid", struct{}{})
	if err != nil {
		return nil, err
	}
	wrapper := &struct {
		Grid *Grid `json:"grid"`
	}{}
	err = json.Unmarshal(resp.Body(), wrapper)
	g = wrapper.Grid
	return
}

type ScheduleMeta struct {
	Type string `json:"type"`
	IsOwn bool `json:"is_own"`
	StreamIDs []int `json:"stream_ids"`
}

type Stream struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

type ScheduleEntry struct {
	Day int `json:"day"`
	Position int `json:"position"`
	Type int `json:"type"` // 1 - practice, 2 - lecture
	RoomName string `json:"room_name"`
	LecturerName string `json:"lecturer_name"`
	StudyID int `json:"study_id"`
	StudyName string `json:"study_name"`
	Subgroup string `json:"subgroup"`
	Streams []Stream `json:"streams"`
}

type Schedule struct {
	Meta ScheduleMeta `json:"meta"`
	Streams []Stream `json:"streams"`
	Entries []ScheduleEntry `json:"entries"`
}

// /lm-vendor/repositories/sched/get_schedule
func (c *Client) GetSchedule(stream uint) (s *Schedule, err error) {
	resp, err := c.Post(Endpoint+"/lm-vendor/repositories/sched/get_schedule", struct{
		StreamID uint `json:"stream_id"`
	}{stream})
	if err != nil {
		return nil, err
	}
	s = &Schedule{}
	err = json.Unmarshal(resp.Body(), s)
	return
}

type Study struct {
	ID uint `json:"id"`
	VerboseName string `json:"verbose_name"`
}

func (c *Client) GetStudies(stream uint) (st []Study, err error) {
	resp, err := c.Post(Endpoint+"/lm-vendor/repositories/learning/get_studies", struct{
		StreamID uint `json:"stream_id"`
	}{stream})
	if err != nil {
		return nil, err
	}
	wrap := &struct {
		Studies []Study `json:"studies"`
	}{}
	err = json.Unmarshal(resp.Body(), wrap)
	st = wrap.Studies
	return
}

// encapsulated into posts
type Post struct {
	ID int `json:"id"`
	StudyID int `json:"study_id"`
	Title string `json:"title"`
	Content string `json:"content"`
	// 2020-10-26T00:48:32.000000Z
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
	Attachments string `json:"attachments"`
	User struct {
		Name string `json:"name"`
	} `json:"user"`
}

// /lm-vendor/repositories/learning/get_posts
func (c *Client) GetPosts() {

}

type Assignment struct {
	ID int `json:"id"`
	TaskID int `json:"task_id"`
	Status int `json:"status"`
	//Choice
	//Variant
	//Options
	LastRepliedAt string `json:"last_replied_at"`
	AssignmentCredits int `json:"assignment_credits"`
	AssignmentDate string `json:"assignment_date"`
	AssessmentLecturerID int `json:"assessment_lecturer_id"`
	ResponsesCount int `json:"responses_count"`
}

type File struct {
	ID int `json:"id"`
	Type string `json:"type"`
	Status int `json:"status"`
	Disk string `json:"disk"`
	Directory *string `json:"directory"`
	UUID string `json:"uuid"`
	Filename string `json:"filename"`
	Description *string `json:"description"`
	MimeType string `json:"mime_type"`
	Size int `json:"size"`
	Width int `json:"width"`
	Height int `json:"height"`
	Sort int `json:"sort"`
	CreatedAt string `json:"created_at"`
	URL string `json:"url"`
	//Preview
	//Thumbnail
	HumanSize string `json:"human_size"`
}

type Task struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Content string `json:"content"`
	Type int `json:"type"`
	CourseID *int `json:"course_id"`
	ChapterID *int `json:"chapter_id"`
	ChapterTitle *int `json:"chapter_title"`
	ChapterHasContent bool `json:"chapter_has_content"`
	Level int `json:"level"`
	Files []File `json:"files"`
	PivotStatus int `json:"pivot_status"`
	PivotCost int `json:"pivot_cost"`
	PivotSort int `json:"pivot_sort"`
	PivotType int `json:"pivot_type"`
	PivotStartAt *string `json:"pivot_start_at"`
	PivotEndAt *string `json:"pivot_end_at"`
	PivotDescription *string `json:"pivot_description"`
}

type Tasks struct {
	Assignments []Assignment `json:"assignments"`
	Tasks []Task `json:"tasks"`
}

// /lm-vendor/repositories/learning/get_tasks
func (c *Client) GetTasks() {

}