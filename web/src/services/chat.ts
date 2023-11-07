import { Resp, post, get } from './method';

/**
 * 发消息给某人
 * @param to_user_id 
 * @param action_type 
 * @param content 
 * @returns 
 */
export const sendMessage = (to_user_id : number, action_type: number, content: string): Resp<any> => post('/message/action', { to_user_id, action_type, content })

export type Message = {
    "id": number
    "to_user_id": number
    "from_user_id": number
    "content": string
    "create_time": number
}

/**
 * 获取与某人的聊天记录
 * @param to_user_id 
 * @returns 
 */

export const getMessage = (to_user_id : number,pre_msg_time?:number): Promise<{
    "code": number
    "message": string,
    "data":{
        message_list:Message[]
    }
}> => get('/message/chat', { to_user_id,pre_msg_time });
