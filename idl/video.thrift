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
namespace go videodemo

struct BaseResp {
    //基本回复
    1:i64 status_code
    2:string status_message
    3:i64 service_time
}

struct Video {
    1:i64 video_id
    2:i64 user_id
    3:string play_url
    4:string cover_url
    5:i64 favorite_count
    6:i64 comment_count
    7:bool is_favorite
}

struct LikeRequest {
    1:string token
    2:i64 video_id
}

struct LikeResponse {
    1:BaseResp base_resp
}

struct UnLikeRequest {
    1:string token
    2:i64 video_id
}

struct UnLikeResponse {
    1:BaseResp base_resp
}


struct QueryVideoListRequest {
    1:string token
    2:i64 user_id
}

struct QueryVideoListResponse {
    1:BaseResp base_resp
    2:list<Video> video_list
}

struct QueryLikeVideoListRequest {
    1:string token
    2:i64 user_id
}

struct QueryLikeVideoListResponse {
    1:BaseResp base_resp
    2:list<Video> video_list
}

struct FeedVideoListRequest{
    1:i64 latest_time
    2:string token
}

struct FeedVideoListResponse{
    1:BaseResp base_resp
    2:i64 next_time
    3:list<Video> video_list
}

service VideoService {
    LikeResponse Like(1:LikeRequest req)
    UnLikeResponse UnLike(1:UnLikeRequest req)
    QueryVideoListResponse QueryVideoList(1:QueryVideoListRequest req)
    QueryLikeVideoListResponse QueryLikeVideoList(1:QueryLikeVideoListRequest req)
    FeedVideoListResponse FeedVideoList(1:FeedVideoListRequest req)


}
