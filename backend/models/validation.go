package models

type ValidationActionType string

var ValidationStatus = struct {
	CREATE ValidationActionType
	UPDATE ValidationActionType
	DELETE ValidationActionType
}{
	CREATE: "create",
	UPDATE: "update",
	DELETE: "delete",
}
