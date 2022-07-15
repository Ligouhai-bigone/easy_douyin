package main

import (
	"context"
	"github.com/Ligouhai-bigone/easy_douyin/cmd/comment/kitex_gen/commentdemo"
)

// NoteServiceImpl implements the last service interface defined in the IDL.
type NoteServiceImpl struct{}

// SendComment implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) SendComment(ctx context.Context, req *commentdemo.SendCommentRequest) (resp *commentdemo.SendCommentResponse, err error) {
	// TODO: Your code here...
	return
}

// GetAllComments implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) GetAllComments(ctx context.Context, req *commentdemo.GetAllCommentsRequest) (resp *commentdemo.GetAllCommentsResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteComment implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) DeleteComment(ctx context.Context, req *commentdemo.DeleteCommentRequset) (resp *commentdemo.DeleteCommentResponse, err error) {
	// TODO: Your code here...
	return
}
