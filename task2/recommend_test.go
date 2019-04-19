package task2

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func testRecommendCreateGroup(t *testing.T, r *Recommend) {
	tests := []*CreateGroupRequest{
		{
			GroupName: "答案组1",
			GroupAttrs: []struct {
				AttrName  string
				AttrValue string
			}{
				{"性别", "男"},
				{"城市", "北京"},
			},
			Answers: []struct {
				KId    uint32
				AId    uint32
				Answer string
			}{
				{1, 1, "御姐风格"},
				{2, 2, "御姐风格"},
			},
		},
		{
			GroupName: "答案组2",
			GroupAttrs: []struct {
				AttrName  string
				AttrValue string
			}{
				{"性别", "女"},
				{"城市", "北京"},
			},
			Answers: []struct {
				KId    uint32
				AId    uint32
				Answer string
			}{
				{1, 3, "大叔风格"},
				{2, 4, "大叔风格"},
			},
		},
		{
			GroupName: "答案组3",
			GroupAttrs: []struct {
				AttrName  string
				AttrValue string
			}{
				{"性别", "男"},
				{"城市", "上海"},
			},
			Answers: []struct {
				KId    uint32
				AId    uint32
				Answer string
			}{
				{1, 5, "萝莉风格"},
				{2, 6, "萝莉风格"},
			},
		},
		{
			GroupName: "答案组4",
			GroupAttrs: []struct {
				AttrName  string
				AttrValue string
			}{
				{"性别", "女"},
				{"城市", "上海"},
			},
			Answers: []struct {
				KId    uint32
				AId    uint32
				Answer string
			}{
				{1, 7, "正太风格"},
				{2, 8, "正太风格"},
			},
		},
		{
			GroupName: "答案组5",
			GroupAttrs: []struct {
				AttrName  string
				AttrValue string
			}{
				{"性别", "*"},
				{"城市", "广州"},
			},
			Answers: []struct {
				KId    uint32
				AId    uint32
				Answer string
			}{
				{1, 9, "普通风格"},
				{2, 10, "普通风格"},
			},
		},
	}
	for _, test := range tests {
		err := r.CreateGroup(context.Background(), test)
		assert.Nil(t, err, test)
	}
}

func testRecommendMatch1(t *testing.T, r *Recommend) {
	tests := []struct {
		request *MatchRequest
		matched bool
		AId     uint32
	}{
		{
			request: &MatchRequest{
				KId: 1,
				UId: 1,
				UserAttrs: []struct {
					AttrName  string
					AttrValue string
				}{
					{"性别", "男"},
					{"城市", "北京"},
				},
			},
			matched: true,
			AId:     1,
		},
		{
			request: &MatchRequest{
				KId: 1,
				UId: 2,
				UserAttrs: []struct {
					AttrName  string
					AttrValue string
				}{
					{"性别", "女"},
					{"城市", "上海"},
				},
			},
			matched: true,
			AId:     7,
		},
		{
			request: &MatchRequest{
				KId: 1,
				UId: 3,
				UserAttrs: []struct {
					AttrName  string
					AttrValue string
				}{
					{"性别", "男"},
					{"城市", "广州"},
				},
			},
			matched: true,
			AId:     9,
		},
	}
	for _, test := range tests {
		matched, AId, err := r.Match(context.Background(), test.request)
		assert.Nil(t, err, test.request)
		assert.Equal(t, test.matched, matched, test.request)
		assert.Equal(t, test.AId, AId, test.request)
	}
}

func testRecommendDeleteGroup(t *testing.T, r *Recommend) {
	err := r.DeleteGroup(context.Background(), "答案组5")
	assert.Nil(t, err)
}

func testRecommendMatch2(t *testing.T, r *Recommend) {
	tests := []struct {
		request *MatchRequest
		matched bool
		AId     uint32
	}{
		{
			request: &MatchRequest{
				KId: 1,
				UId: 3,
				UserAttrs: []struct {
					AttrName  string
					AttrValue string
				}{
					{"性别", "男"},
					{"城市", "广州"},
				},
			},
			matched: false,
			AId:     0,
		},
	}
	for _, test := range tests {
		matched, AId, err := r.Match(context.Background(), test.request)
		assert.Nil(t, err, test.request)
		assert.Equal(t, test.matched, matched, test.request)
		assert.Equal(t, test.AId, AId, test.request)
	}
}

func Test(t *testing.T) {
	r := NewRecommend()
	testRecommendCreateGroup(t, r)
	testRecommendMatch1(t, r)
	testRecommendDeleteGroup(t, r)
	testRecommendMatch2(t, r)
}
