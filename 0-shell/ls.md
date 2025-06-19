# LS Command Tests

## Basic listing
- `ls`  
  ✅ Lists non-hidden files/directories in the current directory.

## Flag: `-a` (all)
- `ls -a`  
  ✅ Includes `.` and `..` as well as hidden files (starting with `.`).

## Flag: `-l` (long listing)
- `ls -l`  
  ✅ Displays permissions, links, owner, group, size, modified time, and name.

## Flag: `-F` (classify)
- `ls -F`  
  ✅ Appends `/` to directories, `*` to executables, etc. (only `/` implemented).

## Combined flags
- `ls -la`  
  ✅ Long listing including hidden files.
  
- `ls -al`  
  ✅ Same as `-la`.

- `ls -lF`  
  ✅ Long listing with directory indicators.

- `ls -Fl`  
  ✅ Same as `-lF`.

- `ls -aF`  
  ✅ All files (including hidden), with `/` for directories.

- `ls -Fa`  
  ✅ Same as `-aF`.

- `ls -laF`, `ls -lFa`, `ls -alF`, etc.  
  ✅ Long listing, hidden files, and directory slashes.

## Error cases
- `ls invalid.txt`  
  ❌ Returns error: no path arguments allowed (if not supported).

- `ls --help`  
  ❌ Returns error (if not supported).

- `ls -z`  
  ❌ Returns: Invalid flag.

- `ls -alZ`  
  ❌ Returns: Invalid flag `Z`.

## Special entries
- `ls -a`  
  ✅ Must include `.` and `..`.

- `ls -alF`  
  ✅ `.` and `..` should appear with full metadata and `/`.
