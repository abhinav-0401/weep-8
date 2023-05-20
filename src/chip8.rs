struct Chip8 {
    registers: [u8; 16],
    memory: [u8; 16],
    index_reg: u16,
    pc: u16,
    stack: [u16; 16],
    sp: u8,
    delay_timer: u8,
    sound_timer: u8,
    vid_buff: [u32; 64*32],
    opcode: u16,
}