import React, { useState, useEffect } from 'react';
import { Menu, MenuProps } from 'antd';
import { NavLink as Link, useLocation } from 'react-router-dom';
import {
  HomeOutlined,
  BarChartOutlined,
  PieChartOutlined,
  DollarCircleOutlined,
  SearchOutlined,
  EditOutlined,
  EyeOutlined,
  LeftOutlined,
} from '@ant-design/icons';
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
  { value: '/admin-dashboard', label: '首页', icon: <HomeOutlined /> },
  { value: '/statistic', label: '统计信息', icon: <BarChartOutlined /> },
  { value: '/chart', label: '可视化状态', icon: <PieChartOutlined /> },
  { value: '/parking-lot-income', label: '收入图', icon: <DollarCircleOutlined /> },
  { value: '/check-status', label: '查询车位状态', icon: <SearchOutlined /> },
  { value: '/update-status', label: '更新车位状态', icon: <EditOutlined /> },
  { value: '/check-space', label: '查看车位', icon: <EyeOutlined /> },
  { value: '/', label: '返回', icon: <LeftOutlined /> },
];

const routeMap: { [key: string]: string } = {};

function getMenuList(): MenuItem[] {
  const tempMenuList: MenuItem[] = [];
  const getList = (list: any[], newList: MenuItem[]) => {
    for (let i = 0; i < list.length; i += 1) {
      const { value, label, icon, children } = list[i] || {};
      const it = getItem(
        <Link to={value}>{label}</Link>,
        value || label,
        icon,
        children &&
          children.map((child) =>
            getItem(<Link to={child.value}>{child.label}</Link>, child.value, child.icon)
          )
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
  const [openKeys, setOpenKeys] = useState<string[]>([]);
  const tempMenuList = getMenuList();

  useEffect(() => {
    setTempPath(path);
  }, [path]);

  const onOpenChange: MenuProps['onOpenChange'] = (keys) => {
    const latestOpenKey = keys.find((key) => openKeys.indexOf(key) === -1);
    if (latestOpenKey) {
      setOpenKeys(keys);
    } else {
      setOpenKeys(keys.filter((key) => openKeys.includes(key)));
    }
  };

  const handleMenuClick: MenuProps['onClick'] = (e) => {
    if (e.key === '/dashboard') {
      if (openKeys.includes('/dashboard')) {
        setOpenKeys(openKeys.filter((key) => key !== '/dashboard'));
      } else {
        setOpenKeys(['/dashboard']);
      }
    } else {
      setTempPath(e.key);
    }
  };

  return (
    <div className="sidebar">
      <h2>管理员菜单</h2>
      <Menu
        onClick={handleMenuClick} // 处理菜单项点击事件
        onOpenChange={onOpenChange} // 处理子菜单项展开/收起
        mode="vertical" // 设置菜单模式为内联
        inlineIndent={24} // 设置子菜单项的缩进量
        selectedKeys={[tempPath]} // 当前选中的菜单项
        openKeys={openKeys} // 当前展开的子菜单项
        items={tempMenuList} // 菜单项列表
      />
    </div>
  );
};

export default Sidebar;
