"use client";

import { Button, TextField } from "@mui/material";
import { useEffect, useState } from "react";
import List from "./component/list";
import {
  createQuote,
  editQuote,
  getQuoteAll,
  searchQuote,
  voting,
} from "./service/quote";
import toast, { Toaster } from "react-hot-toast";
import ModalBox from "./component/modal";
export default function Home() {
  const [quoteData, setQuoteData] = useState([]);
  const [quoteDataVoted, setQuoteDataVoted] = useState([]);
  const [selectQuote, setSelectQuote] = useState();
  const [selectText, setSelectText] = useState();
  const [inputText, setInputText] = useState();
  const [openEdit, setOpenEdit] = useState(false);
  const [openAdd, setOpenAdd] = useState(false);
  const [sort, setSort] = useState(false)
  const handleOpenEdit = () => {
    if (selectText && selectQuote) setOpenEdit(true);
  };
  const handleCloseEdit = () => setOpenEdit(false);

  const handleOpenAdd = () => setOpenAdd(true);
  const handleCloseAdd = () => setOpenAdd(false);

  useEffect(() => {
    const fetchData = async () => {
      const res = await getQuoteAll();
      if (res) {
        setQuoteDataVoted(res);
      }
    };
    fetchData();
  }, []);

  useEffect(() => {
    const fetchData = async () => {
      const res = await getQuoteAll();
      if (res) {
        setQuoteData(res);
      }
    };
    fetchData();
  }, [openAdd == false, openEdit == false]);

  const search = async () => {
    if (inputText == "") {
      const fetchData = async () => {
        const res = await getQuoteAll();
        if (res) {
          setQuoteData(res);
        }
      };
      fetchData();
    }
    const res = await searchQuote(inputText);
    if (res.status) {
      setQuoteData(res.data);
    } else {
      setQuoteData([]);
    }
  };

  const votingFunc = async () => {
    const res = await voting(selectQuote);
    console.log();

    if (res.status) {
      toast.success("คุณโหวตสำเร็จ!");
    } else {
      toast.error("คุณได้โหวตไปแล้ว!");
    }
  };

  const handleSubmitEdit = async (id, data) => {
    const res = await editQuote(id, data);
    if (res.status) {
      toast.success("แก้ไขสำเร็จ");
      handleCloseEdit();
    } else {
      toast.success(res.message);
      handleCloseEdit();
    }
    createQuote;
  };

  const handleSubmitAdd = async (data) => {
    const res = await createQuote(data);
    if (res.status) {
      toast.success("เพิ่มสำเร็จ");
      handleCloseAdd();
    } else {
      toast.success(res.message);
      handleCloseAdd();
    }
  };

useEffect(() => {
  const sorted = [...quoteDataVoted].sort((a, b) => {
    return sort ? b.Voted - a.Voted : a.Voted - b.Voted;
  });
  setQuoteDataVoted(sorted);
}, [sort]); // 🔁 รันทุกครั้งที่ sort เปลี่ยน


  return (
    <div className="bg-blue-200 flex justify-center items-center  w-full min-h-[100vh]">
      <div className="flex justify-center  w-[1200px] h-[500px] p-2  bg-white rounded-3xl ">
        <div className=" w-full basis-1/2 h-[400px]">
          <div className="flex flex-col gap-2 w-full h-[400px] max-h-[400px] justify-end">
            <div className="flex gap-2">
              <TextField
                onChange={(e) => {
                  setInputText(e.target.value);
                }}
                id="outlined-basic"
                label="ค้นหา"
                variant="outlined"
                InputProps={{
                  sx: {
                    height: "40px", // ✅ ใช้ตรงนี้แทน
                  },
                }}
              />
              <Button
                onClick={search}
                size="small"
                variant="outlined"
                sx={{
                  height: "40px",
                  color: "black",
                }}
              >
                ค้นหา
              </Button>
              <Button
                onClick={handleOpenAdd}
                size="small"
                variant="outlined"
                sx={{
                  height: "40px",
                  color: "black",
                }}
              >
                เพิ่มข้อความ
              </Button>
              <ModalBox
                open={openAdd}
                type={"Add"}
                handleClose={handleCloseAdd}
                handleSubmit={handleSubmitAdd}
              />
              <Button
                onClick={handleOpenEdit}
                size="small"
                variant="outlined"
                sx={{
                  height: "40px",
                  color: "black",
                }}
              >
                แก้ไข
              </Button>
              <ModalBox
                open={openEdit}
                id={selectQuote}
                data={selectText}
                type={"Edit"}
                handleClose={handleCloseEdit}
                handleSubmit={handleSubmitEdit}
              />
            </div>
            <List
               data={Array.isArray(quoteData) ? quoteData : []}
              onclick={(id, text) => {
                setSelectQuote(id);
                setSelectText(text);
              }}
            />
            <Button
              onClick={votingFunc}
              size="small"
              sx={{
                height: "40px", // ✅ ตั้งความสูง
                color: "white",
                backgroundColor: "blue",
                borderColor: "white",
              }}
            >
              โหวต
            </Button>
            <Toaster position="top-center" reverseOrder={false} />
          </div>
        </div>
        <div className=" w-full basis-1/2 h-[400px] gap-2">
          <div className="flex flex-col gap-2 w-full h-[400px] max-h-[400px] justify-end">
            <div className=" flex gap-2 h-[40px] w-full">

              <Button
               onClick={()=> setSort(!sort)}
                size="small"
                variant="outlined"
                sx={{
                  height: "40px",
                  color: "black",
                }}
              >
                {sort ? "DESC" : "ASC"}
              </Button>
            </div>
            <List
              data={Array.isArray(quoteDataVoted) ? quoteDataVoted : []}
              count={true}
              onclick={(id, text) => {
                setSelectQuote(id);
                setSelectText(text);
              }}
            />
          </div>
        </div>
      </div>
    </div>
  );
}
