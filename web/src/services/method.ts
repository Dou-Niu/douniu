import  request  from '../utils/request';

export type Resp<T> = Promise<{
    code: number,
    message: string,
    data: T
}>

export const get = (url: string, params?: any): any => request({
    url,
    method: 'GET',
    params
})
export const post = (url: string, data?: any): any => request({
    url,
    method: 'POST',
    data
})
export const put = (url: string, data?: any): any => request({
    url,
    method: 'PUT',
    data
})
export const del = (url: string, params?: any): any => request({
    url,
    method: 'delete',
    params
})
