package view_model

type ClusterSearchCommand struct {
	ConnectionString string `json:"connectionstring" binding:"required"`
	UserName         string `json:"username" binding:"required"`
	Password         string `json:"password" binding:"required"`
}
