package task2

import "context"

type CreateGroupRequest struct {
	// 答案组名
	GroupName string
	// 答案组属性
	GroupAttrs []struct {
		// 属性名
		AttrName string
		// 属性值
		AttrValue string
	}
	Answers []struct {
		// 知识点id
		KId uint32
		// 答案id
		AId uint32
		// 答案内容
		Answer string
	}
}

type MatchRequest struct {
	// 知识点id
	KId uint32
	// 用户id
	UId uint32
	// 用户属性
	UserAttrs []struct {
		AttrName  string
		AttrValue string
	}
}

type Recommend struct {
}

func NewRecommend() *Recommend {
	return &Recommend{}
}

// todo 创建答案组
func (r *Recommend) CreateGroup(ctx context.Context, request *CreateGroupRequest) (err error) {
	return
}

// todo 删除答案组
func (r *Recommend) DeleteGroup(ctx context.Context, GroupName string) (err error) {
	return
}

// todo 匹配答案
func (r *Recommend) Match(ctx context.Context, request *MatchRequest) (matched bool, AId uint32, err error) {
	return
}
