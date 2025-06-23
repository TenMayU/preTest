import Cookies from "js-cookie";

function decodeToken(token) {
  try {
    const payload = token.split('.')[1];
    const decoded = atob(payload);
    return JSON.parse(decoded);
  } catch (e) {
    return null;
  }
}

//// ควรทำแยกเป็น API interface กลางไว้ใช้เลย จะได้ไม่จำเป็นต้องSetting ใหม่ แต่อันนี้จำเป็นต้องทำ เพื่อความง่ายไม่สับสน 

export const getQuoteAll = async () => {
  try {
    const t = Cookies.get("token");
    const apiUrl = process.env.NEXT_PUBLIC_API_URL;
    const res = await fetch(`${apiUrl}/quote/getAllQuote`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        "Authorization": `Bearer ${t}`
      },
    });

    const { data, message } = await res.json();

    if (data.length > 0) {
      return data;
    }
    return message;
  } catch (error) {
    return error;
  }
};

export const voting = async (id) => {
  try {
    const t = Cookies.get("token");
    const userData = decodeToken(t);
    console.log(userData); 
    const apiUrl = process.env.NEXT_PUBLIC_API_URL;
    const res = await fetch(`${apiUrl}/quote/voting`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Authorization": `Bearer ${t}`
      },
      body: JSON.stringify({
        QuoteId: id, 
        User: userData.user_id
      }),
    });

    const { message, status } = await res.json();
    console.log(message,status);
    
    return {message:message, status:status};
  } catch (error) {
    return error;
  }
};


export const editQuote  = async (id,data) =>{
   try {
    const t = Cookies.get("token");
    const apiUrl = process.env.NEXT_PUBLIC_API_URL;
    const res = await fetch(`${apiUrl}/quote/update`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Authorization": `Bearer ${t}`
      },
      body: JSON.stringify({
        ID: id, 
        Text: data
      }),
    });

    const { message, status } = await res.json();
    if(status){
      return { message, status }
    }

    return { message, status }
   } catch (error) {
    return error;
   }

}


export const createQuote  = async (data) =>{
   try {
    const t = Cookies.get("token");
    const apiUrl = process.env.NEXT_PUBLIC_API_URL;
    const res = await fetch(`${apiUrl}/quote/create`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Authorization": `Bearer ${t}`
      },
      body: JSON.stringify({
        Text: data
      }),
    });

    const { message, status } = await res.json();
    if(status){
      return { message, status }
    }

    return { message, status }
   } catch (error) {
    return error;
   }

}


export const searchQuote = async (body) =>{
  try {
    const t = Cookies.get("token");
    const apiUrl = process.env.NEXT_PUBLIC_API_URL;
    const res = await fetch(`${apiUrl}/quote/search`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Authorization": `Bearer ${t}`
      },
      body: JSON.stringify({
        Text: body
      }),
    });

    const { message, status, data } = await res.json();
    if(status){
      return { message, status , data }
    }

    return { message, status, data }
   } catch (error) {
    return error;
   }

}

