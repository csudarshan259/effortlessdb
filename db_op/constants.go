package db_op

const (
	EfdExt                = ".efd"
	UnrecognisedError     = "URE"
	Success               = "SUCCESS"
	InvalidCollectionName = "ICN"
	FileDoesNotExist      = "FNE"
	Create                = "create"
	Read                  = "read"
	Update                = "update"
	Delete                = "delete"
	Id                    = "_id"
	NewLine               = "\n"
)

type KVPair map[string]string
