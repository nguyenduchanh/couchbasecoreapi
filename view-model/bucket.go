package view_model

type BucketCreateCommand struct {
	Name                                     string `json:"name"` // new bucket name
	RamQuotaMB                               string `json:"ramQuotaMB"`
	BucketType                               string `json:"bucketType"`
	EvictionPolicy                           string `json:"evictionPolicy"`
	DurabilityMinLevel                       string `json:"durabilityMinLevel"`
	ThreadsNumber                            string `json:"threadsNumber"`
	ReplicaNumber                            string `json:"replicaNumber"`
	ReplicaIndex                             string `json:"replicaIndex"`
	ConflictResolutionType                   string `json:"conflictResolutionType"`
	FlushEnabled                             string `json:"flushEnabled"`
	AutoCompactionDefined                    string `json:"autoCompactionDefined"`
	ParallelDBAndViewCompaction              string `json:"parallelDBAndViewCompaction"`
	DatabaseFragmentationThresholdPercentage string `json:"databaseFragmentationThresholdPercentage"`
	DatabaseFragmentationThresholdSize       string `json:"databaseFragmentationThresholdSize"`
	ViewFragmentationThresholdPercentage     string `json:"viewFragmentationThresholdPercentage"`
	ViewFragmentationThresholdSize           string `json:"viewFragmentationThresholdSize"`
	IndexCompactionMode                      string `json:"indexCompactionMode"`
	PurgeInterval                            string `json:"purgeInterval"`
	AllowedTimePeriodFromHour                string `json:"allowedTimePeriodFromHour"`
	AllowedTimePeriodFromMinute              string `json:"allowedTimePeriodFromMinute"`
	AllowedTimePeriodToHour                  string `json:"allowedTimePeriodToHour"`
	AllowedTimePeriodToMinute                string `json:"allowedTimePeriodToMinute"`
	AllowedTimePeriodAbortOutside            string `json:"allowedTimePeriodAbortOutside"`
	ConnectionString                         string `json:"connectionstring" binding:"required"`
	UserName                                 string `json:"username" binding:"required"`
	Password                                 string `json:"password" binding:"required"`
	ClusterName                              string `json:"clustername" binding:"required"`
}
type BucketSearchCommand struct {
	ConnectionString string `json:"connectionstring" binding:"required"`
	UserName         string `json:"username" binding:"required"`
	Password         string `json:"password" binding:"required"`
	ClusterName      string `json:"clustername" binding:"required"`
}
