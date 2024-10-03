import { loginURL, registerURL, changePasswordURL } from './URLConfig';
import axios from 'axios';

class AccountManager {
    async register(email,password){
        try {
            const res = await axios.post(registerURL,{
                email,
                password
            });
            const result = res.data;
            if (result.success === true) {
                localStorage.access_token = result.data.access_token;
            }
            return result;
        } catch (error) {
            return {
                success:false,
                errorMessage: '网络错误'
            }
        }

    }

    async login(email,password){
        try {
            const res = await axios.post(loginURL,{
                email,
                password
            });
            const result = res.data;
            if (result.success === true) {
                localStorage.access_token = result.data.access_token;
            }
            return result;
        } catch (error) {
            return {
                success:false,
                errorMessage: '网络错误'
            }
        }

    }

    async changePassword(old_password,new_password){
        try {
            const access_token = localStorage.access_token;
            const res = await axios.post(changePasswordURL,{
                access_token,
                old_password,
                new_password
            });
            const result = res.data;
            return result;
        } catch (error) {
            return {
                success:false,
                errorMessage: '网络错误'
            }
        }

    }

    isLogin(){
        if(localStorage.access_token === '' || !localStorage.access_token){
            return false;
        } else {
            return true;
        }
    }

    logout(){
        localStorage.access_token = '';
    }

}

export default new AccountManager();
