import * as qiniu from 'qiniu-js';
import CryptoJS from 'crypto-js';
import { ref } from 'vue';
import { getCover } from './getCover';

interface Resp {
    videoURL: string
}

// 上传进度
export let uploadPercentage = ref(0);
export let videoCovers = ref<{ url: string, blob: Blob }[]>([]);
export let curCover = ref('');
export let curVideoURL = ref('');
export let curCoverFile = ref<File>();

let AccessKey = "8kRDQhbwFp3LCh4-wyksj82zT1uEJ7ru0wqOUlVq"
let SKey = "4bZqtevQ92UiCkc0vPvOc8IRNTJM4iGLVbwTRi3K"
const bucket = "kecatpostimage"

const getPutPolicy = (bucket: string, key: string) => {
    return JSON.stringify({
        scope: bucket + ':' + key,
        deadline: new Date().getTime() + 3600
    })
}
// 上传凭证
const getToken = (encodeSign: string, encodePutPolicy: string) => {
    return AccessKey + ':' + encodeSign + ':' + encodePutPolicy
}
// 生成封面
const generateCovers = async (file: File) => {
    const { url: url1, duration, blob: blob1 } = await getCover(file);
    videoCovers.value.push({
        url: url1,
        blob: blob1
    });
    curCover.value = url1;
    curCoverFile.value = await (async () => {
        return new window.File([blob1], url1, { type: 'image/webp' })
    })();
    const { url: url2, blob: blob2 } = await getCover(file, Math.floor(Math.random() * (duration as number)));
    videoCovers.value.push({
        url: url2,
        blob: blob2
    });
    const { url: url3, blob: blob3 } = await getCover(file, Math.floor(Math.random() * (duration as number)));
    videoCovers.value.push({
        url: url3,
        blob: blob3
    });
}

export const uploadVideo = async (file: File): Promise<Resp> => {
    return new Promise<Resp>((resolve, reject) => {
        try {
            let key = 'video/' + file.name;
            let putPolicy = getPutPolicy(bucket, key);
            let encodedPutPolicy = CryptoJS.enc.Base64.stringify(CryptoJS.enc.Utf8.parse(putPolicy)).replace(/\+/g, '-').replace(/\//g, '_');
            let encodedSign = CryptoJS.enc.Base64.stringify(CryptoJS.HmacSHA1(encodedPutPolicy, SKey)).replace(/\+/g, '-').replace(/\//g, '_'); // 第一个参数为加密字符串，第二个参数为公共秘钥
            let token = getToken(encodedSign, encodedPutPolicy);
            const observable = qiniu.upload(file, key, token, {}, {
                region: qiniu.region.z2
            })
            observable.subscribe({
                next(res) {
                    // 更新进度
                    uploadPercentage.value = Math.ceil(res.total.percent);
                },
                error(err) {
                    reject(err.message)
                },
                complete(res) {
                    resolve({ videoURL: 'https://www.kecat.top/' + res.key })
                    curVideoURL.value = 'https://www.kecat.top/' + res.key;
                    uploadPercentage.value = 0;
                    generateCovers(file);
                }
            }) // 上传开始
        } catch (error: any) {
            reject(error.message)
        }
    })
}