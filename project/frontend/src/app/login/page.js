"use client";
import {Button, Input, Typography } from "@mui/material";
import { useEffect, useState } from "react";
import LoginService from "../service/login";
import { useRouter } from "next/navigation";
import Cookies from "js-cookie";
import toast, { Toaster } from "react-hot-toast";
export default function Login() {
  const [showPassword, setShowPassword] = useState(false);
  const [userName, setUserName] = useState("");
  const [passWord, setPassword] = useState("");
  const [error, setError] = useState(false)
  const [errorMg, setErrorMg] = useState(false)
  const router = useRouter();

  const login = async () => {
    
    const loginResult = await LoginService(userName, passWord);
    console.log(loginResult.status);
    
    if (loginResult.status) {
    setError(false)
     Cookies.set("token", loginResult.token, {
        expires: 7,
        path: "/",
        secure: true,
        sameSite: "Lax",
      });
      router.push("/");
    } else {
      toast.error(loginResult.message);
    }
  };

  return (
    <div className=" bg-blue-200 w-full min-h-[100vh] flex justify-center items-center">
      <div className="bg-white bg-opacity-10 w-[500px] h-[500px] rounded-3xl flex flex-col items-center justify-center gap-5 p-2 ">
        <Toaster position="top-center" reverseOrder={false} />
        <p className=" text-2xl">ล็อคอิน</p>
        <input
          id="username"
          type="text"
          placeholder="กรอกชื่อผู้ใช้"
          onChange={(e) => {
            setUserName(e.target.value);
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
        <div className="flex gap-2">
        <Button
        onClick={() => {
            router.push("/login/register")
          }}
          variant="outlined"
        >
          สมัครสมาชิก
        </Button>
         <Button
          onClick={() => {
            login();
          }}
          variant="contained"
        >
          เข้าสู่ระบบ
        </Button>
        </div>
         <p className={` text-red-500 text-2xl ${error ? "block" : "hidden"}`}>{errorMg}</p>
      </div>
    </div>
  );
}
