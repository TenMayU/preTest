//// ควรทำแยกเป็น API interface กลางไว้ใช้เลย จะได้ไม่จำเป็นต้องSetting ใหม่ แต่อันนี้จำเป็นต้องทำ เพื่อความง่ายไม่สับสน 
const LoginService = async (username, password) => {
  try {
    const apiUrl = process.env.NEXT_PUBLIC_API_URL;
    const res = await fetch(`${apiUrl}/auth/login`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        userName: username, // เปลี่ยนเป็นค่าจาก state
        Password: password, // เปลี่ยนเป็นค่าจาก state
      }),
    });

    const data = await res.json();

    if (data.token && data.token !== "") {
      console.log("Login success:", data.token);
      return data;
    }

    return data;
  } catch (error) {
    console.log("เข้าสู่ระบบไม่สำเร็จ");
  }
};

export default LoginService;
