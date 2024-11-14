# 智慧停车场管理系统 -React version

## 项目概述

这是一个基于 React 的智慧停车场管理系统前端项目，旨在提供一个高效、便捷的停车管理解决方案。系统包括用户注册、登录、停车场列表、预订、支付、停车记录、管理员后台等功能模块。

## 功能汇总

### 用户

- [x] 注册
- [x] 登录
- [x] 登记1-n辆车
- [x] 查看停车场实时车位情况
- [x] 预定车位
- [x] 支付费用(查出一个费用)
- [x] 查看历史停车记录

### 管理员

- [x] 登录
- [x] 查看停车位状态
- [x] 更新停车位状态
- [x] 统计总占用率(选某个停车场)
- [x] 统计总收入(选某个停车场)
- [x] 查看总违章,发送记录给用户
- [x] 查询某个时段内各停车场的车位占用情况
- [x] 查询某个用户的停车历史记录及支付情况
- [x] 按停车场的收入和停车位的使用率,生成月度或年度报告

### 其他

- [x] 404 页面 (Not Found Page): 处理未找到的页面请求。
- [x] 关于我们 (About Us Page): 介绍公司和项目背景。
- [x] 隐私政策 (Privacy Policy): 隐私政策说明。
- [x] 服务条款 (Terms of Service): 服务条款说明。

## 技术栈

- React: 前端框架，用于构建用户界面。
- React Router: 路由管理库，用于页面导航。
- React Helmet: 用于管理文档的 `<head>` 部分。
- CSS/SCSS: 样式管理。
- ES6+: 现代 JavaScript 语法。

## 安装与运行

- 克隆仓库
  ```sh
  git clone https://github.com/Rethymus/ParkingManagementSystem.git
  ```
- 安装依赖
  ```sh
  npm install
  ```
- 启动开发服务器
  ```sh
  cd ./frontend
  npm start
  ```
- 打包前端文件
  ```sh
  npm run build
  ```
- 启动一个本地服务器，查看打包后的文件
  ```sh
  npm serve -s build
  ```
- 访问应用 打开浏览器，访问 http://localhost:3000 即可看到应用。

## 依赖相关

- 尝试优化依赖树，减少重复的依赖项
  ```sh
  npm dedupe
  ```
- 删除 node_modules 文件夹中未在 package.json 中声明的包
  ```sh
  npm prune
  ```
- 清理 npm 缓存
  ```sh
  npm cache clean --force
  ```
- 更新 npm 包
  ```sh
  npm update
  ```
- 安装仅生产依赖项
  ```sh
  npm install --production
  ```
- 根据 package-lock.json 或 npm-shrinkwrap.json 文件重新安装所有依赖项
  ```sh
  rm -rf node_modules //手动删除 node_modules 目录
  npm cache clean --force //清理 npm 缓存
  npm ci //重新安装依赖
  ```
- 使用pnpm重新安装依赖
  ```sh
  npx pnpm install
  ```
- 强制 npm 获取远程资源
  ```sh
  npm rebuild
  ```
