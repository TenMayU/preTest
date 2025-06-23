"use client";
import {Button, Input, Typography } from "@mui/material";
import Alert from '@mui/material/Alert';
import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import Cookies from "js-cookie";
import RegisterService from "@/app/service/register";
import toast, { Toaster } from "react-hot-toast";
export default function Register() {
  const [showPassword, setShowPassword] = useState(false);
  const [userName, setUserName] = useState("");
  const [name, setName] = useState("");
  const [passWord, setPassword] = useState("");
  const router = useRouter();

  const Register = async () => {
    const registerResult = await RegisterService(userName, name, passWord);
    if (registerResult) {
    toast.error("สมัครสมาชิกสำเร็จ");
      router.push("/");
    } else {
    toast.error(registerResult.message);
    }
  };

  return (
    <div className=" bg-blue-200 w-full min-h-[100vh] flex justify-center items-center">
      <div className="bg-white bg-opacity-10 w-[500px] h-[500px] rounded-3xl flex flex-col items-center justify-center gap-5 p-2 ">
       <Toaster position="top-center" reverseOrder={false} />
        <p className=" text-2xl">สมัครสมาชิก</p>
        <input
          id="username"
          type="text"
          placeholder="กรอกชื่อผู้ใช้"
          onChange={(e) => {
            setUserName(e.target.value);
          }}
          className="w-full rounded border border-gray-300 px-4 py-2 focus:outline-none focus:ring-2 focus:ring-blue-400"
        />
          <input
          id="username"
          type="text"
          placeholder="กรอกชื่อ"
          onChange={(e) => {
            setName(e.target.value);
          }}
          className="w-full rounded border border-gray-300 px-4 py-2 focus:outline-none focus:ring-2 focus:ring-blue-400"
        />
        <div className="flex gap-2 w-full">
          <input
            id="password"
            type={`${showPassword ? "text" : "password"}`}
            placeholder="รหัสผ่าน"
            onChange={(e) => {
              setPassword(e.target.value);
            }}
            className="w-full rounded border border-gray-300 px-4 py-2 focus:outline-none focus:ring-2 focus:ring-blue-400"
          />
          <Button
            onClick={() => {
              setShowPassword(!showPassword);
            }}
            variant="outlined"
          >
            {showPassword ? "ปิดรหัสผ่าน" : "เปิดรหัสผ่าน"}
          </Button>
        </div>
        <Button
          onClick={() => {
            Register();
          }}
          variant="contained"
        >
          สมัครสมาชิก
        </Button>
      </div>
    </div>
  );
}
