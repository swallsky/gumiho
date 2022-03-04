import React from "react";

import { Button } from "antd";

import { Outlet } from "react-router-dom";

import "./App.css";

export default function App() {
  return (
    <div>
      <Button type="primary">累计</Button>
      <Outlet />
    </div>
  );
}
