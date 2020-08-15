package view_model

type BucketCreateCommand struct {
	Name                                     string `json:"name"`
	BucketType                               string `json:"bucketType"`     //bucketType=[ couchbase | ephemeral | memcached ]
	RamQuotaMB                               string `json:"ramQuotaMB"`     //ramQuotaMB=<integer>
	EvictionPolicy                           string `json:"evictionPolicy"` //ramQuotaMB=<integer>
	DurabilityMinLevel                       string `json:"durabilityMinLevel"`
	ThreadsNumber                            string `json:"threadsNumber"`                              //threadsNumber = [3|8]
	ReplicaNumber                            string `json:"replicaNumber"`                              //replicaNumber = [1|2|3]
	CompressionMode                          string `json:"compressionMode"`                            //compressionMode=[ off | passive | active ]
	MaxTTL                                   string `json:"maxTTL"`                                     //maxTTL=<integer>
	ReplicaIndex                             string `json:"replicaIndex"`                               //replicaIndex=[ 0 | 1 ]
	ConflictResolutionType                   string `json:"conflictResolutionType"`                     //conflictResolutionType=[ seqno | lww ]
	FlushEnabled                             string `json:"flushEnabled"`                               //flushEnabled=[ 0 | 1 ]
	AutoCompactionDefined                    string `json:"autoCompactionDefined"`                      //autoCompactionDefined=[ true | false ]
	ParallelDBAndViewCompaction              string `json:"parallelDBAndViewCompaction"`                //parallelDBAndViewCompaction=[ true | false ]
	DatabaseFragmentationThresholdPercentage string `json:"databaseFragmentationThreshold[percentage]"` //databaseFragmentationThreshold[percentage]=<integer>
	DatabaseFragmentationThresholdSize       string `json:"databaseFragmentationThreshold[size]"`       //databaseFragmentationThreshold[size]=<integer>
	ViewFragmentationThresholdPercentage     string `json:"viewFragmentationThreshold[percentage]"`     //viewFragmentationThreshold[percentage]=<integer>
	ViewFragmentationThresholdSize           string `json:"viewFragmentationThreshold[size]"`           //viewFragmentationThreshold[size]=<integer>
	IndexCompactionMode                      string `json:"indexCompactionMode"`                        //indexCompactionMode=[ circular | append ]
	PurgeInterval                            string `json:"purgeInterval"`                              //purgeInterval=[ <float> | <integer> ]
	AllowedTimePeriodFromHour                string `json:"allowedTimePeriod[fromHour]"`                //allowedTimePeriod[fromHour]=<integer>
	AllowedTimePeriodFromMinute              string `json:"allowedTimePeriod[fromMinute]"`              //allowedTimePeriod[fromMinute]=<integer>
	AllowedTimePeriodToHour                  string `json:"allowedTimePeriod[toHour]"`                  //allowedTimePeriod[toHour]=<integer>
	AllowedTimePeriodToMinute                string `json:"allowedTimePeriod[toMinute]"`                //allowedTimePeriod[toMinute]=<integer>
	AllowedTimePeriodAbortOutside            string `json:"allowedTimePeriod[abortOutside]"`            //allowedTimePeriod[abortOutside]=[ true | false ]

}
