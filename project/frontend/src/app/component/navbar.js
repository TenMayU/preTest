"use client";
import { Button } from "@mui/material";
import Cookies from "js-cookie";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";
const Navber = () => {
 const router = useRouter();
 const [token, setToken] = useState("")
  const logout = () => {
    Cookies.remove("token", { path: "/" }); // ลบ token ที่เคย set
    router.push("/")
    console.log("ออกจากระบบแล้ว");
  };

useEffect(() => {
  const t = Cookies.get("token");
  setToken(t || ""); // ✅ set เป็น "" ถ้าไม่มี token แล้ว
}, [token]); // ✅ ใช้ token เป็น dependency
  

  
  return (
    <div className=" w-full h-15 flex justify-end fixed">
     {token && (
        <Button
          onClick={logout}
          size="small"
          sx={{
            color: "black",
            backgroundColor: "white",
            borderColor: "white",
          }}
        >
          ออกจากระบบ
        </Button>
      )}
    </div>
  );
};

export default Navber;
