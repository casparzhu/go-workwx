// Code generated by sdkcodegen; DO NOT EDIT.

package workwx

// execGetAccessToken 获取access_token
func (c *WorkwxApp) execGetAccessToken(req reqAccessToken) (respAccessToken, error) {
	var resp respAccessToken
	err := c.executeQyapiGet("/cgi-bin/gettoken", req, &resp, false)
	if err != nil {
		return respAccessToken{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respAccessToken{}, bizErr
	}

	return resp, nil
}

// execGetJSAPITicket 获取企业的jsapi_ticket
func (c *WorkwxApp) execGetJSAPITicket(req reqJSAPITicket) (respJSAPITicket, error) {
	var resp respJSAPITicket
	err := c.executeQyapiGet("/cgi-bin/get_jsapi_ticket", req, &resp, true)
	if err != nil {
		return respJSAPITicket{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respJSAPITicket{}, bizErr
	}

	return resp, nil
}

// execGetJSAPITicketAgentConfig 获取应用的jsapi_ticket
func (c *WorkwxApp) execGetJSAPITicketAgentConfig(req reqJSAPITicketAgentConfig) (respJSAPITicket, error) {
	var resp respJSAPITicket
	err := c.executeQyapiGet("/cgi-bin/ticket/get", req, &resp, true)
	if err != nil {
		return respJSAPITicket{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respJSAPITicket{}, bizErr
	}

	return resp, nil
}

// execJSCode2Session 临时登录凭证校验code2Session
func (c *WorkwxApp) execJSCode2Session(req reqJSCode2Session) (respJSCode2Session, error) {
	var resp respJSCode2Session
	err := c.executeQyapiGet("/cgi-bin/miniprogram/jscode2session", req, &resp, true)
	if err != nil {
		return respJSCode2Session{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respJSCode2Session{}, bizErr
	}

	return resp, nil
}

// execUserGet 读取成员
func (c *WorkwxApp) execUserGet(req reqUserGet) (respUserGet, error) {
	var resp respUserGet
	err := c.executeQyapiGet("/cgi-bin/user/get", req, &resp, true)
	if err != nil {
		return respUserGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respUserGet{}, bizErr
	}

	return resp, nil
}

// execUserList 获取部门成员详情
func (c *WorkwxApp) execUserList(req reqUserList) (respUserList, error) {
	var resp respUserList
	err := c.executeQyapiGet("/cgi-bin/user/list", req, &resp, true)
	if err != nil {
		return respUserList{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respUserList{}, bizErr
	}

	return resp, nil
}

// execUserIDByMobile 手机号获取userid
func (c *WorkwxApp) execUserIDByMobile(req reqUserIDByMobile) (respUserIDByMobile, error) {
	var resp respUserIDByMobile
	err := c.executeQyapiJSONPost("/cgi-bin/user/getuserid", req, &resp, true)
	if err != nil {
		return respUserIDByMobile{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respUserIDByMobile{}, bizErr
	}

	return resp, nil
}

// execDeptList 获取部门列表
func (c *WorkwxApp) execDeptList(req reqDeptList) (respDeptList, error) {
	var resp respDeptList
	err := c.executeQyapiGet("/cgi-bin/department/list", req, &resp, true)
	if err != nil {
		return respDeptList{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respDeptList{}, bizErr
	}

	return resp, nil
}

// execUserInfoGet 获取访问用户身份
func (c *WorkwxApp) execUserInfoGet(req reqUserInfoGet) (respUserInfoGet, error) {
	var resp respUserInfoGet
	err := c.executeQyapiGet("/cgi-bin/user/getuserinfo", req, &resp, true)
	if err != nil {
		return respUserInfoGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respUserInfoGet{}, bizErr
	}

	return resp, nil
}

// execExternalContactList 获取客户列表
func (c *WorkwxApp) execExternalContactList(req reqExternalContactList) (respExternalContactList, error) {
	var resp respExternalContactList
	err := c.executeQyapiGet("/cgi-bin/externalcontact/list", req, &resp, true)
	if err != nil {
		return respExternalContactList{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respExternalContactList{}, bizErr
	}

	return resp, nil
}

// execExternalContactGet 获取客户详情
func (c *WorkwxApp) execExternalContactGet(req reqExternalContactGet) (respExternalContactGet, error) {
	var resp respExternalContactGet
	err := c.executeQyapiGet("/cgi-bin/externalcontact/get", req, &resp, true)
	if err != nil {
		return respExternalContactGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respExternalContactGet{}, bizErr
	}

	return resp, nil
}

// execExternalContactBatchList 批量获取客户详情
func (c *WorkwxApp) execExternalContactBatchList(req reqExternalContactBatchList) (respExternalContactBatchList, error) {
	var resp respExternalContactBatchList
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/batch/get_by_user", req, &resp, true)
	if err != nil {
		return respExternalContactBatchList{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respExternalContactBatchList{}, bizErr
	}

	return resp, nil
}

// execExternalContactRemark 修改客户备注信息
func (c *WorkwxApp) execExternalContactRemark(req reqExternalContactRemark) (respExternalContactRemark, error) {
	var resp respExternalContactRemark
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/remark", req, &resp, true)
	if err != nil {
		return respExternalContactRemark{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respExternalContactRemark{}, bizErr
	}

	return resp, nil
}

// execExternalContactListCorpTags 获取企业标签库
func (c *WorkwxApp) execExternalContactListCorpTags(req reqExternalContactListCorpTags) (respExternalContactListCorpTags, error) {
	var resp respExternalContactListCorpTags
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/get_corp_tag_list", req, &resp, true)
	if err != nil {
		return respExternalContactListCorpTags{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respExternalContactListCorpTags{}, bizErr
	}

	return resp, nil
}

// execExternalContactAddCorpTag 添加企业客户标签
func (c *WorkwxApp) execExternalContactAddCorpTag(req reqExternalContactAddCorpTag) (respExternalContactAddCorpTag, error) {
	var resp respExternalContactAddCorpTag
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/add_corp_tag", req, &resp, true)
	if err != nil {
		return respExternalContactAddCorpTag{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respExternalContactAddCorpTag{}, bizErr
	}

	return resp, nil
}

// execExternalContactEditCorpTag 编辑企业客户标签
func (c *WorkwxApp) execExternalContactEditCorpTag(req reqExternalContactEditCorpTag) (respExternalContactEditCorpTag, error) {
	var resp respExternalContactEditCorpTag
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/edit_corp_tag", req, &resp, true)
	if err != nil {
		return respExternalContactEditCorpTag{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respExternalContactEditCorpTag{}, bizErr
	}

	return resp, nil
}

// execExternalContactDelCorpTag 删除企业客户标签
func (c *WorkwxApp) execExternalContactDelCorpTag(req reqExternalContactDelCorpTag) (respExternalContactDelCorpTag, error) {
	var resp respExternalContactDelCorpTag
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/del_corp_tag", req, &resp, true)
	if err != nil {
		return respExternalContactDelCorpTag{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respExternalContactDelCorpTag{}, bizErr
	}

	return resp, nil
}

// execExternalContactMarkTag 标记客户企业标签
func (c *WorkwxApp) execExternalContactMarkTag(req reqExternalContactMarkTag) (respExternalContactMarkTag, error) {
	var resp respExternalContactMarkTag
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/mark_tag", req, &resp, true)
	if err != nil {
		return respExternalContactMarkTag{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respExternalContactMarkTag{}, bizErr
	}

	return resp, nil
}

// execListUnassignedExternalContact 获取离职成员的客户列表
func (c *WorkwxApp) execListUnassignedExternalContact(req reqListUnassignedExternalContact) (respListUnassignedExternalContact, error) {
	var resp respListUnassignedExternalContact
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/get_unassigned_list", req, &resp, true)
	if err != nil {
		return respListUnassignedExternalContact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respListUnassignedExternalContact{}, bizErr
	}

	return resp, nil
}

// execTransferExternalContact 分配成员的客户
func (c *WorkwxApp) execTransferExternalContact(req reqTransferExternalContact) (respTransferExternalContact, error) {
	var resp respTransferExternalContact
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/transfer", req, &resp, true)
	if err != nil {
		return respTransferExternalContact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respTransferExternalContact{}, bizErr
	}

	return resp, nil
}

// execGetTransferExternalContactResult 查询客户接替结果
func (c *WorkwxApp) execGetTransferExternalContactResult(req reqGetTransferExternalContactResult) (respGetTransferExternalContactResult, error) {
	var resp respGetTransferExternalContactResult
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/get_transfer_result", req, &resp, true)
	if err != nil {
		return respGetTransferExternalContactResult{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respGetTransferExternalContactResult{}, bizErr
	}

	return resp, nil
}

// execTransferGroupChatExternalContact 离职成员的群再分配
func (c *WorkwxApp) execTransferGroupChatExternalContact(req reqTransferGroupChatExternalContact) (respTransferGroupChatExternalContact, error) {
	var resp respTransferGroupChatExternalContact
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/groupchat/transfer", req, &resp, true)
	if err != nil {
		return respTransferGroupChatExternalContact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respTransferGroupChatExternalContact{}, bizErr
	}

	return resp, nil
}

// execExternalContractGroupChatGet 获取客户群详情
func (c *WorkwxApp) execExternalContractGroupChatGet(req reqGroupChatExternalContact) (respGetExternalContractGroupChatResult, error) {
	var resp respGetExternalContractGroupChatResult
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/groupchat/get", req, &resp, true)
	if err != nil {
		return respGetExternalContractGroupChatResult{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respGetExternalContractGroupChatResult{}, bizErr
	}

	return resp, nil
}

// execAppchatCreate 创建群聊会话
func (c *WorkwxApp) execAppchatCreate(req reqAppchatCreate) (respAppchatCreate, error) {
	var resp respAppchatCreate
	err := c.executeQyapiJSONPost("/cgi-bin/appchat/create", req, &resp, true)
	if err != nil {
		return respAppchatCreate{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respAppchatCreate{}, bizErr
	}

	return resp, nil
}

// execAppchatGet 获取群聊会话
func (c *WorkwxApp) execAppchatGet(req reqAppchatGet) (respAppchatGet, error) {
	var resp respAppchatGet
	err := c.executeQyapiGet("/cgi-bin/appchat/get", req, &resp, true)
	if err != nil {
		return respAppchatGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respAppchatGet{}, bizErr
	}

	return resp, nil
}

// execMessageSend 发送应用消息
func (c *WorkwxApp) execMessageSend(req reqMessage) (respMessageSend, error) {
	var resp respMessageSend
	err := c.executeQyapiJSONPost("/cgi-bin/message/send", req, &resp, true)
	if err != nil {
		return respMessageSend{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respMessageSend{}, bizErr
	}

	return resp, nil
}

// execAppchatSend 应用推送消息
func (c *WorkwxApp) execAppchatSend(req reqMessage) (respMessageSend, error) {
	var resp respMessageSend
	err := c.executeQyapiJSONPost("/cgi-bin/appchat/send", req, &resp, true)
	if err != nil {
		return respMessageSend{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respMessageSend{}, bizErr
	}

	return resp, nil
}

// execMediaUpload 上传临时素材
func (c *WorkwxApp) execMediaUpload(req reqMediaUpload) (respMediaUpload, error) {
	var resp respMediaUpload
	err := c.executeQyapiMediaUpload("/cgi-bin/media/upload", req, &resp, true)
	if err != nil {
		return respMediaUpload{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respMediaUpload{}, bizErr
	}

	return resp, nil
}

// execMediaUploadImg 上传永久图片
func (c *WorkwxApp) execMediaUploadImg(req reqMediaUploadImg) (respMediaUploadImg, error) {
	var resp respMediaUploadImg
	err := c.executeQyapiMediaUpload("/cgi-bin/media/uploadimg", req, &resp, true)
	if err != nil {
		return respMediaUploadImg{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respMediaUploadImg{}, bizErr
	}

	return resp, nil
}

// execOAGetTemplateDetail 获取审批模板详情
func (c *WorkwxApp) execOAGetTemplateDetail(req reqOAGetTemplateDetail) (respOAGetTemplateDetail, error) {
	var resp respOAGetTemplateDetail
	err := c.executeQyapiJSONPost("/cgi-bin/oa/gettemplatedetail", req, &resp, true)
	if err != nil {
		return respOAGetTemplateDetail{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respOAGetTemplateDetail{}, bizErr
	}

	return resp, nil
}

// execOAApplyEvent 提交审批申请
func (c *WorkwxApp) execOAApplyEvent(req reqOAApplyEvent) (respOAApplyEvent, error) {
	var resp respOAApplyEvent
	err := c.executeQyapiJSONPost("/cgi-bin/oa/applyevent", req, &resp, true)
	if err != nil {
		return respOAApplyEvent{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respOAApplyEvent{}, bizErr
	}

	return resp, nil
}

// execOAGetApprovalInfo 批量获取审批单号
func (c *WorkwxApp) execOAGetApprovalInfo(req reqOAGetApprovalInfo) (respOAGetApprovalInfo, error) {
	var resp respOAGetApprovalInfo
	err := c.executeQyapiJSONPost("/cgi-bin/oa/getapprovalinfo", req, &resp, true)
	if err != nil {
		return respOAGetApprovalInfo{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respOAGetApprovalInfo{}, bizErr
	}

	return resp, nil
}

// execOAGetApprovalDetail 获取审批申请详情
func (c *WorkwxApp) execOAGetApprovalDetail(req reqOAGetApprovalDetail) (respOAGetApprovalDetail, error) {
	var resp respOAGetApprovalDetail
	err := c.executeQyapiJSONPost("/cgi-bin/oa/getapprovaldetail", req, &resp, true)
	if err != nil {
		return respOAGetApprovalDetail{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respOAGetApprovalDetail{}, bizErr
	}

	return resp, nil
}

// execMsgAuditListPermitUser 获取会话内容存档开启成员列表
func (c *WorkwxApp) execMsgAuditListPermitUser(req reqMsgAuditListPermitUser) (respMsgAuditListPermitUser, error) {
	var resp respMsgAuditListPermitUser
	err := c.executeQyapiJSONPost("/cgi-bin/msgaudit/get_permit_user_list", req, &resp, true)
	if err != nil {
		return respMsgAuditListPermitUser{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respMsgAuditListPermitUser{}, bizErr
	}

	return resp, nil
}

// execMsgAuditCheckSingleAgree 获取会话同意情况（单聊）
func (c *WorkwxApp) execMsgAuditCheckSingleAgree(req reqMsgAuditCheckSingleAgree) (respMsgAuditCheckSingleAgree, error) {
	var resp respMsgAuditCheckSingleAgree
	err := c.executeQyapiJSONPost("/cgi-bin/msgaudit/check_single_agree", req, &resp, true)
	if err != nil {
		return respMsgAuditCheckSingleAgree{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respMsgAuditCheckSingleAgree{}, bizErr
	}

	return resp, nil
}

// execMsgAuditCheckRoomAgree 获取会话同意情况（群聊）
func (c *WorkwxApp) execMsgAuditCheckRoomAgree(req reqMsgAuditCheckRoomAgree) (respMsgAuditCheckRoomAgree, error) {
	var resp respMsgAuditCheckRoomAgree
	err := c.executeQyapiJSONPost("/cgi-bin/msgaudit/check_room_agree", req, &resp, true)
	if err != nil {
		return respMsgAuditCheckRoomAgree{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respMsgAuditCheckRoomAgree{}, bizErr
	}

	return resp, nil
}

// execMsgAuditGetGroupChat 获取会话内容存档内部群信息
func (c *WorkwxApp) execMsgAuditGetGroupChat(req reqMsgAuditGetGroupChat) (respMsgAuditGetGroupChat, error) {
	var resp respMsgAuditGetGroupChat
	err := c.executeQyapiJSONPost("/cgi-bin/msgaudit/groupchat/get", req, &resp, true)
	if err != nil {
		return respMsgAuditGetGroupChat{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respMsgAuditGetGroupChat{}, bizErr
	}

	return resp, nil
}
