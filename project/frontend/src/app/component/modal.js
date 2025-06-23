'use client'
import {
  Box,
  Button,
  Modal,
  TextField,
  Typography,
} from '@mui/material';
import { useEffect, useState } from 'react';

  const style = {
    position: "absolute",
    top: "50%",
    left: "50%",
    transform: "translate(-50%, -50%)",
    width: 400,
    bgcolor: "background.paper",
    borderRadius: 2,
    boxShadow: 24,
    p: 4,
  };

export default function ModalBox({
  open,
  id,
  data,
  type,
  handleClose,
  handleSubmit,
}) {
  const [selectText, setSelectText] = useState('');


  useEffect(() => {
    if (type === 'Edit' && data) {
      setSelectText(data);
    } else {
      setSelectText('');
    }
  }, [open, data, type]);

  const onSubmit = () => {
    if (type === 'Edit') {
      handleSubmit(id, selectText);
    } else {
      handleSubmit(selectText);
    }
  };

  return (
    <Modal
      open={open}
      onClose={handleClose}
      aria-labelledby="modal-form-title"
      aria-describedby="modal-form-description"
    >
      <Box sx={style}>
        <Typography id="modal-form-title" variant="h6" component="h2" mb={2}>
          {type === 'Edit' ? 'แก้ไขข้อความ' : 'เพิ่มข้อความ'}
        </Typography>
        <TextField
          fullWidth
          label="ข้อความ"
          variant="outlined"
          margin="normal"
          value={selectText}
          onChange={(e) => setSelectText(e.target.value)}
        />
        <Box mt={3} display="flex" justifyContent="flex-end" gap={1}>
          <Button variant="outlined" onClick={handleClose}>
            ยกเลิก
          </Button>
          <Button variant="contained" onClick={onSubmit}>
            ส่งข้อมูล
          </Button>
        </Box>
      </Box>
    </Modal>
  );
}