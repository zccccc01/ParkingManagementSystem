import axios from 'axios';

import {
    findUserURL,
    followURL,
    unFollowURL,
    getFollowURL, 
} from './URLConfig';

class FriendManager {

    async findUser(nickname){
        try {
            const access_token = localStorage.access_token;
            const res = await axios.post(findUserURL,{
                access_token,
                nickname
            })
            return res.data;
        } catch (error) {
            return {
                success:false,
                errorMessage:'网络错误'
            }
        }  
    }

    async followUser(userId){
        try {
            const access_token = localStorage.access_token;
            const res = await axios.post(followURL,{
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

    async unFollowUser(userId){
        try {
            const access_token = localStorage.access_token;
            const res = await axios.post(unFollowURL,{
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

    async getFollows(){
        try {
            const access_token = localStorage.access_token;
            const res = await axios.post(getFollowURL,{
                access_token,
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

export default new FriendManager();
