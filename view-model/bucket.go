package view_model

type BucketCreateCommand struct {
	Name                                     string `json:"name"` // new bucket name
	RamQuotaMB                               int    `json:"ramQuotaMB"`
	AuthType                                 string `json:"authType"` //none | sasl
	ReplicaNumber                            string `json:"replicaNumber"`
	ProxyPort                                string `json:"proxyPort"`
	BucketType                               string `json:"bucketType"`
	CompressionMode                          string `json:"compressionMode"`
	MaxTTL                                   string `json:"maxTTL"`
	FlushEnabled                             string `json:"flushEnabled"`
	ReplicaIndex                             string `json:"replicaIndex"`
	ThreadsNumber                            string `json:"threadsNumber"`
	EvictionPolicy                           string `json:"evictionPolicy"`
	DurabilityMinLevel                       string `json:"durabilityMinLevel"`
	ConflictResolutionType                   string `json:"conflictResolutionType"`
	AutoCompactionDefined                    string `json:"autoCompactionDefined"`
	ParallelDBAndViewCompaction              string `json:"parallelDBAndViewCompaction"`
	DatabaseFragmentationThresholdPercentage string `json:"databaseFragmentationThreshold[percentage]"`
	DatabaseFragmentationThresholdSize       string `json:"databaseFragmentationThreshold[size]"`
	ViewFragmentationThresholdPercentage     string `json:"viewFragmentationThreshold[percentage]"`
	ViewFragmentationThresholdSize           string `json:"viewFragmentationThreshold[size]"`
	IndexCompactionMode                      string `json:"indexCompactionMode"`
	PurgeInterval                            string `json:"purgeInterval"`
	AllowedTimePeriodFromHour                string `json:"allowedTimePeriod[fromHour]"`
	AllowedTimePeriodFromMinute              string `json:"allowedTimePeriod[fromMinute]"`
	AllowedTimePeriodToHour                  string `json:"allowedTimePeriod[toHour]"`
	AllowedTimePeriodToMinute                string `json:"allowedTimePeriod[toMinute]"`
	AllowedTimePeriodAbortOutside            string `json:"allowedTimePeriod[abortOutside]"`

	ConnectionString string `json:"connectionstring" binding:"required"`
	UserName         string `json:"username" binding:"required"`
	Password         string `json:"password" binding:"required"`
	ClusterName      string `json:"clustername" binding:"required"`
}
type BucketSearchCommand struct {
	ConnectionString string `json:"connectionstring" binding:"required"`
	UserName         string `json:"username" binding:"required"`
	Password         string `json:"password" binding:"required"`
	ClusterName      string `json:"clustername" binding:"required"`
}
