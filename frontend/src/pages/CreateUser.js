import React, { Component } from 'react'

import { 
    Button,
    Toast,
    NavBar,
    WingBlank, 
    WhiteSpace ,
    List,
    InputItem,
    Icon,
    ImagePicker,
    Modal
} from 'antd-mobile';

import userManager from '../DataServer/UserManager.js';



export default class CreateUserScreen extends Component {


    constructor(props) {
      super(props)

      this.state = {
         nickname:'',
         sign:'',
         files:[]
      }
    }


  render() {
    return (
      <div>
        <NavBar
            mode="dark"
            icon={<Icon type="left" />}
            onLeftClick={() => {this.props.history.goBack()}}
        >完善用户信息</NavBar>
        <WhiteSpace/>
        <ImagePicker
            files={this.state.files}
            onChange={(files)=>{this.setState({files})}}
            selectable={this.state.files.length < 1}
        />
        <List>
            <InputItem
                type={'text'}
                value={this.state.nickname}
                onChange={(nickname)=>{this.setState({nickname})}}
                placeholder={'请输入昵称'}
            >
                昵称
            </InputItem>
            <InputItem
                type={'text'}
                value={this.state.sign}
                onChange={(sign)=>{this.setState({sign})}}
                placeholder={'请输入个性签名'}
            >
                签名
            </InputItem>
        </List>
        <WhiteSpace/>
        <WingBlank>
            <Button
                type={'primary'}
                onClick={async()=>{

                    if(this.state.files.length === 0){
                        Toast.fail('请选择头像',1);
                        return;
                    }

                    Toast.loading('内容上传中...',0);
                    const resutl = await userManager.createUser(this.state.nickname,this.state.sign,this.state.files[0]);
                    Toast.hide();
                    if(resutl.success === false){
                        Toast.fail(resutl.errorMessage);
                        return;
                    }
                    Modal.alert('成功','点击确认键进入主页',[{
                        text:'确认',
                        onPress:()=>{this.props.history.replace('/TabBarScreen')}
                    }])

                }}
            >
                完成注册
            </Button>
        </WingBlank>
      </div>
    )
  }
}
