import React, { useState, useEffect } from 'react';
import { Menu, MenuProps } from 'antd';
import { NavLink as Link, useLocation } from 'react-router-dom';
import './Sidebar.scss';

type MenuItem = Required<MenuProps>['items'][number];

function getItem(
  label: React.ReactNode,
  key: React.Key,
  icon?: React.ReactNode,
  children?: MenuItem[]
): MenuItem {
  return {
    key,
    icon,
    children,
    label,
  } as MenuItem;
}

const menuList = [
  { value: '/', label: '首页', icon: '' },
  { value: '/dashboard', label: '仪表盘', icon: '' },
  { value: '/parking-records', label: '停车记录', icon: '' },
  { value: '/vehicle-info', label: '车辆信息', icon: '' },
  { value: '/violations', label: '违章停车', icon: '' },
  { value: '/parking-space', label: '停车位', icon: '' },
  { value: '/bookings', label: '停车位预约', icon: '' },
  { value: '/parking-lots', label: '停车场列表', icon: '' },
  { value: '/admin-dashboard', label: '管理员仪表盘', icon: '' },
  { value: '/space-status', label: '查询车位状态', icon: '' },
  { value: '/check-space', label: '查看车位', icon: '' },
  { value: '/user', label: '用户页面', icon: '' },
];

const routeMap: { [key: string]: string } = {};

function getMenuList(): MenuItem[] {
  const tempMenuList: MenuItem[] = [];
  const getList = (list: any[], newList: MenuItem[]) => {
    for (let i = 0; i < list.length; i += 1) {
      const { value, label, icon } = list[i] || {};
      const it = getItem(
        <Link to={value}>{label}</Link>,
        value || label,
        icon && <span>{icon}</span>
      );
      newList.push(it);
      routeMap[value] = label;
    }
  };
  getList(menuList, tempMenuList);
  return tempMenuList;
}

const Sidebar: React.FC = () => {
  const path = useLocation().pathname;
  const [tempPath, setTempPath] = useState(path);
  const tempMenuList = getMenuList();

  useEffect(() => {
    setTempPath(path);
  }, [path]);

  const onClick: MenuProps['onClick'] = (e) => {
    setTempPath(e.key);
  };

  return (
    <div className="sidebar">
      <h2>菜单</h2>
      <Menu onClick={onClick} mode="inline" selectedKeys={[tempPath]} items={tempMenuList} />
    </div>
  );
};

export default Sidebar;
