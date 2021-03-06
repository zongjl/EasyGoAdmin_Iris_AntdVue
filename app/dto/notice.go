// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ 赋能开发者，助力企业发展 ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2022 深圳EasyGoAdmin研发中心
// +----------------------------------------------------------------------
// | Licensed LGPL-3.0 EasyGoAdmin并不是自由软件，未经许可禁止去掉相关版权
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: @半城风雨 团队荣誉出品
// +----------------------------------------------------------------------
// | 版权和免责声明:
// | 本团队对该软件框架产品拥有知识产权（包括但不限于商标权、专利权、著作权、商业秘密等）
// | 均受到相关法律法规的保护，任何个人、组织和单位不得在未经本团队书面授权的情况下对所授权
// | 软件框架产品本身申请相关的知识产权，禁止用于任何违法、侵害他人合法权益等恶意的行为，禁
// | 止用于任何违反我国法律法规的一切项目研发，任何个人、组织和单位用于项目研发而产生的任何
// | 意外、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、
// | 附带或衍生的损失等)，本团队不承担任何法律责任，本软件框架禁止任何单位和个人、组织用于
// | 任何违法、侵害他人合法利益等恶意的行为，如有发现违规、违法的犯罪行为，本团队将无条件配
// | 合公安机关调查取证同时保留一切以法律手段起诉的权利，本软件框架只能用于公司和个人内部的
// | 法律所允许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
// +----------------------------------------------------------------------

/**
 * 通知公告Dto
 * @author 半城风雨
 * @since 2021/11/13
 * @File : notice
 */
package dto

import "github.com/gookit/validate"

// 分页查询
type NoticePageReq struct {
	Page   int    `form:"page"`   // 页码
	Limit  int    `form:"limit"`  // 每页数
	Title  string `form:"title"`  // 通知标题
	Source int    `form:"source"` // 通知来源
}

// 添加通知公告
type NoticeAddReq struct {
	Id      int    `form:"id"`
	Title   string `form:"title" validate:"required"`   // 通知标题
	Content string `form:"content" validate:"required"` // 通知内容
	Source  int    `form:"source" validate:"int"`       // 来源：1内部通知 2外部新闻
	IsTop   int    `form:"isTop" validate:"int"`        // 是否置顶：1是 2否
	Status  int    `form:"status" validate:"int"`       // 状态：1已发布 2待发布
}

// 添加通知表单验证
func (v NoticeAddReq) Messages() map[string]string {
	return validate.MS{
		"Title.required":   "通知标题不能为空.",
		"Content.required": "通知内容不能为空.",
		"Source.int":       "请选择通知来源.",
		"IsTop.int":        "请选择是否置顶.",
		"Status.int":       "请选择通知状态.",
	}
}

// 更新通知公告
type NoticeUpdateReq struct {
	Id      int    `form:"id" validate:"int"`
	Title   string `form:"title" validate:"required"`   // 通知标题
	Content string `form:"content" validate:"required"` // 通知内容
	Source  int    `form:"source" validate:"int"`       // 来源：1内部通知 2外部新闻
	IsTop   int    `form:"isTop" validate:"int"`        // 是否置顶：1是 2否
	Status  int    `form:"status" validate:"int"`       // 状态：1已发布 2待发布
}

// 更新通知表单验证
func (v NoticeUpdateReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":           "通知ID不能为空.",
		"Title.required":   "通知标题不能为空.",
		"Content.required": "通知内容不能为空.",
		"Source.int":       "请选择通知来源.",
		"IsTop.int":        "请选择是否置顶.",
		"Status.int":       "请选择通知状态.",
	}
}

// 设置状态
type NoticeStatusReq struct {
	Id     int `form:"id" validate:"int"`
	Status int `form:"status" validate:"int"`
}

// 设置状态参数验证
func (v NoticeStatusReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":     "通知ID不能为空.",
		"Status.int": "请选择通知状态.",
	}
}
