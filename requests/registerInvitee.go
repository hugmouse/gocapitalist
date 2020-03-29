package requests

type RegisterInvitee struct {
	InviteeLogin    string
	InviteeEmail    string
	InviteeNickname string
	InviteeMobile   string
}

func (r *RegisterInvitee) Params() (map[string]string, map[string]interface{}) {
	params := map[string]string{}
	logParams := map[string]interface{}{}
	if r == nil {
		return params, logParams
	}

	logParams["operation"] = "register_invitee"
	params["operation"] = "register_invitee"

	if r.InviteeLogin != "" {
		logParams["invitee_login"] = r.InviteeLogin
		params["invitee_login"] = r.InviteeLogin
	}

	if r.InviteeEmail != "" {
		logParams["invitee_email"] = r.InviteeEmail
		params["invitee_email"] = r.InviteeEmail
	}

	if r.InviteeNickname != "" {
		logParams["invitee_nickname"] = r.InviteeNickname
		params["invitee_nickname"] = r.InviteeNickname
	}

	if r.InviteeMobile != "" {
		logParams["invitee_mobile"] = r.InviteeMobile
		params["invitee_mobile"] = r.InviteeMobile
	}

	return params, logParams
}
