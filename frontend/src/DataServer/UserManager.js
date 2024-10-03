import {
    createUserURL,
    updateUserURL,
    getUserURL,
} from './URLConfig';

import axios from 'axios';

class UserManager {
    async createUser(nickname,sign,image){
        try {
            const access_token = localStorage.access_token;
            const formData = new FormData();
            formData.append('access_token',access_token);
            formData.append('nickname',nickname);
            formData.append('sign',sign);
            formData.append('image',image.file);

            const res = await axios({
                url:createUserURL,
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

    async updateUser(userInfo){
        try {
            const access_token = localStorage.access_token;
            const formData = new FormData();
            formData.append('access_token',access_token);
            if (userInfo.nickname) {
                formData.append('nickname',userInfo.nickname);
            }
            if (userInfo.sign) {
                formData.append('sign',userInfo.sign);
            }
            if (userInfo.image) {
                formData.append('image',userInfo.image.file);
            }

            const res = await axios({
                url:updateUserURL,
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

    async getUserInfo(userId = 0){
        try {
            const access_token = localStorage.access_token;
            const res = await axios.post(getUserURL,{
                access_token,
                userId
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

export default new UserManager();
