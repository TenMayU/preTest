
//// ควรทำแยกเป็น API interface กลางไว้ใช้เลย จะได้ไม่จำเป็นต้องSetting ใหม่ แต่อันนี้จำเป็นต้องทำ เพื่อความง่ายไม่สับสน 
export default async function  RegisterService(username, name , password) {
   try {
     const apiUrl = process.env.NEXT_PUBLIC_API_URL
    const res = await fetch(`${apiUrl}/auth/register`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        userName: username, 
        name: name,
        Password: password,     
      }),
    });
    const data = await res.json();

    if (data.status) {
      return true;
    }

    return false;
   } catch (error) {
    console.log('สมัครสมาชิกไม่สำเร็จ');
   } 
   
}