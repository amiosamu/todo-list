package entity

type Task struct {
	ID       string `bson:"_id,omitempty" json:"id,omitempty"`
	Title    string `bson:"title,omitempty" json:"title,omitempty"`
	ActiveAt string `bson:"activeAt,omitempty" json:"activeAt,omitempty"`
	Status   string `bson:"status,omitempty" json:"status,omitempty"`
}

type UpdateTask struct {
	Title    string `bson:"title,omitempty" json:"title,omitempty"`
	ActiveAt string `bson:"activeAt,omitempty" json:"activeAt,omitempty"`
	Status   string `bson:"status,omitempty" json:"status,omitempty"`
}
