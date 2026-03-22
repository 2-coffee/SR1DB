# SR1DB

A standard relational Database.

SR1DB/
├── Makefile
├── README.md             
│
├── include/              # header files (.h / .hpp) - The "What"
│   ├── pager.h           # Disk I/O and memory buffer definitions
│   ├── page_layout.h     # PageHeader, LinePointer, and memory cast logic
│   ├── row.h             # Tuple struct and serialization definitions
│   ├── cursor.h          # Database iterator definitions
│   └── table.h           # High-level engine logic (Insert, Select, Delete)
│
├── src/                  # Your source files (.cpp) - The "How"
│   ├── pager.cpp         # fstream logic to read/write 4KB chunks
│   ├── page_layout.cpp   # Logic to calculate free space and offsets
│   ├── row.cpp           # memcpy logic to pack/unpack bytes
│   ├── cursor.cpp        # Logic to advance through slots and pages
│   └── table.cpp         # Tying it all together to execute operations
│
├── cli/                  # Your user interface
│   └── main.cpp          # The REPL (Read-Eval-Print Loop) prompt
│
└── tests/                # (Highly Recommended)
    ├── test_pager.cpp    # Write a page, read it back, assert they match
    ├── test_row.cpp      # Serialize a row, deserialize it, assert match
    └── test_table.cpp    # Insert 10k rows and ensure no crashes