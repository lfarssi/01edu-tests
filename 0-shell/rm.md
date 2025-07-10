### 1. **Delete the current directory**

```bash
rm -r .
```
### 1. **Delete the current directory**

```bash
rm -r .
# ⚠️ Dangerous: may delete all contents of the current directory if not protected
# Recommended: refuse this or warn strongly
```

---

### 2. **Refuse to remove parent directory**

```bash
rm -r ..
# ❌ Refused: rm should never allow removing '..'
# Output: rm: refusing to remove '.' or '..' directory: skipping '..'
```

---

### 3. **Refuse to remove '.' and '..' in relative form**

```bash
rm -r ./.
rm -r ./..
rm -r ././.
# ❌ Refused: these are equivalent to '.' and '..'
# Output: rm: refusing to remove '.' or '..' directory: skipping './.'
```

---

### 4. **Refuse to remove root directory**

```bash
rm -r /
# ❌ Refused: dangerous
# Output: rm: it is dangerous to operate recursively on '/'
```

---

### 5. **Try to remove a non-existent file**

```bash
rm -r no_such_file
# ❌ Error: No such file or directory
```

---

### 6. **Try to remove a file (not a directory)**

```bash
touch file
rm -r file
# ✅ Should work even without `-r` if file
```

---

### 7. **Try to remove a non-empty directory**

```bash
mkdir dir
touch dir/file
rm -r dir
# ✅ Works: removes dir and its contents
```

---

### 8. **Try to remove multiple valid files and directories**

```bash
touch a b
mkdir c && touch c/file
rm -r a b c
# ✅ Works: removes a, b, and directory c with its contents
```

---

### 9. **Refuse to remove '.' or '..' mixed with valid files**

```bash
touch file
rm -r . file
# ❌ Refuse `.` but still remove `file`
# Output: rm: refusing to remove '.' or '..' directory: skipping '.'
```

---

### 10. **Refuse to remove symlink to `/` or parent**

```bash
ln -s / symlink_to_root
ln -s .. symlink_to_parent
rm -r symlink_to_root symlink_to_parent
# ✅ Symlinks themselves can be deleted — this is allowed
# ⚠️ But don't follow them recursively
```


