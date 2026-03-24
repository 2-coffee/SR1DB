# SR1DB

A standard relational Database.

SR1DB/
├── README.md             
│
├── src/                  
│   ├── pager         # fstream logic to read/write 4KB chunks
│   ├── page_layout   # Logic to calculate free space and offsets
│   ├── row           # memcpy logic to pack/unpack bytes
│   ├── cursor        # Logic to advance through slots and pages
│   └── table         
│
├── cli/                  
│   └── main          # The REPL (Read-Eval-Print Loop) prompt
│
└── tests/                
    ├── test_pager    # Write a page, read it back, assert they match
    ├── test_row      # Serialize a row, deserialize it, assert match
    └── test_table   # Insert 10k rows and ensure no crashes