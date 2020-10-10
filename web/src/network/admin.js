import {request} from "./request";

// 上传视频
export function uploadVideo(data) {
    return request({
        url: "/v0/videos",
        method: "post",
        data: data
    })
}

// 编辑视频
export function editVideo(vid, data) {
    return request({
        url: `/v0/videos/${vid}`,
        method: "put",
        data: data
    })
}

// 删除视频
export function deleteVideo(vid) {
    return request({
        url: `/v0/videos/${vid}`,
        method: "delete"
    })
}
