## How am I going to go about this?

- My first thought is to encapsulate the machine in a struct:
    => the simplest and most obvious way

- According to the reference I am following, I should load the ROM next.
    1. Open the file.
    2. Seek to EOF, get n = length of file.
    3. Seek back to 0, read n bytes from the file.
    4. Transfer the bytes into Weep's memory, starting from 0x200.

- 