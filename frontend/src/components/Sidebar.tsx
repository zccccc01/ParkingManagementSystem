import React, { useState, useEffect, useMemo } from 'react';
import { Menu, Modal, Input, Button, MenuProps } from 'antd';
import {
  HomeOutlined,
  DashboardOutlined,
  CarOutlined,
  FileSearchOutlined,
  DollarCircleOutlined,
  CalendarOutlined,
  UserOutlined,
  LockOutlined,
} from '@ant-design/icons';
import { useLocation, useNavigate } from 'react-router-dom';
import './Sidebar.scss';

type MenuItem = {
  key: string;
  icon?: React.ReactNode;
  children?: MenuItem[];
  label: React.ReactNode;
};

const getItem = (
  label: React.ReactNode,
  key: string,
  icon?: React.ReactNode,
  children?: MenuItem[]
): MenuItem => ({
  key,
  icon,
  children,
  label,
});

const menuList: MenuItem[] = [
  { key: '/', label: '首页', icon: <HomeOutlined /> },
  {
    key: '/dashboard',
    label: '用户仪表盘',
    icon: <DashboardOutlined />,
    children: [
      { key: '/parking-records', label: '停车记录', icon: <FileSearchOutlined /> },
      { key: '/create-vehicle', label: '绑定车辆', icon: <CarOutlined /> },
      { key: '/vehicle-info', label: '车辆信息', icon: <CarOutlined /> },
      { key: '/violations', label: '违章停车', icon: <LockOutlined /> },
      { key: '/bookings', label: '停车位预约', icon: <CalendarOutlined /> },
      { key: '/payments', label: '支付', icon: <DollarCircleOutlined /> },
    ],
  },
  { key: '/parking-space', label: '空闲车位', icon: <CarOutlined /> },
  { key: '/parking-lots', label: '停车场列表', icon: <HomeOutlined /> },
  { key: '/user', label: '用户页面', icon: <UserOutlined /> },
  { key: '/admin-dashboard', label: '管理员面板', icon: <LockOutlined /> },
];

const Sidebar: React.FC = () => {
  const path = useLocation().pathname;
  const navigate = useNavigate();
  const [openKeys, setOpenKeys] = useState<string[]>([]);
  const [passwordModalVisible, setPasswordModalVisible] = useState(false);
  const [password, setPassword] = useState('');

  const tempMenuList = useMemo(() => {
    const buildMenu = (list: MenuItem[]): MenuItem[] =>
      list.map(({ key, label, icon, children }) => ({
        key,
        icon,
        label,
        children: children ? buildMenu(children) : undefined,
      }));
    return buildMenu(
      menuList.map((item) => getItem(item.label, item.key, item.icon, item.children))
    );
  }, []);

  useEffect(() => {
    // 这里不再需要更新 tempPath，因为我们直接使用 navigate 进行跳转
  }, [path]);

  const onOpenChange: MenuProps['onOpenChange'] = (keys) => {
    setOpenKeys(keys);
  };

  const handleMenuClick: MenuProps['onClick'] = ({ key }) => {
    if (key === '/admin-dashboard') {
      setPasswordModalVisible(true); // 显示密码输入模态框
    } else {
      navigate(key); // 对于其他菜单项，直接跳转
    }
  };

  const handlePasswordSubmit = () => {
    if (password === 'root') {
      navigate('/admin-dashboard'); // 密码正确，跳转到管理员面板
      setPasswordModalVisible(false); // 关闭密码输入模态框
    } else {
      // 密码错误，可以在这里添加错误提示
      console.log('密码错误');
    }
  };

  const handleKeyPress = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === 'Enter') {
      handlePasswordSubmit();
    }
  };

  return (
    <div className="sidebar">
      <h2>菜单</h2>
      <Menu
        onClick={handleMenuClick}
        onOpenChange={onOpenChange}
        mode="vertical"
        inlineIndent={24}
        selectedKeys={[path]} // 使用当前路径作为选中的菜单项
        openKeys={openKeys}
        items={tempMenuList}
      />
      {passwordModalVisible && (
        <Modal
          title="请输入密码"
          visible={passwordModalVisible}
          onCancel={() => setPasswordModalVisible(false)}
          footer={null}
        >
          <Input
            type="password"
            className="password-input"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            placeholder="请输入密码"
            onKeyPress={handleKeyPress} // 添加 onKeyPress 事件处理函数
          />
          <br />
          <br />
          <Button className="submit-button" type="primary" onClick={handlePasswordSubmit}>
            提交
          </Button>
        </Modal>
      )}
    </div>
  );
};

export default Sidebar;
