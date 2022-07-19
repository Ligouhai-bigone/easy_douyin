all: video_client comment_client
# user_client:
# 	kitex -I ./idl/ -module github.com/Ligouhai-bigone/easy_douyin -type protobuf ./idl/user.proto

video_client:
	kitex -I ./idl/ -module github.com/Ligouhai-bigone/easy_douyin ./idl/video.thrift

comment_client:
	kitex -I ./idl/ -module github.com/Ligouhai-bigone/easy_douyin ./idl/comment.thrift