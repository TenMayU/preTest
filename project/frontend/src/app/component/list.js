'use client'
import { useState } from "react";

const List =({data,onclick,count})=>{
 const [selectedId, setSelectedId] = useState(null);

  return (
    <div className="w-full h-[400px] flex flex-col gap-2 overflow-y-scroll">
      {
      data?.length > 0 ?       
      data?.map((item) => (
        <div
          key={item?.ID}
          onClick={() => {setSelectedId(item?.ID); onclick(item?.ID,item?.Text)}}
          className={`cursor-pointer flex justify-between px-4 py-2 rounded-md border transition-colors duration-200 ${
            selectedId === item?.ID ? "bg-blue-100 border-blue-400" : "bg-white border-gray-300"
          }`}
        >
          <p className="text-sm text-gray-800">{item?.Text}</p>
          {
            count ? <p className="text-sm text-gray-800">{item?.Voted}</p> : ""
          }
          
        </div>
      ))
      
      : ""
      }
    </div>
  );
}

export default List