interface Resp {
    blob: Blob,
    url: string,
    duration?: number
}

// 将图片转为base64
export function getBase64(url: string) {
    const img = document.createElement("img");
    img.src = url;
    const cvs = document.createElement("canvas");
    cvs.width = img.width;
    cvs.height = img.height;
    const ctx = cvs.getContext("2d");
    (ctx as CanvasRenderingContext2D).drawImage(img, 0, 0, cvs.width, cvs.height);
    const ext = img.src.substring(img.src.lastIndexOf(".") + 1).toLowerCase();
    let dataUrl = cvs.toDataURL("image/jpeg" + ext);
    return dataUrl;
}

export const getCover = (file: File, time: number = 0): Promise<Resp> => {
    return new Promise((resolve, reject) => {
        try {
            let src = URL.createObjectURL(file);
            let video = document.createElement('video');
            video.src = src;
            video.currentTime = time;
            video.muted = true;
            video.autoplay = true;
            video.oncanplay = async () => {
                const frame = await drawVideo(video);
                frame.duration = video.duration;
                resolve(frame)
            }
        } catch (error: any) {
            reject(error.message)
        }
    })
}
function drawVideo(vdo: HTMLVideoElement): Promise<Resp> {
    return new Promise((resolve, reject) => {
        try {
            const cvs = document.createElement("canvas") as HTMLCanvasElement; // 创建canvas元素
            const ctx = cvs.getContext("2d") as any; // 返回一个canvas渲染的上下文
            cvs.width = vdo.videoWidth; // 获取视频的宽度
            cvs.height = vdo.videoHeight; // 获取视频的高度
            ctx.drawImage(vdo, 0, 0, cvs.width, cvs.height); // 从视频的左上角画满视频的宽高
            cvs.toBlob((blob) => {
                resolve({
                    blob: blob as Blob,
                    url: URL.createObjectURL(blob as Blob),
                });
            });
        } catch (error: any) {
            reject(error.messsage)
        }
    });
}




