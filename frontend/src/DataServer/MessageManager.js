import axios from 'axios';

import {
    postMessageURL,
    getMessageURL,
    homeMessageURL,
    deleteMessageURL,
} from './URLConfig';

class MessageManager {
    async postMessage(content,images){
        try {
            const access_token = localStorage.access_token;
            const formData = new FormData();
            formData.append('access_token',access_token);
            formData.append('content',content);
            if (images) {
                images.map((image,index)=>{
                    formData.append(`image${index+1}`,image.file);
                    return '';
                })
            }

            const res = await axios({
                url:postMessageURL,
                method:'POST',
                data:formData,
                headers: {'Content-Type': 'multipart/form-data'}
            })
            return res.data;

        } catch (error) {
            return {
                success:false,
                errorMessage:'网络错误'
            }
        }
    }

    async deleteMessage(messageId){
        try {
            const access_token = localStorage.access_token;
            const res = await axios.post(deleteMessageURL,{
                access_token,
                messageId
            })
            return res.data;
        } catch (error) {
            return {
                success:false,
                errorMessage:'网络错误'
            }
        }  
    }

    async homeMessage(minId){
        try {
            const access_token = localStorage.access_token;
            const res = await axios.post(homeMessageURL,{
                access_token,
                minId
            })
            return res.data;
        } catch (error) {
            return {
                success:false,
                errorMessage:'网络错误'
            }
        }  
    }

    async getUserMessage(userId,minId){
        try {
            const access_token = localStorage.access_token;
            const res = await axios.post(getMessageURL,{
                access_token,
                userId,
                minId
            })
            return res.data;
        } catch (error) {
            return {
                success:false,
                errorMessage:'网络错误'
            }
        }  
    }




}

export default new MessageManager();
