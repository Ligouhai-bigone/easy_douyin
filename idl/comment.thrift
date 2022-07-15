// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

//命名空间
namespace go commentdemo

struct BaseResp {
    //基本回复
    1:i64 status_code
    2:string status_message
    3:i64 service_time
}

struct User {
    1:int64 user_id;
    2:string user_name;
    3:int64 follow_count;
    4:int64 follower_count;
    5:bool is_follow;
}

struct Comment {
    1:i64 comment_id
    2:User user
    3:string content
    4:i64 create_time
}

struct SendCommentRequest {
    1:string token
    2:int64 video_id
    3:string comment_text

}

struct SendCommentResponse {
    1:BaseResp base_resp
    2:Comment comment
}



struct DeleteCommentRequset {
    1:string token
    2:int64 video_id
    3:int64 comment_id
}

struct DeleteCommentResponse {
    1:BaseResp base_resp
}


struct GetAllCommentsRequest {
    1:string token
    2:i64 video_id
}

struct GetAllCommentsResponse {
    1:BaseResp base_resp
    2:list<Comment> comments
}


service NoteService {
    SendCommentResponse SendComment(1:SendCommentRequest req)
    GetAllCommentsResponse GetAllComments(1:GetAllCommentsRequest req)
    DeleteCommentResponse DeleteComment(1:DeleteCommentRequset req)
}
