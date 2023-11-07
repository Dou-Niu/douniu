import { defineStore } from 'pinia';
import { Video } from '@/types';

interface VideoInfo {
    video_list: Video[],
    currentVideo:Video,
    currentIndex:number
}
export const video = defineStore('video', {
    state: (): VideoInfo => {
        return {
            video_list: [],
            currentVideo: {
                video_id: 0,
                author:{
                    id : (0),
                    phone: "",
                    nickname: "",
                    follow_count: 0,
                    follower_count: 0,
                    is_follow: false,
                    avatar: "",
                    background_image: "",
                    signature: "",
                    total_favorited: "",
                    work_count: 0,
                    favorite_count: 0,
                    collection_count: 0
                },
                play_url:"",
                cover_url:"",
                favorite_count:0,
                collection_count:0,
                comment_count:0,
                is_favorite:false,
                is_collect:false,
                title:"",
                partition:0,
                create_time:new Date(),
            },
            currentIndex:0
        }
    },
    actions: {
        setVideoList(video_list: Video[]): void{
            this.video_list = video_list
        },
        getVideoList(){
            return this.video_list
        },
        setCurrentVideo(video:Video): void{
            if(this.video_list.length === 0){
                this.currentVideo = video
            }else{
                this.currentVideo = this.video_list[this.currentIndex]
            }
        },
        setCurrentIndex(index:number): void{
            this.currentIndex = index
            this.currentVideo = this.video_list[index]
        },
        getCurrentVideo(){
            return this.currentVideo
        }
    }
})
