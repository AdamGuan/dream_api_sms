package models

type MResp struct {
	responseNo  int
	responseMsg string
}

type MFindPwdResp struct {
	responseNo  int
	responseMsg string
	password string
}

type MUserExistsResp struct {
	responseNo  int
	responseMsg string
	exists string
}