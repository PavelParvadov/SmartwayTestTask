package services

const (
	OpSave            = "save"
	OpUpdate          = "update"
	OpDelete          = "delete"
	OpGetByCompany    = "get_by_company"
	OpGetByDepartment = "get_by_department"

	LogFieldOperation = "op"

	LogMsgMappedError           = "mapped repository error"
	LogMsgSaveFailed            = "save employee failed"
	LogMsgUpdateFailed          = "update employee failed"
	LogMsgDeleteFailed          = "delete employee failed"
	LogMsgGetByCompanyFailed    = "get employees by company failed"
	LogMsgGetByDepartmentFailed = "get employees by department failed"
	LogMsgInvalidInput          = "invalid input"
)
